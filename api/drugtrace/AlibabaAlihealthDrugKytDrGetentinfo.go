package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) 
alibaba.alihealth.drug.kyt.dr.getentinfo

根据企业名称查询ID
*/
func AlibabaAlihealthDrugKytDrGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetentinfoRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrGetentinfoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDrGetentinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
