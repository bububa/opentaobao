package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyShopUpdaterange 更新渠道店销售范围
// alibaba.aelophy.shop.updaterange
//
// 更新渠道店销售范围
func AlibabaAelophyShopUpdaterange(clt *core.SDKClient, req *wdk.AlibabaAelophyShopUpdaterangeAPIRequest, session string) (*wdk.AlibabaAelophyShopUpdaterangeAPIResponse, error) {
	var resp wdk.AlibabaAelophyShopUpdaterangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
