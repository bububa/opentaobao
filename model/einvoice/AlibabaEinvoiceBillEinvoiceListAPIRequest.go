package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceBillEinvoiceListAPIRequest 扫码开票列表 API请求
// alibaba.einvoice.bill.einvoice.list
//
// 扫码开票列表，包括用户扫二维码开票和结算单同步前的开票数据
type AlibabaEinvoiceBillEinvoiceListAPIRequest struct {
	model.Params
	// 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
	_einvoiceType []string
	// 结算单同步的ERP平台系统
	_platform string
	// 收款方税号
	_payeeRegisterNo string
	// 订单ID
	_orderId string
}

// NewAlibabaEinvoiceBillEinvoiceListRequest 初始化AlibabaEinvoiceBillEinvoiceListAPIRequest对象
func NewAlibabaEinvoiceBillEinvoiceListRequest() *AlibabaEinvoiceBillEinvoiceListAPIRequest {
	return &AlibabaEinvoiceBillEinvoiceListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.bill.einvoice.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEinvoiceType is EinvoiceType Setter
// 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
func (r *AlibabaEinvoiceBillEinvoiceListAPIRequest) SetEinvoiceType(_einvoiceType []string) error {
	r._einvoiceType = _einvoiceType
	r.Set("einvoice_type", _einvoiceType)
	return nil
}

// GetEinvoiceType EinvoiceType Getter
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetEinvoiceType() []string {
	return r._einvoiceType
}

// SetPlatform is Platform Setter
// 结算单同步的ERP平台系统
func (r *AlibabaEinvoiceBillEinvoiceListAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetPlatform() string {
	return r._platform
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税号
func (r *AlibabaEinvoiceBillEinvoiceListAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabaEinvoiceBillEinvoiceListAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEinvoiceBillEinvoiceListAPIRequest) GetOrderId() string {
	return r._orderId
}
