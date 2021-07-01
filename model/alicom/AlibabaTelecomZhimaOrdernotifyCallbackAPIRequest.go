package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest
阿里通信芝麻订单通知 API请求
alibaba.telecom.zhima.ordernotify.callback

商家通知阿里通信，芝麻订单状态，阿里通信侧进行代扣支付、发货并确认收货 */
type AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *OrderStatusNotifyRequest
}

// New
