package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbDelete RDS数据库删除
// taobao.rds.db.delete
//
// 通过api删除用户RDS的数据库
func TaobaoRdsDbDelete(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoRdsDbDeleteAPIRequest, resp *jst.TaobaoRdsDbDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
