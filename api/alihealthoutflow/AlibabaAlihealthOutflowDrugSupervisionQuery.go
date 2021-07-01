package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
监管平台药品查询 
alibaba.alihealth.outflow.drug.supervision.query

获取监管平台药品数据
*/
func AlibabaAlihealthOutflowDrugSupervisionQuery(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
