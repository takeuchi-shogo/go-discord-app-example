package database

import (
	"errors"
	"fmt"
)

const (
	MySQLDSN = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true"
)

func setDsn(dbEngine string) (string, error) {
	switch dbEngine {
	case DBMySQL:
		return fmt.Sprintf(MySQLDSN), nil
	}
	return "", fmt.Errorf("%w", errors.New("DBエンジンを設定してください"))
}
