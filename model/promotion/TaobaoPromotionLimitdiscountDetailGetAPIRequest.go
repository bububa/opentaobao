package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionLimitdiscountDetailGetAPIRequest
限时打折详情查询 API请求
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。 */
type TaobaoPromotionLimitdiscountDetailGetAPIRequest struct {
	model.Params
	// 限时打折ID。这个针对查询唯一限时打折情况。
	_limitDiscountId int64
}

// New
