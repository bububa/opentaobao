package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabamallitemcenterentitledservicesupplierquery 根据天猫id查询门店服务授权
// alibaba.mallitemcenter.entitledservice.supplier.query
//
// 根据天猫id查询门店服务授权
func Alibabamallitemcenterentitledservicesupplierquery(clt *core.SDKClient, req *tmallservice.AlibabamallitemcenterentitledservicesupplierqueryAPIRequest, session string) (*tmallservice.AlibabamallitemcenterentitledservicesupplierqueryAPIResponse, error) {
	var resp tmallservice.AlibabamallitemcenterentitledservicesupplierqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
