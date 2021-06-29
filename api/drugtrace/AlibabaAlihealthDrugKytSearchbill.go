package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
通过时间段批量查询入出库单信息 
alibaba.alihealth.drug.kyt.searchbill

通过时间段批量查询入出库单信息
*/
func AlibabaAlihealthDrugKytSearchbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSearchbillRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSearchbillAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytSearchbillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
