package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkPurchasePrice
rt回传采购价
alibaba.wdk.purchase.price

猫超共享库存项目-rt回传采购价 */
func AlibabaWdkPurchasePrice(clt *core.SDKClient, req *wdk.AlibabaWdkPurchasePriceAPIRequest, session string) (*wdk.AlibabaWdkPurchasePriceAPIResponse, error) {
	var resp wdk.AlibabaWdkPurchasePriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
