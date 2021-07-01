package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaPricePromotionItemDelete
批量删除档期
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品 */
func AlibabaPricePromotionItemDelete(clt *core.SDKClient, req *wdk.AlibabaPricePromotionItemDeleteAPIRequest, session string) (*wdk.AlibabaPricePromotionItemDeleteAPIResponse, error) {
	var resp wdk.AlibabaPricePromotionItemDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
