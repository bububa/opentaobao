package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomesystemseller 商品发布授权
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
func Alibabaalihousenewhomesystemseller(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomesystemsellerAPIRequest, session string) (*alihouse.AlibabaalihousenewhomesystemsellerAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomesystemsellerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
