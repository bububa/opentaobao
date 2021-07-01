package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTxpItemItemlistgetAPIRequest
淘小铺商品接口 API请求
taobao.txp.item.itemlistget

淘小铺商品的查询服务。 */
type TaobaoTxpItemItemlistgetAPIRequest struct {
	model.Params
	// 第几页
	_beginPage int64
	// 每页多少条
	_pageSize int64
}

// New
