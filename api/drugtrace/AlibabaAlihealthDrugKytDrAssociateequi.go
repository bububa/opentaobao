package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
疫苗单据与设备绑定 
alibaba.alihealth.drug.kyt.dr.associateequi

疫苗单据与设备绑定
*/
func AlibabaAlihealthDrugKytDrAssociateequi(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrAssociateequiAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrAssociateequiAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDrAssociateequiAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
