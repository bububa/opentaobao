package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询出入库单处理结果api 
alibaba.alihealth.tracecodeseller.bill.result.search

查询出入库单处理结果api
*/
func AlibabaAlihealthTracecodesellerBillResultSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
