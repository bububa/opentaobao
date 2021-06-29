package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
药品是否赋码 
alibaba.alihealth.drug.kyt.drugcodes

药品是否赋码
*/
func AlibabaAlihealthDrugKytDrugcodes(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugcodesRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrugcodesAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDrugcodesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
