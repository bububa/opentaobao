package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceBillConfirmAPIRequest 供应商账单确认 API请求
// tmall.txcs.finance.bill.confirm
//
// 提供天猫超市外部合作商家：财务账单对账
type TmallTxcsFinanceBillConfirmAPIRequest struct {
	model.Params
	// 门店ID
	_ouCode string
	// 系统自动生成
	_statementBillConfirmDTO *StatementBillConfirmDto
}

// NewTmallTxcsFinanceBillConfirmRequest 初始化TmallTxcsFinanceBillConfirmAPIRequest对象
func NewTmallTxcsFinanceBillConfirmRequest() *TmallTxcsFinanceBillConfirmAPIRequest {
	return &TmallTxcsFinanceBillConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.bill.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuCode is OuCode Setter
// 门店ID
func (r *TmallTxcsFinanceBillConfirmAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetStatementBillConfirmDTO is StatementBillConfirmDTO Setter
// 系统自动生成
func (r *TmallTxcsFinanceBillConfirmAPIRequest) SetStatementBillConfirmDTO(_statementBillConfirmDTO *StatementBillConfirmDto) error {
	r._statementBillConfirmDTO = _statementBillConfirmDTO
	r.Set("statement_bill_confirm_d_t_o", _statementBillConfirmDTO)
	return nil
}

// GetStatementBillConfirmDTO StatementBillConfirmDTO Getter
func (r TmallTxcsFinanceBillConfirmAPIRequest) GetStatementBillConfirmDTO() *StatementBillConfirmDto {
	return r._statementBillConfirmDTO
}
