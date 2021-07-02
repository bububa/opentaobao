package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionItemAdd 新增档期商品
// alibaba.price.promotion.item.add
//
// 批量新增档期活动商品
func AlibabaPricePromotionItemAdd(clt *core.SDKClient, req *wdk.AlibabaPricePromotionItemAddAPIRequest, session string) (*wdk.AlibabaPricePromotionItemAddAPIResponse, error) {
	var resp wdk.AlibabaPricePromotionItemAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
