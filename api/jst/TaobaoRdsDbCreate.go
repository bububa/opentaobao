package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbCreate rds创建数据库
// taobao.rds.db.create
//
// 在rds实例里创建数据库
func TaobaoRdsDbCreate(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoRdsDbCreateAPIRequest, resp *jst.TaobaoRdsDbCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
