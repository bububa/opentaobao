package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceDeductGetAPIRequest
发票扣减的接口 API请求
alibaba.einvoice.deduct.get

获取历史发票扣减量、每日发票扣减量的接口 */
type AlibabaEinvoiceDeductGetAPIRequest struct {
	model.Params
	// 税号
	_payeeRegisterNo string
	// 业务日期
	_bizDate string
	// 类型 1：所有 2：当日
	_type int64
}

// NewAlibabaEinvoiceDeductGetRequest 初始化AlibabaEinvoiceDeductGetAPIRequest对象
func NewAlibabaEinvoiceDeductGetRequest() *AlibabaEinvoiceDeductGetAPIRequest {
	return &AlibabaEinvoiceDeductGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.deduct.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeductGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceDeductGetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// Get PayeeRegisterNo Getter
func (r AlibabaEinvoiceDeductGetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// Set is BizDate Setter
// 业务日期
func (r *AlibabaEinvoiceDeductGetAPIRequest) SetBizDate(_bizDate string) error {
	r._bizDate = _bizDate
	r.Set("biz_date", _bizDate)
	return nil
}

// Get BizDate Getter
func (r AlibabaEinvoiceDeductGetAPIRequest) GetBizDate() string {
	return r._bizDate
}

// Set is Type Setter
// 类型 1：所有 2：当日
func (r *AlibabaEinvoiceDeductGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaEinvoiceDeductGetAPIRequest) GetType() int64 {
	return r._type
}
