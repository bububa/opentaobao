package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商订购单上传核对 APIRequest
alibaba.einvoice.unitorder.check

开票服务商回传收到的订购单用于电子发票平台核对
*/
type AlibabaEinvoiceUnitorderCheckRequest struct {
    model.Params

    // 订购单列表
    orders   []SimpleUnitOrder 

    // 开始时间,来自于查询消息
    begin   string 

    // 结束时间,来自于查询消息
    end   string 

}

func NewAlibabaEinvoiceUnitorderCheckRequest() *AlibabaEinvoiceUnitorderCheckRequest{
    return &AlibabaEinvoiceUnitorderCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceUnitorderCheckRequest) GetApiMethodName() string {
    return "alibaba.einvoice.unitorder.check"
}

func (r AlibabaEinvoiceUnitorderCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceUnitorderCheckRequest) SetOrders(orders []SimpleUnitOrder) error {
    r.orders = orders
    r.Set("orders", orders)
    return nil
}

func (r AlibabaEinvoiceUnitorderCheckRequest) GetOrders() []SimpleUnitOrder {
    return r.orders
}

func (r *AlibabaEinvoiceUnitorderCheckRequest) SetBegin(begin string) error {
    r.begin = begin
    r.Set("begin", begin)
    return nil
}

func (r AlibabaEinvoiceUnitorderCheckRequest) GetBegin() string {
    return r.begin
}

func (r *AlibabaEinvoiceUnitorderCheckRequest) SetEnd(end string) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r AlibabaEinvoiceUnitorderCheckRequest) GetEnd() string {
    return r.end
}

