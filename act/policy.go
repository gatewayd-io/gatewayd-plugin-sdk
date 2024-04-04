package act

import (
	"context"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

type IPolicy interface {
	MustCompile(opts ...expr.Option) error
	Eval(ctx context.Context, input Input) (any, error)
}

type Policy struct {
	prg  *vm.Program   `json:"-"` // Compiled policy
	opts []expr.Option `json:"-"` // Extra options for compilation

	Name     string         `json:"name"`
	Policy   string         `json:"policy"`
	Metadata map[string]any `json:"metadata"`
}

var _ IPolicy = (*Policy)(nil)

func (p *Policy) MustCompile(extraOpts ...expr.Option) error {
	var err error

	extraOpts = append(p.opts, extraOpts...)

	p.prg, err = expr.Compile(p.Policy, extraOpts...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Policy) Eval(ctx context.Context, input Input) (any, error) {
	result := make(chan interface{})
	errChan := make(chan error)

	go func() {
		output, err := expr.Run(p.prg, input)
		if err != nil {
			errChan <- err
			return
		}
		result <- output
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case out := <-result:
		return out, nil
	case err := <-errChan:
		return nil, err
	}
}

func MustNewPolicy(
	name string, policy string, metadata map[string]any, opts ...expr.Option,
) *Policy {
	plc := Policy{
		Name:     name,
		Policy:   policy,
		Metadata: metadata,
		opts:     opts,
	}

	if err := plc.MustCompile(opts...); err != nil {
		return nil
	}

	return &plc
}

func NewPolicy(
	name string, policy string, metadata map[string]any, opts ...expr.Option,
) (*Policy, error) {
	plc := &Policy{
		Name:     name,
		Policy:   policy,
		Metadata: metadata,
		opts:     opts,
	}

	if err := plc.MustCompile(opts...); err != nil {
		return nil, err
	}

	return plc, nil
}
