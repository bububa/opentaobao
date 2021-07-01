package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentOrderGetAPIRequest
闲鱼帮卖订单查询 API请求
alibaba.idle.consignment.order.get

闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息 */
type AlibabaIdleConsignmentOrderGetAPIRequest struct {
	model.Params
	// 闲鱼订单ID
	_bizOrderId int64
}

// New
