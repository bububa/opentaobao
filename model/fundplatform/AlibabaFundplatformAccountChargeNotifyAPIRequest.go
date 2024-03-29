package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountChargeNotifyAPIRequest 账户充值成功通知 API请求
// alibaba.fundplatform.account.charge.notify
//
// 通知外部业务方充值成功
type AlibabaFundplatformAccountChargeNotifyAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabaFundplatformAccountChargeNotifyStruct
}

// NewAlibabaFundplatformAccountChargeNotifyRequest 初始化AlibabaFundplatformAccountChargeNotifyAPIRequest对象
func NewAlibabaFundplatformAccountChargeNotifyRequest() *AlibabaFundplatformAccountChargeNotifyAPIRequest {
	return &AlibabaFundplatformAccountChargeNotifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformAccountChargeNotifyAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.charge.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 入参对象
func (r *AlibabaFundplatformAccountChargeNotifyAPIRequest) SetRequest(_request *AlibabaFundplatformAccountChargeNotifyStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetRequest() *AlibabaFundplatformAccountChargeNotifyStruct {
	return r._request
}

var poolAlibabaFundplatformAccountChargeNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformAccountChargeNotifyRequest()
	},
}

// GetAlibabaFundplatformAccountChargeNotifyRequest 从 sync.Pool 获取 AlibabaFundplatformAccountChargeNotifyAPIRequest
func GetAlibabaFundplatformAccountChargeNotifyAPIRequest() *AlibabaFundplatformAccountChargeNotifyAPIRequest {
	return poolAlibabaFundplatformAccountChargeNotifyAPIRequest.Get().(*AlibabaFundplatformAccountChargeNotifyAPIRequest)
}

// ReleaseAlibabaFundplatformAccountChargeNotifyAPIRequest 将 AlibabaFundplatformAccountChargeNotifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformAccountChargeNotifyAPIRequest(v *AlibabaFundplatformAccountChargeNotifyAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformAccountChargeNotifyAPIRequest.Put(v)
}
