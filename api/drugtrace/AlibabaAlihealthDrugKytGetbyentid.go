package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
根据企业主键查看企业详细信息 
alibaba.alihealth.drug.kyt.getbyentid

根据企业主键查看企业详细信息
*/
func AlibabaAlihealthDrugKytGetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetbyentidRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytGetbyentidAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytGetbyentidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
