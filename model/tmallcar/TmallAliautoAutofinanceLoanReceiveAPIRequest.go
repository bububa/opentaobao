package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoAutofinanceLoanReceiveAPIRequest 接收放款结果通知 API请求
// tmall.aliauto.autofinance.loan.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
type TmallAliautoAutofinanceLoanReceiveAPIRequest struct {
	model.Params
	// 接收外部金融结构的放款结果通知参数
	_loanReceiveDto *LoanReceiveDto
}

// NewTmallAliautoAutofinanceLoanReceiveRequest 初始化TmallAliautoAutofinanceLoanReceiveAPIRequest对象
func NewTmallAliautoAutofinanceLoanReceiveRequest() *TmallAliautoAutofinanceLoanReceiveAPIRequest {
	return &TmallAliautoAutofinanceLoanReceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.autofinance.loan.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLoanReceiveDto is LoanReceiveDto Setter
// 接收外部金融结构的放款结果通知参数
func (r *TmallAliautoAutofinanceLoanReceiveAPIRequest) SetLoanReceiveDto(_loanReceiveDto *LoanReceiveDto) error {
	r._loanReceiveDto = _loanReceiveDto
	r.Set("loan_receive_dto", _loanReceiveDto)
	return nil
}

// GetLoanReceiveDto LoanReceiveDto Getter
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetLoanReceiveDto() *LoanReceiveDto {
	return r._loanReceiveDto
}
