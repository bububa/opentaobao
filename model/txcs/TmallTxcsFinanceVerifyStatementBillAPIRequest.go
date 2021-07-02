package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceVerifyStatementBillAPIRequest 供应商核销单录入 API请求
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
type TmallTxcsFinanceVerifyStatementBillAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 核销单内容
	_verificationBillDTO *VerificationBillDto
}

// NewTmallTxcsFinanceVerifyStatementBillRequest 初始化TmallTxcsFinanceVerifyStatementBillAPIRequest对象
func NewTmallTxcsFinanceVerifyStatementBillRequest() *TmallTxcsFinanceVerifyStatementBillAPIRequest {
	return &TmallTxcsFinanceVerifyStatementBillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.verify.statement.bill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuCode is OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceVerifyStatementBillAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetVerificationBillDTO is VerificationBillDTO Setter
// 核销单内容
func (r *TmallTxcsFinanceVerifyStatementBillAPIRequest) SetVerificationBillDTO(_verificationBillDTO *VerificationBillDto) error {
	r._verificationBillDTO = _verificationBillDTO
	r.Set("verification_bill_d_t_o", _verificationBillDTO)
	return nil
}

// GetVerificationBillDTO VerificationBillDTO Getter
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetVerificationBillDTO() *VerificationBillDto {
	return r._verificationBillDTO
}
