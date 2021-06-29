package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
往来单位查询 
alibaba.alihealth.drugtrace.top.yljg.query.listparts

查询往来单位列表
*/
func AlibabaAlihealthDrugtraceTopYljgQueryListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryListpartsRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryListpartsAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryListpartsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
