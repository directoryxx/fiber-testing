package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenDB(t *testing.T) {
	dsn := "root:passworddev2020@tcp(127.0.0.1:3306)/rest-api-test?charset=utf8mb4&parseTime=True"
	db,err := OpenDBMysql(dsn)
	if assert.NoError(t, err) {
		assert.NotNil(t, db)
	}
}
