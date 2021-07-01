package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
药店查询往来单位 
alibaba.alihealth.drug.kyt.smyx.listparts

查询往来单位列表
*/
func AlibabaAlihealthDrugKytSmyxListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSmyxListpartsAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSmyxListpartsAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytSmyxListpartsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
