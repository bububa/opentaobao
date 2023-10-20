package cityretail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cityretail"
)

// TmallCityretailWmflOrderLogisticsQuery 完美履约订单物流状态查询接口
// tmall.cityretail.wmfl.order.logistics.query
//
// 完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
func TmallCityretailWmflOrderLogisticsQuery(clt *core.SDKClient, req *cityretail.TmallCityretailWmflOrderLogisticsQueryAPIRequest, session string) (*cityretail.TmallCityretailWmflOrderLogisticsQueryAPIResponse, error) {
	var resp cityretail.TmallCityretailWmflOrderLogisticsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
