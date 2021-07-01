package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
查企业标识信息 
alibaba.alihealth.drug.kyt.smyx.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
func AlibabaAlihealthDrugKytSmyxGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytSmyxGetentinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
