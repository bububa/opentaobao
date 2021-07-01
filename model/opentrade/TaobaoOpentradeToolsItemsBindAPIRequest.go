package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeToolsItemsBindAPIRequest
交易开放商品绑定 API请求
taobao.opentrade.tools.items.bind

交易开放商品绑定 */
type TaobaoOpentradeToolsItemsBindAPIRequest struct {
	model.Params
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
	// 待绑定商品id
	_itemIds []int64
}

// New
