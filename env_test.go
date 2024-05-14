package osenv_test

import (
	"testing"

	"github.com/theplant/osenv"
)

var (
	dbURL     = osenv.Get("DB_URL", "string to connect to database", "localhost:5432")
	boolValue = osenv.GetBool("BOOL_VALUE", "bool value", false)
	intValue  = osenv.GetInt64("INT_VALUE", "int value", 0)
)

func TestBasic(t *testing.T) {
	t.Log(dbURL, boolValue, intValue)
}
