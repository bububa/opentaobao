package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeApartmentLine 公寓上下架
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
func AlibabaAlihouseNewhomeApartmentLine(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeApartmentLineAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeApartmentLineAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeApartmentLineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
