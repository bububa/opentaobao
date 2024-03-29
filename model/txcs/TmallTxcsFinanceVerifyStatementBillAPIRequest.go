package txcs

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTxcsFinanceVerifyStatementBillAPIRequest) Reset() {
	r._ouCode = ""
	r._verificationBillDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.verify.statement.bill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTxcsFinanceVerifyStatementBillAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallTxcsFinanceVerifyStatementBillAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTxcsFinanceVerifyStatementBillRequest()
	},
}

// GetTmallTxcsFinanceVerifyStatementBillRequest 从 sync.Pool 获取 TmallTxcsFinanceVerifyStatementBillAPIRequest
func GetTmallTxcsFinanceVerifyStatementBillAPIRequest() *TmallTxcsFinanceVerifyStatementBillAPIRequest {
	return poolTmallTxcsFinanceVerifyStatementBillAPIRequest.Get().(*TmallTxcsFinanceVerifyStatementBillAPIRequest)
}

// ReleaseTmallTxcsFinanceVerifyStatementBillAPIRequest 将 TmallTxcsFinanceVerifyStatementBillAPIRequest 放入 sync.Pool
func ReleaseTmallTxcsFinanceVerifyStatementBillAPIRequest(v *TmallTxcsFinanceVerifyStatementBillAPIRequest) {
	v.Reset()
	poolTmallTxcsFinanceVerifyStatementBillAPIRequest.Put(v)
}
