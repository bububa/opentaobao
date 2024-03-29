package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaBuynowOrderCreate 阿里巴巴买家buynow下单接口
// alibaba.buynow.order.create
//
// 阿里巴巴买家下单接口
func AlibabaBuynowOrderCreate(clt *core.SDKClient, req *icbudropshipping.AlibabaBuynowOrderCreateAPIRequest, resp *icbudropshipping.AlibabaBuynowOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
