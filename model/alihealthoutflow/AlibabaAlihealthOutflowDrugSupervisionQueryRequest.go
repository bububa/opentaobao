package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
监管平台药品查询 APIRequest
alibaba.alihealth.outflow.drug.supervision.query

获取监管平台药品数据
*/
type AlibabaAlihealthOutflowDrugSupervisionQueryRequest struct {
    model.Params

    // 请求
    request1   *OuterDrugVo 

}

func NewAlibabaAlihealthOutflowDrugSupervisionQueryRequest() *AlibabaAlihealthOutflowDrugSupervisionQueryRequest{
    return &AlibabaAlihealthOutflowDrugSupervisionQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowDrugSupervisionQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.drug.supervision.query"
}

func (r AlibabaAlihealthOutflowDrugSupervisionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowDrugSupervisionQueryRequest) SetRequest1(request1 *OuterDrugVo) error {
    r.request1 = request1
    r.Set("request1", request1)
    return nil
}

func (r AlibabaAlihealthOutflowDrugSupervisionQueryRequest) GetRequest1() *OuterDrugVo {
    return r.request1
}

