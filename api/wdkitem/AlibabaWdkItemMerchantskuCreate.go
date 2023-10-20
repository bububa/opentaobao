package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantskuCreate 商家商品信息新建
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
func AlibabaWdkItemMerchantskuCreate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuCreateAPIRequest, session string) (*wdkitem.AlibabaWdkItemMerchantskuCreateAPIResponse, error) {
	var resp wdkitem.AlibabaWdkItemMerchantskuCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
