package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderAdjustpriceAPIRequest
闲鱼服务商订单价格修改接口 API请求
alibaba.idle.isv.order.adjustprice

闲鱼用户通过授权的服务商修改订单价格和邮费 */
type AlibabaIdleIsvOrderAdjustpriceAPIRequest struct {
	model.Params
	// 输入参数
	_paramAdjustOrderPrice *IsvAdjustOrderPriceDto
}

// New
