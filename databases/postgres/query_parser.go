package postgres

import (
	"fmt"

	pgQuery "github.com/pganalyze/pg_query_go/v4"
)

const MinPgSQLMessageLength = 5

// GetQueryFromRequest decodes the request and returns the query.
func GetQueryFromRequest(req interface{}) (string, error) {
	var requestDecoded []byte

	if req, ok := req.([]byte); ok {
		requestDecoded = req
	} else {
		return "", fmt.Errorf("unknown request type: %T", req)
	}

	if len(requestDecoded) < MinPgSQLMessageLength {
		return "", nil
	}

	// The first byte is the message type.
	// The next 4 bytes are the length of the message.
	// The rest of the message is the query.
	// See https://www.postgresql.org/docs/13/protocol-message-formats.html
	// for more information.
	size := int(requestDecoded[1])<<24 + int(requestDecoded[2])<<16 + int(requestDecoded[3])<<8 + int(requestDecoded[4])
	return string(requestDecoded[MinPgSQLMessageLength:size]), nil
}

// isMulti checks if the query is a union, intersect, or except.
func isMulti(stmt *pgQuery.SelectStmt) bool {
	return stmt.GetOp() == pgQuery.SetOperation_SETOP_UNION ||
		stmt.GetOp() == pgQuery.SetOperation_SETOP_INTERSECT ||
		stmt.GetOp() == pgQuery.SetOperation_SETOP_EXCEPT
}

// getSingleTable returns the tables used in a query.
func getSingleTable(stmt *pgQuery.SelectStmt) []string {
	tables := []string{}
	for _, from := range stmt.FromClause {
		rangeVar := from.GetRangeVar()
		if rangeVar != nil {
			tables = append(tables, rangeVar.Relname)
		}
	}

	return tables
}

// getMultiTable returns the tables used in a union, intersect, or except query.
func getMultiTable(stmt *pgQuery.SelectStmt) []string {
	tables := []string{}
	// Get the tables from the left side.
	left := stmt.GetLarg()
	tables = append(tables, getSingleTable(left)...)
	// Get the tables from the right side.
	right := stmt.GetRarg()
	tables = append(tables, getSingleTable(right)...)

	return tables
}

// GetTablesFromQuery returns the tables used in a query.
func GetTablesFromQuery(query string) ([]string, error) {
	stmt, err := pgQuery.Parse(query)
	if err != nil {
		return nil, err
	}

	if len(stmt.Stmts) == 0 {
		return nil, nil
	}

	tables := []string{}

	for _, stmt := range stmt.Stmts {
		// Get the tables from the left and right side of the complex query.
		selectStatement := stmt.Stmt.GetSelectStmt()
		if isMulti(selectStatement) {
			tables = append(tables, getMultiTable(selectStatement)...)
		}

		// Get the table from the WITH clause.
		if withClause := stmt.Stmt.GetSelectStmt().GetWithClause(); withClause != nil {
			for _, cte := range withClause.Ctes {
				selectStmt := cte.GetCommonTableExpr().Ctequery.GetSelectStmt()
				if isMulti(selectStmt) {
					tables = append(tables, getMultiTable(selectStmt)...)
				} else {
					tables = append(tables, getSingleTable(selectStmt)...)
				}
			}
		} else {
			// Get the table from the FROM clause.
			if selectStatement := stmt.Stmt.GetSelectStmt(); selectStatement != nil {
				tables = append(tables, getSingleTable(selectStatement)...)
			}
		}

		if insertQuery := stmt.Stmt.GetInsertStmt(); insertQuery != nil {
			tables = append(tables, insertQuery.Relation.Relname)
		}

		if updateQuery := stmt.Stmt.GetUpdateStmt(); updateQuery != nil {
			tables = append(tables, updateQuery.Relation.Relname)
		}

		if deleteQuery := stmt.Stmt.GetDeleteStmt(); deleteQuery != nil {
			tables = append(tables, deleteQuery.Relation.Relname)
		}

		if truncateQuery := stmt.Stmt.GetTruncateStmt(); truncateQuery != nil {
			for _, relation := range truncateQuery.Relations {
				tables = append(tables, relation.GetRangeVar().Relname)
			}
		}

		if dropTable := stmt.Stmt.GetDropStmt(); dropTable != nil {
			for _, object := range dropTable.GetObjects() {
				for _, table := range object.GetList().GetItems() {
					tables = append(tables, table.String())
				}
			}
		}

		if alterTable := stmt.Stmt.GetAlterTableStmt(); alterTable != nil {
			tables = append(tables, alterTable.Relation.Relname)
		}
	}

	return tables, nil
}
