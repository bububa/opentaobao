package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
仓库换证审批 
alibaba.alihealth.store.certificate.create

仓库侧换证发起审批
*/
func AlibabaAlihealthStoreCertificateCreate(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthStoreCertificateCreateRequest, session string) (*alihealth2.AlibabaAlihealthStoreCertificateCreateAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthStoreCertificateCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
