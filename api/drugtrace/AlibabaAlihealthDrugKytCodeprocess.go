package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
关联关系处理查询 
alibaba.alihealth.drug.kyt.codeprocess

关联关系处理查询
*/
func AlibabaAlihealthDrugKytCodeprocess(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodeprocessRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytCodeprocessAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytCodeprocessAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
