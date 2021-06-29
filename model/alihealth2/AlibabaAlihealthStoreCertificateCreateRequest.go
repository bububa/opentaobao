package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库换证审批 API请求
alibaba.alihealth.store.certificate.create

仓库侧换证发起审批
*/
type AlibabaAlihealthStoreCertificateCreateRequest struct {
    model.Params
    // 仓库code
    storeCode   string
    // 审批业务类型
    auditType   string
    // 审批内容
    content   string
    // 业务单号
    bizNo   string
}

// 初始化AlibabaAlihealthStoreCertificateCreateRequest对象
func NewAlibabaAlihealthStoreCertificateCreateRequest() *AlibabaAlihealthStoreCertificateCreateRequest{
    return &AlibabaAlihealthStoreCertificateCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthStoreCertificateCreateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.store.certificate.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthStoreCertificateCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓库code
func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaAlihealthStoreCertificateCreateRequest) GetStoreCode() string {
    return r.storeCode
}
// AuditType Setter
// 审批业务类型
func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetAuditType(auditType string) error {
    r.auditType = auditType
    r.Set("audit_type", auditType)
    return nil
}

// AuditType Getter
func (r AlibabaAlihealthStoreCertificateCreateRequest) GetAuditType() string {
    return r.auditType
}
// Content Setter
// 审批内容
func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaAlihealthStoreCertificateCreateRequest) GetContent() string {
    return r.content
}
// BizNo Setter
// 业务单号
func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetBizNo(bizNo string) error {
    r.bizNo = bizNo
    r.Set("biz_no", bizNo)
    return nil
}

// BizNo Getter
func (r AlibabaAlihealthStoreCertificateCreateRequest) GetBizNo() string {
    return r.bizNo
}
