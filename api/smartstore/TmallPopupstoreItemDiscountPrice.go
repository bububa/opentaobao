package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TmallPopupstoreItemDiscountPrice 商品优惠价格查询
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
func TmallPopupstoreItemDiscountPrice(clt *core.SDKClient, req *smartstore.TmallPopupstoreItemDiscountPriceAPIRequest, resp *smartstore.TmallPopupstoreItemDiscountPriceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
