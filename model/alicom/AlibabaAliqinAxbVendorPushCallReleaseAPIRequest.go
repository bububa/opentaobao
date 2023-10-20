package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorPushCallReleaseAPIRequest 供应商推送通话结束事件 API请求
// alibaba.aliqin.axb.vendor.push.call.release
//
// 通话结束挂断的时候，供应商推送通话结束事件给阿里侧
type AlibabaAliqinAxbVendorPushCallReleaseAPIRequest struct {
	model.Params
	// end_call_request
	_endCallRequest *EndCallRequest
}

// NewAlibabaAliqinAxbVendorPushCallReleaseRequest 初始化AlibabaAliqinAxbVendorPushCallReleaseAPIRequest对象
func NewAlibabaAliqinAxbVendorPushCallReleaseRequest() *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest {
	return &AlibabaAliqinAxbVendorPushCallReleaseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) Reset() {
	r._endCallRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.push.call.release"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndCallRequest is EndCallRequest Setter
// end_call_request
func (r *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) SetEndCallRequest(_endCallRequest *EndCallRequest) error {
	r._endCallRequest = _endCallRequest
	r.Set("end_call_request", _endCallRequest)
	return nil
}

// GetEndCallRequest EndCallRequest Getter
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetEndCallRequest() *EndCallRequest {
	return r._endCallRequest
}

var poolAlibabaAliqinAxbVendorPushCallReleaseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinAxbVendorPushCallReleaseRequest()
	},
}

// GetAlibabaAliqinAxbVendorPushCallReleaseRequest 从 sync.Pool 获取 AlibabaAliqinAxbVendorPushCallReleaseAPIRequest
func GetAlibabaAliqinAxbVendorPushCallReleaseAPIRequest() *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest {
	return poolAlibabaAliqinAxbVendorPushCallReleaseAPIRequest.Get().(*AlibabaAliqinAxbVendorPushCallReleaseAPIRequest)
}

// ReleaseAlibabaAliqinAxbVendorPushCallReleaseAPIRequest 将 AlibabaAliqinAxbVendorPushCallReleaseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinAxbVendorPushCallReleaseAPIRequest(v *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) {
	v.Reset()
	poolAlibabaAliqinAxbVendorPushCallReleaseAPIRequest.Put(v)
}
