package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// Alibabalstvendingshippingcallback 售货机出货回传接口
// alibaba.lst.vending.shipping.callback
//
// 零售通自动售货机商品出货回传接口，同步商品出库最新状态。
func Alibabalstvendingshippingcallback(clt *core.SDKClient, req *lstvending.AlibabalstvendingshippingcallbackAPIRequest, session string) (*lstvending.AlibabalstvendingshippingcallbackAPIResponse, error) {
	var resp lstvending.AlibabalstvendingshippingcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
