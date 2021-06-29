package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
通过企业名得到唯一标识ref_ent_id及企业ent_id 
alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo

根据企业名称查询ID
*/
func AlibabaAlihealthDrugtraceTopLsydQueryGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
