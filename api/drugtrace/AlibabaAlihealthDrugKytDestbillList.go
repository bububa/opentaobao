package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
直调单据查询 
alibaba.alihealth.drug.kyt.destbill.list

为药企提供直调单据查询功能
*/
func AlibabaAlihealthDrugKytDestbillList(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDestbillListRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDestbillListAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDestbillListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
