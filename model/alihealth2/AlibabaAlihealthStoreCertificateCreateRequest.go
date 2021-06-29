package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库换证审批 APIRequest
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

func NewAlibabaAlihealthStoreCertificateCreateRequest() *AlibabaAlihealthStoreCertificateCreateRequest{
    return &AlibabaAlihealthStoreCertificateCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthStoreCertificateCreateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.store.certificate.create"
}

func (r AlibabaAlihealthStoreCertificateCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r AlibabaAlihealthStoreCertificateCreateRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetAuditType(auditType string) error {
    r.auditType = auditType
    r.Set("audit_type", auditType)
    return nil
}

func (r AlibabaAlihealthStoreCertificateCreateRequest) GetAuditType() string {
    return r.auditType
}

func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaAlihealthStoreCertificateCreateRequest) GetContent() string {
    return r.content
}

func (r *AlibabaAlihealthStoreCertificateCreateRequest) SetBizNo(bizNo string) error {
    r.bizNo = bizNo
    r.Set("biz_no", bizNo)
    return nil
}

func (r AlibabaAlihealthStoreCertificateCreateRequest) GetBizNo() string {
    return r.bizNo
}

