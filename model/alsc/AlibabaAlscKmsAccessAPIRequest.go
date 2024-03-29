package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscKmsAccessAPIRequest 本地生活风控数据接入 API请求
// alibaba.alsc.kms.access
//
// 第三方使用本地生活数据对外提供服务，上报访问日志信息接口
type AlibabaAlscKmsAccessAPIRequest struct {
	model.Params
	// 接入参数
	_requestdata string
}

// NewAlibabaAlscKmsAccessRequest 初始化AlibabaAlscKmsAccessAPIRequest对象
func NewAlibabaAlscKmsAccessRequest() *AlibabaAlscKmsAccessAPIRequest {
	return &AlibabaAlscKmsAccessAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscKmsAccessAPIRequest) Reset() {
	r._requestdata = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscKmsAccessAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.kms.access"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscKmsAccessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscKmsAccessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestdata is Requestdata Setter
// 接入参数
func (r *AlibabaAlscKmsAccessAPIRequest) SetRequestdata(_requestdata string) error {
	r._requestdata = _requestdata
	r.Set("requestdata", _requestdata)
	return nil
}

// GetRequestdata Requestdata Getter
func (r AlibabaAlscKmsAccessAPIRequest) GetRequestdata() string {
	return r._requestdata
}

var poolAlibabaAlscKmsAccessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscKmsAccessRequest()
	},
}

// GetAlibabaAlscKmsAccessRequest 从 sync.Pool 获取 AlibabaAlscKmsAccessAPIRequest
func GetAlibabaAlscKmsAccessAPIRequest() *AlibabaAlscKmsAccessAPIRequest {
	return poolAlibabaAlscKmsAccessAPIRequest.Get().(*AlibabaAlscKmsAccessAPIRequest)
}

// ReleaseAlibabaAlscKmsAccessAPIRequest 将 AlibabaAlscKmsAccessAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscKmsAccessAPIRequest(v *AlibabaAlscKmsAccessAPIRequest) {
	v.Reset()
	poolAlibabaAlscKmsAccessAPIRequest.Put(v)
}
