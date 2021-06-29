package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 
alibaba.alihealth.drug.kyt.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
func AlibabaAlihealthDrugKytGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetentinfoRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytGetentinfoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytGetentinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
