package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeSpecialItemsQueryAPIRequest
专属下单获取商品绑定信息 API请求
taobao.opentrade.special.items.query

专属下单获取商品绑定信息 */
type TaobaoOpentradeSpecialItemsQueryAPIRequest struct {
	model.Params
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
}

// New
