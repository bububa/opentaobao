package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceDeductGetAPIRequest 发票扣减的接口 API请求
// alibaba.einvoice.deduct.get
//
// 获取历史发票扣减量、每日发票扣减量的接口
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceDeductGetAPIRequest) Reset() {
	r._payeeRegisterNo = ""
	r._bizDate = ""
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.deduct.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceDeductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceDeductGetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceDeductGetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetBizDate is BizDate Setter
// 业务日期
func (r *AlibabaEinvoiceDeductGetAPIRequest) SetBizDate(_bizDate string) error {
	r._bizDate = _bizDate
	r.Set("biz_date", _bizDate)
	return nil
}

// GetBizDate BizDate Getter
func (r AlibabaEinvoiceDeductGetAPIRequest) GetBizDate() string {
	return r._bizDate
}

// SetType is Type Setter
// 类型 1：所有 2：当日
func (r *AlibabaEinvoiceDeductGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaEinvoiceDeductGetAPIRequest) GetType() int64 {
	return r._type
}

var poolAlibabaEinvoiceDeductGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceDeductGetRequest()
	},
}

// GetAlibabaEinvoiceDeductGetRequest 从 sync.Pool 获取 AlibabaEinvoiceDeductGetAPIRequest
func GetAlibabaEinvoiceDeductGetAPIRequest() *AlibabaEinvoiceDeductGetAPIRequest {
	return poolAlibabaEinvoiceDeductGetAPIRequest.Get().(*AlibabaEinvoiceDeductGetAPIRequest)
}

// ReleaseAlibabaEinvoiceDeductGetAPIRequest 将 AlibabaEinvoiceDeductGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceDeductGetAPIRequest(v *AlibabaEinvoiceDeductGetAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceDeductGetAPIRequest.Put(v)
}
