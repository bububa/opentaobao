package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeToolsItemsQueryAPIRequest
交易开放获取商品绑定信息 API请求
taobao.opentrade.tools.items.query

交易开放获取商品绑定信息 */
type TaobaoOpentradeToolsItemsQueryAPIRequest struct {
	model.Params
	// 交易开放C端小程序ID
	_miniappId int64
}

// New
