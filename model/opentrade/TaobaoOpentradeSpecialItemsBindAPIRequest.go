package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeSpecialItemsBindAPIRequest
专属下单场景商品绑定 API请求
taobao.opentrade.special.items.bind

专属下单场景商品绑定 */
type TaobaoOpentradeSpecialItemsBindAPIRequest struct {
	model.Params
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
	// 本次待绑定的商品ID列表
	_itemIds []int64
}

// New
