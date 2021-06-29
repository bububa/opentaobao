package tax

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方开票回调API APIRequest
alibaba.tax.invoice.sync

该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
*/
type AlibabaTaxInvoiceSyncRequest struct {
    model.Params

    // 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
    paramJson   string 

}

func NewAlibabaTaxInvoiceSyncRequest() *AlibabaTaxInvoiceSyncRequest{
    return &AlibabaTaxInvoiceSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTaxInvoiceSyncRequest) GetApiMethodName() string {
    return "alibaba.tax.invoice.sync"
}

func (r AlibabaTaxInvoiceSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTaxInvoiceSyncRequest) SetParamJson(paramJson string) error {
    r.paramJson = paramJson
    r.Set("param_json", paramJson)
    return nil
}

func (r AlibabaTaxInvoiceSyncRequest) GetParamJson() string {
    return r.paramJson
}

