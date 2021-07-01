package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderDetailSearchAPIRequest
订单详情查询 API请求
taobao.xhotel.order.detail.search

提供订单详情查询 */
type TaobaoXhotelOrderDetailSearchAPIRequest struct {
	model.Params
	// 外部订单号
	_outOid string
	// 外部订单号
	_tid int64
}

// New
