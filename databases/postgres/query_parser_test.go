package postgres

import (
	"encoding/base64"
	"testing"

	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/stretchr/testify/assert"
)

func testQueryRequest() (string, string) {
	query := "SELECT * FROM users"
	queryMsg := pgproto3.Query{String: query}
	// Encode the data to base64.
	queryBytes, _ := queryMsg.Encode(nil)
	return query, base64.StdEncoding.EncodeToString(queryBytes)
}

func Test_GetQueryFromRequest(t *testing.T) {
	query, request := testQueryRequest()
	// Decode the request and extract the query.
	decodedQuery, err := GetQueryFromRequest(request)
	assert.Nil(t, err)
	assert.Equal(t, query, decodedQuery)
}

func Test_GetQueryFromRequest_Empty(t *testing.T) {
	// Decode the request and extract the query.
	decodedQuery, err := GetQueryFromRequest("")
	assert.Nil(t, err)
	assert.Equal(t, "", decodedQuery)
}

func Test_GetQueryFromRequest_Invalid(t *testing.T) {
	// Decode the request and extract the query.
	decodedQuery, err := GetQueryFromRequest("invalid")
	assert.NotNil(t, err)
	assert.Equal(t, "", decodedQuery)
}

func Test_GetQueryFromRequest_Short(t *testing.T) {
	// Decode the request and extract the query.
	decodedQuery, err := GetQueryFromRequest("Q")
	assert.NotNil(t, err)
	assert.Equal(t, "", decodedQuery)
}

func Test_GetQueryFromRequest_Shorter(t *testing.T) {
	// Decode the request and extract the query.
	decodedQuery, err := GetQueryFromRequest("QAAAA")
	assert.NotNil(t, err)
	assert.Equal(t, "", decodedQuery)
}

func Test_GetTablesFromQuery(t *testing.T) {
	query := "SELECT * FROM users"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_Multiple(t *testing.T) {
	query := "SELECT * FROM users, posts"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_Union(t *testing.T) {
	query := "SELECT * FROM users UNION SELECT * FROM posts"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_UnionAll(t *testing.T) {
	query := "SELECT * FROM users UNION ALL SELECT * FROM posts"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_Intersection(t *testing.T) {
	query := "SELECT * FROM users INTERSECT SELECT * FROM posts"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_Except(t *testing.T) {
	query := "SELECT * FROM users EXCEPT SELECT * FROM posts"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_With(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_WithMultiple(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users), t2 AS (SELECT * FROM posts) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_WithUnion(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION SELECT * FROM posts) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_WithUnionAll(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION ALL SELECT * FROM posts) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_WithIntersection(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users INTERSECT SELECT * FROM posts) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_WithExcept(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users EXCEPT SELECT * FROM posts) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleUnion(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION SELECT * FROM posts), t2 AS (SELECT * FROM comments UNION SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleUnionAll(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION ALL SELECT * FROM posts), t2 AS (SELECT * FROM comments UNION ALL SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleIntersection(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users INTERSECT SELECT * FROM posts), t2 AS (SELECT * FROM comments INTERSECT SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleExcept(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users EXCEPT SELECT * FROM posts), t2 AS (SELECT * FROM comments EXCEPT SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleUnionAndIntersection(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION SELECT * FROM posts), t2 AS (SELECT * FROM comments INTERSECT SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleUnionAndExcept(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION SELECT * FROM posts), t2 AS (SELECT * FROM comments EXCEPT SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleUnionAllAndIntersection(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION ALL SELECT * FROM posts), t2 AS (SELECT * FROM comments INTERSECT SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_WithMultipleUnionAllAndExcept(t *testing.T) {
	query := "WITH t AS (SELECT * FROM users UNION ALL SELECT * FROM posts), t2 AS (SELECT * FROM comments EXCEPT SELECT * FROM likes) SELECT * FROM t"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "posts", "comments", "likes"}, tables)
}

func Test_GetTablesFromQuery_Insert(t *testing.T) {
	query := "INSERT INTO users SELECT * FROM posts"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_InsertMultiple(t *testing.T) {
	query := "INSERT INTO users SELECT * FROM posts; INSERT INTO comments SELECT * FROM likes"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "comments"}, tables)
}

func Test_GetTablesFromQuery_InsertWithUnion(t *testing.T) {
	query := "INSERT INTO users SELECT * FROM posts UNION SELECT * FROM comments"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_Update(t *testing.T) {
	query := "UPDATE users SET name = 'John' WHERE id = 1"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_UpdateMultiple(t *testing.T) {
	query := "UPDATE users SET name = 'John' WHERE id = 1; UPDATE comments SET name = 'John' WHERE id = 1"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "comments"}, tables)
}

func Test_GetTablesFromQuery_Delete(t *testing.T) {
	query := "DELETE FROM users WHERE id = 1"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_DeleteMultiple(t *testing.T) {
	query := "DELETE FROM users WHERE id = 1; DELETE FROM comments WHERE id = 1"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "comments"}, tables)
}

func Test_GetTablesFromQuery_Truncate(t *testing.T) {
	query := "TRUNCATE users"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_TruncateMultiple(t *testing.T) {
	query := "TRUNCATE users, comments"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "comments"}, tables)
}

func Test_GetTablesFromQuery_Drop(t *testing.T) {
	query := "DROP TABLE users"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_DropMultiple(t *testing.T) {
	query := "DROP TABLE users, comments"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "comments"}, tables)
}

func Test_GetTablesFromQuery_Alter(t *testing.T) {
	query := "ALTER TABLE users ADD COLUMN name VARCHAR(255)"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users"}, tables)
}

func Test_GetTablesFromQuery_AlterMultiple(t *testing.T) {
	query := "ALTER TABLE users ADD COLUMN name VARCHAR(255); ALTER TABLE comments ADD COLUMN name VARCHAR(255)"
	tables, err := GetTablesFromQuery(query)
	assert.Nil(t, err)
	assert.Equal(t, []string{"users", "comments"}, tables)
}
