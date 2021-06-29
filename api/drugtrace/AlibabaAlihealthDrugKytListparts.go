package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查询往来单位列表 
alibaba.alihealth.drug.kyt.listparts

查询往来单位列表
*/
func AlibabaAlihealthDrugKytListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListpartsRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytListpartsAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytListpartsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
