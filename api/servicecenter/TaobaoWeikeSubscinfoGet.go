package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikeSubscinfoGet 需求订单查询接口
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
func TaobaoWeikeSubscinfoGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoWeikeSubscinfoGetAPIRequest, resp *servicecenter.TaobaoWeikeSubscinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
