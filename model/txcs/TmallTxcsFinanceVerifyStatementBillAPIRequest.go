package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltxcsfinanceverifystatementbillAPIRequest 供应商核销单录入 API请求
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
type TmalltxcsfinanceverifystatementbillAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 核销单内容
	_verificationBillDTO *VerificationBillDto
}

// NewTmalltxcsfinanceverifystatementbillRequest 初始化TmalltxcsfinanceverifystatementbillAPIRequest对象
func NewTmalltxcsfinanceverifystatementbillRequest() *TmalltxcsfinanceverifystatementbillAPIRequest {
	return &TmalltxcsfinanceverifystatementbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltxcsfinanceverifystatementbillAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.verify.statement.bill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltxcsfinanceverifystatementbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltxcsfinanceverifystatementbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuCode is OuCode Setter
// 门店ID
func (r *TmalltxcsfinanceverifystatementbillAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmalltxcsfinanceverifystatementbillAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetVerificationBillDTO is VerificationBillDTO Setter
// 核销单内容
func (r *TmalltxcsfinanceverifystatementbillAPIRequest) SetVerificationBillDTO(_verificationBillDTO *VerificationBillDto) error {
	r._verificationBillDTO = _verificationBillDTO
	r.Set("verification_bill_d_t_o", _verificationBillDTO)
	return nil
}

// GetVerificationBillDTO VerificationBillDTO Getter
func (r TmalltxcsfinanceverifystatementbillAPIRequest) GetVerificationBillDTO() *VerificationBillDto {
	return r._verificationBillDTO
}
