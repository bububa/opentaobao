package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
上游出库单单据明细查询 
alibaba.alihealth.drugtrace.top.yljg.listupout.detail

查询上游出库单明细(带追溯码信息)
*/
func AlibabaAlihealthDrugtraceTopYljgListupoutDetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
