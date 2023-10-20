package alihealthmedical

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalItemPublishAPIRequest 三方入驻-开通服务 API请求
// alibaba.alihealth.medical.item.publish
//
// 三方入驻-开通服务
type AlibabaAlihealthMedicalItemPublishAPIRequest struct {
	model.Params
	// 请求
	_request1 *ItemPublishRequest
}

// NewAlibabaAlihealthMedicalItemPublishRequest 初始化AlibabaAlihealthMedicalItemPublishAPIRequest对象
func NewAlibabaAlihealthMedicalItemPublishRequest() *AlibabaAlihealthMedicalItemPublishAPIRequest {
	return &AlibabaAlihealthMedicalItemPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalItemPublishAPIRequest) Reset() {
	r._request1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.item.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest1 is Request1 Setter
// 请求
func (r *AlibabaAlihealthMedicalItemPublishAPIRequest) SetRequest1(_request1 *ItemPublishRequest) error {
	r._request1 = _request1
	r.Set("request1", _request1)
	return nil
}

// GetRequest1 Request1 Getter
func (r AlibabaAlihealthMedicalItemPublishAPIRequest) GetRequest1() *ItemPublishRequest {
	return r._request1
}

var poolAlibabaAlihealthMedicalItemPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalItemPublishRequest()
	},
}

// GetAlibabaAlihealthMedicalItemPublishRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalItemPublishAPIRequest
func GetAlibabaAlihealthMedicalItemPublishAPIRequest() *AlibabaAlihealthMedicalItemPublishAPIRequest {
	return poolAlibabaAlihealthMedicalItemPublishAPIRequest.Get().(*AlibabaAlihealthMedicalItemPublishAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalItemPublishAPIRequest 将 AlibabaAlihealthMedicalItemPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalItemPublishAPIRequest(v *AlibabaAlihealthMedicalItemPublishAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalItemPublishAPIRequest.Put(v)
}
