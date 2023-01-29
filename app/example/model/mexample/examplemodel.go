package mexample

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExampleModel = (*customExampleModel)(nil)

type (
	// ExampleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExampleModel.
	ExampleModel interface {
		exampleModel
		xexampleModel
	}

	customExampleModel struct {
		*defaultExampleModel
		*xExampleModel
	}
)

// NewExampleModel returns a model for the database table.
func NewExampleModel(conn sqlx.SqlConn, rd ...*redis.Redis) ExampleModel {
	var rc *redis.Redis
	if len(rd) > 0 {
		rc = rd[0]
	}
	return &customExampleModel{
		defaultExampleModel: newExampleModel(conn),
		xExampleModel:       newxExampleModel(conn, rc),
	}
}
