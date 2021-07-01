package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderShowAPIRequest
闲鱼回收订单查询V1.1 API请求
alibaba.idle.recycle.order.show

查询回收订单 */
type AlibabaIdleRecycleOrderShowAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// New
