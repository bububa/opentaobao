package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeapartmentline 公寓上下架
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
func Alibabaalihousenewhomeapartmentline(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeapartmentlineAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeapartmentlineAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeapartmentlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
