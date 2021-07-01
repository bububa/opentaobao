package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemItemDeleteAPIRequest
全渠道商品删除 API请求
taobao.omniitem.item.delete

全渠道商品删除，能够对门店商品库商品进行删除动作 */
type TaobaoOmniitemItemDeleteAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
	// 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
	_itemId int64
	// 商品outerId
	_outerId string
}

// New
