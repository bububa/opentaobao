package txcs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltxcsfinancebillcheckAPIRequest 天猫超市外部商家财务账单对账 API请求
// tmall.txcs.finance.bill.check
//
// 提供天猫超市外部合作商家财务账单对账
type TmalltxcsfinancebillcheckAPIRequest struct {
	model.Params
	// 门店编码
	_ouCode string
	// 查询参数
	_statementBillFeeDetailQuery *StatementBillFeeDetailQuery
}

// NewTmalltxcsfinancebillcheckRequest 初始化TmalltxcsfinancebillcheckAPIRequest对象
func NewTmalltxcsfinancebillcheckRequest() *TmalltxcsfinancebillcheckAPIRequest {
	return &TmalltxcsfinancebillcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltxcsfinancebillcheckAPIRequest) GetApiMethodName() string {
	return "tmall.txcs.finance.bill.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltxcsfinancebillcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltxcsfinancebillcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuCode is OuCode Setter
// 门店编码
func (r *TmalltxcsfinancebillcheckAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r TmalltxcsfinancebillcheckAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetStatementBillFeeDetailQuery is StatementBillFeeDetailQuery Setter
// 查询参数
func (r *TmalltxcsfinancebillcheckAPIRequest) SetStatementBillFeeDetailQuery(_statementBillFeeDetailQuery *StatementBillFeeDetailQuery) error {
	r._statementBillFeeDetailQuery = _statementBillFeeDetailQuery
	r.Set("statement_bill_fee_detail_query", _statementBillFeeDetailQuery)
	return nil
}

// GetStatementBillFeeDetailQuery StatementBillFeeDetailQuery Getter
func (r TmalltxcsfinancebillcheckAPIRequest) GetStatementBillFeeDetailQuery() *StatementBillFeeDetailQuery {
	return r._statementBillFeeDetailQuery
}
