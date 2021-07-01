package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderQueryAPIRequest
闲鱼回收订单查询V1 API请求
alibaba.idle.recycle.order.query

查询回收订单 */
type AlibabaIdleRecycleOrderQueryAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// New
