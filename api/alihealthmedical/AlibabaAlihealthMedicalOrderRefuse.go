package alihealthmedical

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthmedical"
)

/* 
三方机构通知平台"医生拒诊" 
alibaba.alihealth.medical.order.refuse

三方机构通知平台"医生拒诊"
*/
func AlibabaAlihealthMedicalOrderRefuse(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalOrderRefuseRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIResponse, error) {
    var resp alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
