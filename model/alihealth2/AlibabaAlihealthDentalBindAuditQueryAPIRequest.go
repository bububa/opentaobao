package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalBindAuditQueryAPIRequest
ISV查询绑定审核状态 API请求
alibaba.alihealth.dental.bind.audit.query

ISV查询绑定审核状态 */
type AlibabaAlihealthDentalBindAuditQueryAPIRequest struct {
	model.Params
	// 绑定ID列表
	_bindIds []int64
}

// NewAlibabaAlihealthDentalBindAuditQueryRequest 初始化AlibabaAlihealthDentalBindAuditQueryAPIRequest对象
func NewAlibabaAlihealthDentalBindAuditQueryRequest() *AlibabaAlihealthDentalBindAuditQueryAPIRequest {
	return &AlibabaAlihealthDentalBindAuditQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.bind.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BindIds Setter
// 绑定ID列表
func (r *AlibabaAlihealthDentalBindAuditQueryAPIRequest) SetBindIds(_bindIds []int64) error {
	r._bindIds = _bindIds
	r.Set("bind_ids", _bindIds)
	return nil
}

// Get BindIds Getter
func (r AlibabaAlihealthDentalBindAuditQueryAPIRequest) GetBindIds() []int64 {
	return r._bindIds
}
