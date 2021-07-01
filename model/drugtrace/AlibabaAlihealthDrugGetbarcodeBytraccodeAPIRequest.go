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
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest struct {
    model.Params
    // 追溯码
    _traceCode   string
}

// 初始化AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest对象
func NewAlibabaAlihealthDrugGetbarcodeBytraccodeRequest() *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest{
    return &AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.getbarcode.bytraccode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceCode Setter
// 追溯码
func (r *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) SetTraceCode(_traceCode string) error {
    r._traceCode = _traceCode
    r.Set("trace_code", _traceCode)
    return nil
}

// TraceCode Getter
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetTraceCode() string {
    return r._traceCode
}
