package cityretail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cityretail"
)

// TmallCityretailWmflOrderLogisticsQuery 完美履约订单物流状态查询接口
// tmall.cityretail.wmfl.order.logistics.query
//
// 完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
func TmallCityretailWmflOrderLogisticsQuery(ctx context.Context, clt *core.SDKClient, req *cityretail.TmallCityretailWmflOrderLogisticsQueryAPIRequest, resp *cityretail.TmallCityretailWmflOrderLogisticsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
