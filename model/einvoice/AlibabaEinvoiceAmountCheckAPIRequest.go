package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceamountcheckAPIRequest 开票量核对接口 API请求
// alibaba.einvoice.amount.check
//
// 跟开票服务商核对历史开票量，用来对账
type AlibabaeinvoiceamountcheckAPIRequest struct {
	model.Params
	// 税号
	_payeeRegisterNo string
	// 开票日期开始时间
	_startDate string
	// 开票日期结束时间
	_endDate string
}

// NewAlibabaeinvoiceamountcheckRequest 初始化AlibabaeinvoiceamountcheckAPIRequest对象
func NewAlibabaeinvoiceamountcheckRequest() *AlibabaeinvoiceamountcheckAPIRequest {
	return &AlibabaeinvoiceamountcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceamountcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.amount.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceamountcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceamountcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税号
func (r *AlibabaeinvoiceamountcheckAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoiceamountcheckAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetStartDate is StartDate Setter
// 开票日期开始时间
func (r *AlibabaeinvoiceamountcheckAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaeinvoiceamountcheckAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 开票日期结束时间
func (r *AlibabaeinvoiceamountcheckAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaeinvoiceamountcheckAPIRequest) GetEndDate() string {
	return r._endDate
}
