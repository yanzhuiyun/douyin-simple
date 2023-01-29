package mexample

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

//自定义数据库行为
//redis，mysql

type (
	xexampleModel interface {
		FindContentByUsername(ctx context.Context, username string) ([]string, error)
	}

	xExampleModel struct {
		conn  sqlx.SqlConn
		rd    *redis.Redis
		table string
	}
)

func newxExampleModel(conn sqlx.SqlConn, rd *redis.Redis) *xExampleModel {
	return &xExampleModel{
		conn:  conn,
		table: "`example`",
		rd:    rd,
	}
}

func (x *xExampleModel) FindContentByUsername(ctx context.Context, username string) ([]string, error) {
	query := fmt.Sprintf("SELECT content FROM %s WHERE %s=?", x.table, "username")
	var resp []string
	return resp, x.conn.QueryRowCtx(ctx, &resp, query, username)
}
