package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidlelogisticscompaniesquery 快递公司列表查询
// alibaba.idle.logistics.companies.query
//
// 支持发货的快递公司列表查询
func Alibabaidlelogisticscompaniesquery(clt *core.SDKClient, req *idleisv.AlibabaidlelogisticscompaniesqueryAPIRequest, session string) (*idleisv.AlibabaidlelogisticscompaniesqueryAPIResponse, error) {
	var resp idleisv.AlibabaidlelogisticscompaniesqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
