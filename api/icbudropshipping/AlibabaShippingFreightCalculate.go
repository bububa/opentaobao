package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaShippingFreightCalculate 阿里巴巴商品运费计算查询接口
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
func AlibabaShippingFreightCalculate(clt *core.SDKClient, req *icbudropshipping.AlibabaShippingFreightCalculateAPIRequest, resp *icbudropshipping.AlibabaShippingFreightCalculateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
