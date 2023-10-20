package alihealthpw

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchronodeAPIRequest 合作方同步状态至阿里健康 API请求
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
type AlibabaAlihealthPwSpecialSynchronodeAPIRequest struct {
	model.Params
	// 状态信息入参
	_body *SNodeDto
}

// NewAlibabaAlihealthPwSpecialSynchronodeRequest 初始化AlibabaAlihealthPwSpecialSynchronodeAPIRequest对象
func NewAlibabaAlihealthPwSpecialSynchronodeRequest() *AlibabaAlihealthPwSpecialSynchronodeAPIRequest {
	return &AlibabaAlihealthPwSpecialSynchronodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwSpecialSynchronodeAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchronode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 状态信息入参
func (r *AlibabaAlihealthPwSpecialSynchronodeAPIRequest) SetBody(_body *SNodeDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwSpecialSynchronodeAPIRequest) GetBody() *SNodeDto {
	return r._body
}

var poolAlibabaAlihealthPwSpecialSynchronodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwSpecialSynchronodeRequest()
	},
}

// GetAlibabaAlihealthPwSpecialSynchronodeRequest 从 sync.Pool 获取 AlibabaAlihealthPwSpecialSynchronodeAPIRequest
func GetAlibabaAlihealthPwSpecialSynchronodeAPIRequest() *AlibabaAlihealthPwSpecialSynchronodeAPIRequest {
	return poolAlibabaAlihealthPwSpecialSynchronodeAPIRequest.Get().(*AlibabaAlihealthPwSpecialSynchronodeAPIRequest)
}

// ReleaseAlibabaAlihealthPwSpecialSynchronodeAPIRequest 将 AlibabaAlihealthPwSpecialSynchronodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwSpecialSynchronodeAPIRequest(v *AlibabaAlihealthPwSpecialSynchronodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwSpecialSynchronodeAPIRequest.Put(v)
}
