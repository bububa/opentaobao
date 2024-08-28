package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbGet 查询rds下的数据库
// taobao.rds.db.get
//
// 查询rds实例下的数据库
func TaobaoRdsDbGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoRdsDbGetAPIRequest, resp *jst.TaobaoRdsDbGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
