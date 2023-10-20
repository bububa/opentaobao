package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderMakeSuccessAPIRequest 通知制卡成功 API请求
// alibaba.fundplatform.cardorder.make.success
//
// 当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
type AlibabaFundplatformCardorderMakeSuccessAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabaFundplatformCardorderMakeSuccessStruct
}

// NewAlibabaFundplatformCardorderMakeSuccessRequest 初始化AlibabaFundplatformCardorderMakeSuccessAPIRequest对象
func NewAlibabaFundplatformCardorderMakeSuccessRequest() *AlibabaFundplatformCardorderMakeSuccessAPIRequest {
	return &AlibabaFundplatformCardorderMakeSuccessAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformCardorderMakeSuccessAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.make.success"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 入参对象
func (r *AlibabaFundplatformCardorderMakeSuccessAPIRequest) SetRequest(_request *AlibabaFundplatformCardorderMakeSuccessStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetRequest() *AlibabaFundplatformCardorderMakeSuccessStruct {
	return r._request
}

var poolAlibabaFundplatformCardorderMakeSuccessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformCardorderMakeSuccessRequest()
	},
}

// GetAlibabaFundplatformCardorderMakeSuccessRequest 从 sync.Pool 获取 AlibabaFundplatformCardorderMakeSuccessAPIRequest
func GetAlibabaFundplatformCardorderMakeSuccessAPIRequest() *AlibabaFundplatformCardorderMakeSuccessAPIRequest {
	return poolAlibabaFundplatformCardorderMakeSuccessAPIRequest.Get().(*AlibabaFundplatformCardorderMakeSuccessAPIRequest)
}

// ReleaseAlibabaFundplatformCardorderMakeSuccessAPIRequest 将 AlibabaFundplatformCardorderMakeSuccessAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformCardorderMakeSuccessAPIRequest(v *AlibabaFundplatformCardorderMakeSuccessAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformCardorderMakeSuccessAPIRequest.Put(v)
}
