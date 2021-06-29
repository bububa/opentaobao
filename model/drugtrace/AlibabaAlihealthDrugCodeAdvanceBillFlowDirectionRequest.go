package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据流向查询 API请求
alibaba.alihealth.drug.code.advance.bill.flow.direction

单据流向查询
*/
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest struct {
    model.Params
    // 追溯码
    _code   string
}

// 初始化AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest对象
func NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest() *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest{
    return &AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.advance.bill.flow.direction"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) GetCode() string {
    return r._code
}
