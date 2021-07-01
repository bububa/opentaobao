package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpPromotionGetAPIRequest
商品优惠详情查询 API请求
taobao.ump.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。 */
type TaobaoUmpPromotionGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
