package cityretail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cityretail"
)

// Tmallcityretailwmflorderlogisticsquery 完美履约订单物流状态查询接口
// tmall.cityretail.wmfl.order.logistics.query
//
// 完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
func Tmallcityretailwmflorderlogisticsquery(clt *core.SDKClient, req *cityretail.TmallcityretailwmflorderlogisticsqueryAPIRequest, session string) (*cityretail.TmallcityretailwmflorderlogisticsqueryAPIResponse, error) {
	var resp cityretail.TmallcityretailwmflorderlogisticsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
