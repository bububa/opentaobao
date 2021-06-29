package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票列表 API请求
alibaba.einvoice.bill.einvoice.list

扫码开票列表，包括用户扫二维码开票和结算单同步前的开票数据
*/
type AlibabaEinvoiceBillEinvoiceListRequest struct {
    model.Params
    // 结算单同步的ERP平台系统
    _platform   string
    // 收款方税号
    _payeeRegisterNo   string
    // 订单ID
    _orderId   string
    // 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
    _einvoiceType   []int64
}

// 初始化AlibabaEinvoiceBillEinvoiceListRequest对象
func NewAlibabaEinvoiceBillEinvoiceListRequest() *AlibabaEinvoiceBillEinvoiceListRequest{
    return &AlibabaEinvoiceBillEinvoiceListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetApiMethodName() string {
    return "alibaba.einvoice.bill.einvoice.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Platform Setter
// 结算单同步的ERP平台系统
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetPlatform() string {
    return r._platform
}
// PayeeRegisterNo Setter
// 收款方税号
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// OrderId Setter
// 订单ID
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetOrderId() string {
    return r._orderId
}
// EinvoiceType Setter
// 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetEinvoiceType(_einvoiceType []int64) error {
    r._einvoiceType = _einvoiceType
    r.Set("einvoice_type", _einvoiceType)
    return nil
}

// EinvoiceType Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetEinvoiceType() []int64 {
    return r._einvoiceType
}
