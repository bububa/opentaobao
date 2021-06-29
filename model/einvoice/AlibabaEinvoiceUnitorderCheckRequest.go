package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商订购单上传核对 API请求
alibaba.einvoice.unitorder.check

开票服务商回传收到的订购单用于电子发票平台核对
*/
type AlibabaEinvoiceUnitorderCheckRequest struct {
    model.Params
    // 订购单列表
    _orders   []SimpleUnitOrder
    // 开始时间,来自于查询消息
    _begin   string
    // 结束时间,来自于查询消息
    _end   string
}

// 初始化AlibabaEinvoiceUnitorderCheckRequest对象
func NewAlibabaEinvoiceUnitorderCheckRequest() *AlibabaEinvoiceUnitorderCheckRequest{
    return &AlibabaEinvoiceUnitorderCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceUnitorderCheckRequest) GetApiMethodName() string {
    return "alibaba.einvoice.unitorder.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceUnitorderCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Orders Setter
// 订购单列表
func (r *AlibabaEinvoiceUnitorderCheckRequest) SetOrders(_orders []SimpleUnitOrder) error {
    r._orders = _orders
    r.Set("orders", _orders)
    return nil
}

// Orders Getter
func (r AlibabaEinvoiceUnitorderCheckRequest) GetOrders() []SimpleUnitOrder {
    return r._orders
}
// Begin Setter
// 开始时间,来自于查询消息
func (r *AlibabaEinvoiceUnitorderCheckRequest) SetBegin(_begin string) error {
    r._begin = _begin
    r.Set("begin", _begin)
    return nil
}

// Begin Getter
func (r AlibabaEinvoiceUnitorderCheckRequest) GetBegin() string {
    return r._begin
}
// End Setter
// 结束时间,来自于查询消息
func (r *AlibabaEinvoiceUnitorderCheckRequest) SetEnd(_end string) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r AlibabaEinvoiceUnitorderCheckRequest) GetEnd() string {
    return r._end
}
