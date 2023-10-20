package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeApartmentLine 公寓上下架
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
func AlibabaAlihouseNewhomeApartmentLine(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeApartmentLineAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeApartmentLineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
