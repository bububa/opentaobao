package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
直调审批 
alibaba.alihealth.drug.kyt.destbill.check

为药企提供直调单据审批操作
*/
func AlibabaAlihealthDrugKytDestbillCheck(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDestbillCheckAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDestbillCheckAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDestbillCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
