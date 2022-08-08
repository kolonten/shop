package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemLoginLogModel = (*customSystemLoginLogModel)(nil)

type (
	// SystemLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemLoginLogModel.
	SystemLoginLogModel interface {
		systemLoginLogModel
	}

	customSystemLoginLogModel struct {
		*defaultSystemLoginLogModel
	}
)

// NewSystemLoginLogModel returns a model for the database table.
func NewSystemLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf) SystemLoginLogModel {
	return &customSystemLoginLogModel{
		defaultSystemLoginLogModel: newSystemLoginLogModel(conn, c),
	}
}
