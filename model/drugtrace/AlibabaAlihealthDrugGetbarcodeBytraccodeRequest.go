package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据追溯码获取69码 API请求
alibaba.alihealth.drug.getbarcode.bytraccode

根据追溯码获取69码
*/
type AlibabaAlihealthDrugGetbarcodeBytraccodeRequest struct {
    model.Params
    // 追溯码
    traceCode   string
}

// 初始化AlibabaAlihealthDrugGetbarcodeBytraccodeRequest对象
func NewAlibabaAlihealthDrugGetbarcodeBytraccodeRequest() *AlibabaAlihealthDrugGetbarcodeBytraccodeRequest{
    return &AlibabaAlihealthDrugGetbarcodeBytraccodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.getbarcode.bytraccode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceCode Setter
// 追溯码
func (r *AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) SetTraceCode(traceCode string) error {
    r.traceCode = traceCode
    r.Set("trace_code", traceCode)
    return nil
}

// TraceCode Getter
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) GetTraceCode() string {
    return r.traceCode
}
