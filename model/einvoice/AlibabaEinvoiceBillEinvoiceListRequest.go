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
    platform   string
    // 收款方税号
    payeeRegisterNo   string
    // 订单ID
    orderId   string
    // 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
    einvoiceType   []int64
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
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetPlatform() string {
    return r.platform
}
// PayeeRegisterNo Setter
// 收款方税号
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// OrderId Setter
// 订单ID
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetOrderId() string {
    return r.orderId
}
// EinvoiceType Setter
// 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
func (r *AlibabaEinvoiceBillEinvoiceListRequest) SetEinvoiceType(einvoiceType []int64) error {
    r.einvoiceType = einvoiceType
    r.Set("einvoice_type", einvoiceType)
    return nil
}

// EinvoiceType Getter
func (r AlibabaEinvoiceBillEinvoiceListRequest) GetEinvoiceType() []int64 {
    return r.einvoiceType
}
