package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiSearch 查询面单服务订购及面单使用情况
// cainiao.waybill.ii.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
func CainiaoWaybillIiSearch(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiSearchAPIRequest, resp *waybill.CainiaoWaybillIiSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
