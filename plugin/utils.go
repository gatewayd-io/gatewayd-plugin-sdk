package plugin

import v1 "github.com/gatewayd-io/gatewayd-plugin-sdk/plugin/v1"

func GetAttr(req *v1.Struct, key string, defaultValue interface{}) interface{} {
	if req == nil {
		return defaultValue
	}

	if value, ok := req.Fields[key]; ok {
		return value.AsInterface()
	}

	return defaultValue
}
