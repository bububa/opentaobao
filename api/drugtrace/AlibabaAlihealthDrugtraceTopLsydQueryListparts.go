package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
往来单位查询 
alibaba.alihealth.drugtrace.top.lsyd.query.listparts

查询往来单位列表
*/
func AlibabaAlihealthDrugtraceTopLsydQueryListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
