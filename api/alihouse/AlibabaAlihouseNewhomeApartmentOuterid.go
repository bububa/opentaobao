package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeapartmentouterid 公寓更新outerid
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
func Alibabaalihousenewhomeapartmentouterid(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeapartmentouteridAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeapartmentouteridAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeapartmentouteridAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
