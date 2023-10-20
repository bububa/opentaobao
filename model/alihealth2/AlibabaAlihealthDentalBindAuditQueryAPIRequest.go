package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalBindAuditQueryAPIRequest ISV查询绑定审核状态 API请求
// alibaba.alihealth.dental.bind.audit.query
//
// ISV查询绑定审核状态
type AlibabaAlihealthDentalBindAuditQueryAPIRequest struct {
	model.Params
	// 绑定ID列表
	_bindIds []string
}

// NewAlibabaAlihealthDentalBindAuditQueryRequest 初始化AlibabaAlihealthDentalBindAuditQueryAPIRequest对象
func NewAlibabaAlihealthDentalBindAuditQueryRequest() *AlibabaAlihealthDentalBindAuditQueryAPIRequest {
	return &AlibabaAlihealthDentalBindAuditQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDentalBindAuditQueryAPIRequest) Reset() {
	r._bindIds = r._bindIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.bind.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindIds is BindIds Setter
// 绑定ID列表
func (r *AlibabaAlihealthDentalBindAuditQueryAPIRequest) SetBindIds(_bindIds []string) error {
	r._bindIds = _bindIds
	r.Set("bind_ids", _bindIds)
	return nil
}

// GetBindIds BindIds Getter
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetBindIds() []string {
	return r._bindIds
}

var poolAlibabaAlihealthDentalBindAuditQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDentalBindAuditQueryRequest()
	},
}

// GetAlibabaAlihealthDentalBindAuditQueryRequest 从 sync.Pool 获取 AlibabaAlihealthDentalBindAuditQueryAPIRequest
func GetAlibabaAlihealthDentalBindAuditQueryAPIRequest() *AlibabaAlihealthDentalBindAuditQueryAPIRequest {
	return poolAlibabaAlihealthDentalBindAuditQueryAPIRequest.Get().(*AlibabaAlihealthDentalBindAuditQueryAPIRequest)
}

// ReleaseAlibabaAlihealthDentalBindAuditQueryAPIRequest 将 AlibabaAlihealthDentalBindAuditQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDentalBindAuditQueryAPIRequest(v *AlibabaAlihealthDentalBindAuditQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDentalBindAuditQueryAPIRequest.Put(v)
}
