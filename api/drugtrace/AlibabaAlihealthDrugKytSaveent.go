package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
新增往来单位企业 
alibaba.alihealth.drug.kyt.saveent

新增往来单位企业记录
*/
func AlibabaAlihealthDrugKytSaveent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSaveentAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSaveentAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytSaveentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
