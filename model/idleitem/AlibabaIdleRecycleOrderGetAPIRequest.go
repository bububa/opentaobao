package idleitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderGetAPIRequest
闲鱼回收订单查询V2 API请求
alibaba.idle.recycle.order.get

闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单 */
type AlibabaIdleRecycleOrderGetAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// New
