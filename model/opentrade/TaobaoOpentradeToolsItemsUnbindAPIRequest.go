package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeToolsItemsUnbindAPIRequest
交易开放商品解绑 API请求
taobao.opentrade.tools.items.unbind

交易开放商品解绑 */
type TaobaoOpentradeToolsItemsUnbindAPIRequest struct {
	model.Params
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
	// 商品id
	_itemIds []int64
}

// New
