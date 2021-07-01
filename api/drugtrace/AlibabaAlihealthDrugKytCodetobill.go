package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
通过追溯码查单据 
alibaba.alihealth.drug.kyt.codetobill

通过追溯码查单据
*/
func AlibabaAlihealthDrugKytCodetobill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodetobillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytCodetobillAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytCodetobillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
