package brock_test

import "testing"

func Test_Brock(t *testing.T) {
	t.Parallel()

	_ = t.Run("crypto", test_crypto)
	_ = t.Run("fsm", test_fsm)
	_ = t.Run("http", test_http)
	_ = t.Run("parser", test_parser)
	_ = t.Run("sql", test_sql)
}
