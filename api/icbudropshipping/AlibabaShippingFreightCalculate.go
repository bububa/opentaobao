package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabashippingfreightcalculate 阿里巴巴商品运费计算查询接口
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
func Alibabashippingfreightcalculate(clt *core.SDKClient, req *icbudropshipping.AlibabashippingfreightcalculateAPIRequest, session string) (*icbudropshipping.AlibabashippingfreightcalculateAPIResponse, error) {
	var resp icbudropshipping.AlibabashippingfreightcalculateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
