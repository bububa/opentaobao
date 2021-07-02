package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthStoreCertificateCreateAPIRequest 仓库换证审批 API请求
// alibaba.alihealth.store.certificate.create
//
// 仓库侧换证发起审批
type AlibabaAlihealthStoreCertificateCreateAPIRequest struct {
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

// NewAlibabaAlihealthStoreCertificateCreateRequest 初始化AlibabaAlihealthStoreCertificateCreateAPIRequest对象
func NewAlibabaAlihealthStoreCertificateCreateRequest() *AlibabaAlihealthStoreCertificateCreateAPIRequest {
	return &AlibabaAlihealthStoreCertificateCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthStoreCertificateCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.store.certificate.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthStoreCertificateCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreCode is StoreCode Setter
// 仓库code
func (r *AlibabaAlihealthStoreCertificateCreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabaAlihealthStoreCertificateCreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetAuditType is AuditType Setter
// 审批业务类型
func (r *AlibabaAlihealthStoreCertificateCreateAPIRequest) SetAuditType(_auditType string) error {
	r._auditType = _auditType
	r.Set("audit_type", _auditType)
	return nil
}

// GetAuditType AuditType Getter
func (r AlibabaAlihealthStoreCertificateCreateAPIRequest) GetAuditType() string {
	return r._auditType
}

// SetContent is Content Setter
// 审批内容
func (r *AlibabaAlihealthStoreCertificateCreateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaAlihealthStoreCertificateCreateAPIRequest) GetContent() string {
	return r._content
}

// SetBizNo is BizNo Setter
// 业务单号
func (r *AlibabaAlihealthStoreCertificateCreateAPIRequest) SetBizNo(_bizNo string) error {
	r._bizNo = _bizNo
	r.Set("biz_no", _bizNo)
	return nil
}

// GetBizNo BizNo Getter
func (r AlibabaAlihealthStoreCertificateCreateAPIRequest) GetBizNo() string {
	return r._bizNo
}
