package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* AlibabaMallitemcenterEntitledserviceSupplierQuery
根据天猫id查询门店服务授权
alibaba.mallitemcenter.entitledservice.supplier.query

根据天猫id查询门店服务授权 */
func AlibabaMallitemcenterEntitledserviceSupplierQuery(clt *core.SDKClient, req *tmallservice.AlibabaMallitemcenterEntitledserviceSupplierQueryAPIRequest, session string) (*tmallservice.AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse, error) {
	var resp tmallservice.AlibabaMallitemcenterEntitledserviceSupplierQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
