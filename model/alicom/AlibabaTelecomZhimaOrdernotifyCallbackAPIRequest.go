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

// NewAlibabaTelecomZhimaOrdernotifyCallbackRequest 初始化AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest对象
func NewAlibabaTelecomZhimaOrdernotifyCallbackRequest() *AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest {
	return &AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.telecom.zhima.ordernotify.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 入参对象
func (r *AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest) SetParam0(_param0 *OrderStatusNotifyRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaTelecomZhimaOrdernotifyCallbackAPIRequest) GetParam0() *OrderStatusNotifyRequest {
	return r._param0
}
