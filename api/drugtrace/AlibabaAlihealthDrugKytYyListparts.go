package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查询往来单位 
alibaba.alihealth.drug.kyt.yy.listparts

查询往来单位列表
*/
func AlibabaAlihealthDrugKytYyListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyListpartsRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytYyListpartsAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytYyListpartsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
