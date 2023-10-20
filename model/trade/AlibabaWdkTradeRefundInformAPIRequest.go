package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundInformAPIRequest 外部渠道通知淘鲜达退款成功接口 API请求
// alibaba.wdk.trade.refund.inform
//
// 该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
type AlibabaWdkTradeRefundInformAPIRequest struct {
	model.Params
	// 通知退款成功请求
	_informRefundSuccessRequest *InformRefundSuccessRequest
}

// NewAlibabaWdkTradeRefundInformRequest 初始化AlibabaWdkTradeRefundInformAPIRequest对象
func NewAlibabaWdkTradeRefundInformRequest() *AlibabaWdkTradeRefundInformAPIRequest {
	return &AlibabaWdkTradeRefundInformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeRefundInformAPIRequest) Reset() {
	r._informRefundSuccessRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundInformAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.inform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundInformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeRefundInformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInformRefundSuccessRequest is InformRefundSuccessRequest Setter
// 通知退款成功请求
func (r *AlibabaWdkTradeRefundInformAPIRequest) SetInformRefundSuccessRequest(_informRefundSuccessRequest *InformRefundSuccessRequest) error {
	r._informRefundSuccessRequest = _informRefundSuccessRequest
	r.Set("inform_refund_success_request", _informRefundSuccessRequest)
	return nil
}

// GetInformRefundSuccessRequest InformRefundSuccessRequest Getter
func (r AlibabaWdkTradeRefundInformAPIRequest) GetInformRefundSuccessRequest() *InformRefundSuccessRequest {
	return r._informRefundSuccessRequest
}

var poolAlibabaWdkTradeRefundInformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeRefundInformRequest()
	},
}

// GetAlibabaWdkTradeRefundInformRequest 从 sync.Pool 获取 AlibabaWdkTradeRefundInformAPIRequest
func GetAlibabaWdkTradeRefundInformAPIRequest() *AlibabaWdkTradeRefundInformAPIRequest {
	return poolAlibabaWdkTradeRefundInformAPIRequest.Get().(*AlibabaWdkTradeRefundInformAPIRequest)
}

// ReleaseAlibabaWdkTradeRefundInformAPIRequest 将 AlibabaWdkTradeRefundInformAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeRefundInformAPIRequest(v *AlibabaWdkTradeRefundInformAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeRefundInformAPIRequest.Put(v)
}
