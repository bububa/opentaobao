package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票流水号批量生成接口 APIRequest
alibaba.einvoice.serialno.batch.generate

批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。
*/
type AlibabaEinvoiceSerialnoBatchGenerateRequest struct {
    model.Params

}

func NewAlibabaEinvoiceSerialnoBatchGenerateRequest() *AlibabaEinvoiceSerialnoBatchGenerateRequest{
    return &AlibabaEinvoiceSerialnoBatchGenerateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceSerialnoBatchGenerateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.serialno.batch.generate"
}

func (r AlibabaEinvoiceSerialnoBatchGenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


