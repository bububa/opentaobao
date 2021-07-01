package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemItemGetAPIRequest
获取全渠道门店商品 API请求
taobao.omniitem.item.get

通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息 */
type TaobaoOmniitemItemGetAPIRequest struct {
	model.Params
	// 分页当前页数
	_pageNo int64
	// 分页单页大小
	_pageSize int64
	// 可选，指定获取的商品id
	_itemId int64
	// 可选，指定获取的商品外部id
	_outerId string
}

// New
