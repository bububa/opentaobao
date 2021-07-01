package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查询药品是否赋码 
alibaba.alihealth.drug.kyt.yy.drugcodes

药品是否赋码
*/
func AlibabaAlihealthDrugKytYyDrugcodes(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyDrugcodesAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytYyDrugcodesAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytYyDrugcodesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
