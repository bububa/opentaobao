package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
监管平台药品查询 API请求
alibaba.alihealth.outflow.drug.supervision.query

获取监管平台药品数据
*/
type AlibabaAlihealthOutflowDrugSupervisionQueryRequest struct {
    model.Params
    // 请求
    request1   *OuterDrugVo
}

// 初始化AlibabaAlihealthOutflowDrugSupervisionQueryRequest对象
func NewAlibabaAlihealthOutflowDrugSupervisionQueryRequest() *AlibabaAlihealthOutflowDrugSupervisionQueryRequest{
    return &AlibabaAlihealthOutflowDrugSupervisionQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowDrugSupervisionQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.drug.supervision.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowDrugSupervisionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request1 Setter
// 请求
func (r *AlibabaAlihealthOutflowDrugSupervisionQueryRequest) SetRequest1(request1 *OuterDrugVo) error {
    r.request1 = request1
    r.Set("request1", request1)
    return nil
}

// Request1 Getter
func (r AlibabaAlihealthOutflowDrugSupervisionQueryRequest) GetRequest1() *OuterDrugVo {
    return r.request1
}
