package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthstorecertificatecreateAPIRequest 仓库换证审批 API请求
// alibaba.alihealth.store.certificate.create
//
// 仓库侧换证发起审批
type AlibabaalihealthstorecertificatecreateAPIRequest struct {
	model.Params
	// 仓库code
	_storeCode string
	// 审批业务类型
	_auditType string
	// 审批内容
	_content string
	// 业务单号
	_bizNo string
}

// NewAlibabaalihealthstorecertificatecreateRequest 初始化AlibabaalihealthstorecertificatecreateAPIRequest对象
func NewAlibabaalihealthstorecertificatecreateRequest() *AlibabaalihealthstorecertificatecreateAPIRequest {
	return &AlibabaalihealthstorecertificatecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.store.certificate.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 仓库code
func (r *AlibabaalihealthstorecertificatecreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetAuditType is AuditType Setter
// 审批业务类型
func (r *AlibabaalihealthstorecertificatecreateAPIRequest) SetAuditType(_auditType string) error {
	r._auditType = _auditType
	r.Set("audit_type", _auditType)
	return nil
}

// GetAuditType AuditType Getter
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetAuditType() string {
	return r._auditType
}

// SetContent is Content Setter
// 审批内容
func (r *AlibabaalihealthstorecertificatecreateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetContent() string {
	return r._content
}

// SetBizNo is BizNo Setter
// 业务单号
func (r *AlibabaalihealthstorecertificatecreateAPIRequest) SetBizNo(_bizNo string) error {
	r._bizNo = _bizNo
	r.Set("biz_no", _bizNo)
	return nil
}

// GetBizNo BizNo Getter
func (r AlibabaalihealthstorecertificatecreateAPIRequest) GetBizNo() string {
	return r._bizNo
}
