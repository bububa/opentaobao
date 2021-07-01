package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkCouponGetAPIRequest
淘宝客-公用-阿里妈妈推广券详情查询 API请求
taobao.tbk.coupon.get

传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。 */
type TaobaoTbkCouponGetAPIRequest struct {
	model.Params
	// 带券ID与商品ID的加密串
	_me string
	// 商品ID
	_itemId int64
	// 券ID
	_activityId string
}

// New
