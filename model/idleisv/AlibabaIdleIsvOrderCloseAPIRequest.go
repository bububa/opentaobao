package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderCloseAPIRequest
服务商闲鱼卖家主动关闭订单 API请求
alibaba.idle.isv.order.close

供外部服务商 isv 提供卖家主动关闭交易订单的功能 */
type AlibabaIdleIsvOrderCloseAPIRequest struct {
	model.Params
	// 输入参数
	_isvAppraiseIsvOrderCloseDto *AppraiseIsvOrderCloseDto
}

// New
