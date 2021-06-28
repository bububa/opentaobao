package alihealth

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth"
)

/* 
阿里健康处方平台获取授权码 
alibaba.alihealth.prescription.auth.get

获取处方用户授权
*/
func AlibabaAlihealthPrescriptionAuthGet(clt *core.SDKClient, req *alihealth.AlibabaAlihealthPrescriptionAuthGetRequest, session string) (*alihealth.AlibabaAlihealthPrescriptionAuthGetAPIResponse, error) {
    var resp alihealth.AlibabaAlihealthPrescriptionAuthGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
