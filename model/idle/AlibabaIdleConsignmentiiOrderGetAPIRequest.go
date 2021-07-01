package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentiiOrderGetAPIRequest
闲鱼寄卖V2订单查询 API请求
alibaba.idle.consignmentii.order.get

闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息 */
type AlibabaIdleConsignmentiiOrderGetAPIRequest struct {
	model.Params
	// 闲鱼订单ID
	_bizOrderId int64
}

// New
