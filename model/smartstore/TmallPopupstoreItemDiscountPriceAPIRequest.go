package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPopupstoreItemDiscountPriceAPIRequest
商品优惠价格查询 API请求
tmall.popupstore.item.discount.price

商品优惠价格查询 */
type TmallPopupstoreItemDiscountPriceAPIRequest struct {
	model.Params
	// 商品id列表
	_itemIds []int64
}

// New
