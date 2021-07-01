package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAppraiseOrderQueryAPIRequest
闲鱼验货担保订单详情查询V1 API请求
alibaba.idle.appraise.order.query

鉴定商调用该接口获取订单状态 */
type AlibabaIdleAppraiseOrderQueryAPIRequest struct {
	model.Params
	// orderId
	_bizOrderId int64
}

// New
