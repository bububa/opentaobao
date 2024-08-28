package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillISearch 查询面单服务订购及面单使用情况v1.0
// taobao.wlb.waybill.i.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
func TaobaoWlbWaybillISearch(ctx context.Context, clt *core.SDKClient, req *waybill.TaobaoWlbWaybillISearchAPIRequest, resp *waybill.TaobaoWlbWaybillISearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
