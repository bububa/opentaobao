package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsCompaniesGet 查询物流公司信息
// taobao.logistics.companies.get
//
// 查询淘宝网合作的物流公司信息，用于发货接口。
func TaobaoLogisticsCompaniesGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsCompaniesGetAPIRequest, resp *tblogistics.TaobaoLogisticsCompaniesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
