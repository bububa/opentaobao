package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMerchantItemUpdate 修改商家商品
// alibaba.wdk.merchant.item.update
//
// 修改商家商品
func AlibabaWdkMerchantItemUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantItemUpdateAPIRequest, session string) (*wdk.AlibabaWdkMerchantItemUpdateAPIResponse, error) {
	var resp wdk.AlibabaWdkMerchantItemUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
