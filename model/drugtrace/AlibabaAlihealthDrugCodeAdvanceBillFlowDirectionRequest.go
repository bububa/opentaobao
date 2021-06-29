package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据流向查询 APIRequest
alibaba.alihealth.drug.code.advance.bill.flow.direction

单据流向查询
*/
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest struct {
    model.Params

    // 追溯码
    code   string 

}

func NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest() *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest{
    return &AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.advance.bill.flow.direction"
}

func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest) GetCode() string {
    return r.code
}

