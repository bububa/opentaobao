package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktraderefundinformAPIRequest 外部渠道通知淘鲜达退款成功接口 API请求
// alibaba.wdk.trade.refund.inform
//
// 该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
type AlibabawdktraderefundinformAPIRequest struct {
	model.Params
	// 通知退款成功请求
	_informRefundSuccessRequest *InformRefundSuccessRequest
}

// NewAlibabawdktraderefundinformRequest 初始化AlibabawdktraderefundinformAPIRequest对象
func NewAlibabawdktraderefundinformRequest() *AlibabawdktraderefundinformAPIRequest {
	return &AlibabawdktraderefundinformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktraderefundinformAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.inform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktraderefundinformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktraderefundinformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInformRefundSuccessRequest is InformRefundSuccessRequest Setter
// 通知退款成功请求
func (r *AlibabawdktraderefundinformAPIRequest) SetInformRefundSuccessRequest(_informRefundSuccessRequest *InformRefundSuccessRequest) error {
	r._informRefundSuccessRequest = _informRefundSuccessRequest
	r.Set("inform_refund_success_request", _informRefundSuccessRequest)
	return nil
}

// GetInformRefundSuccessRequest InformRefundSuccessRequest Getter
func (r AlibabawdktraderefundinformAPIRequest) GetInformRefundSuccessRequest() *InformRefundSuccessRequest {
	return r._informRefundSuccessRequest
}
