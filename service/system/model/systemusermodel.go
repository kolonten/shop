package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemUserModel = (*customSystemUserModel)(nil)

type (
	// SystemUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemUserModel.
	SystemUserModel interface {
		systemUserModel
	}

	customSystemUserModel struct {
		*defaultSystemUserModel
	}
)

// NewSystemUserModel returns a model for the database table.
func NewSystemUserModel(conn sqlx.SqlConn, c cache.CacheConf) SystemUserModel {
	return &customSystemUserModel{
		defaultSystemUserModel: newSystemUserModel(conn, c),
	}
}
