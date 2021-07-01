package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeSpecialItemsUnbindAPIRequest
专属下单场景商品解绑 API请求
taobao.opentrade.special.items.unbind

专属下单场景商品解绑 */
type TaobaoOpentradeSpecialItemsUnbindAPIRequest struct {
	model.Params
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
	// 本次待解绑的商品ID列表
	_itemIds []int64
}

// New
