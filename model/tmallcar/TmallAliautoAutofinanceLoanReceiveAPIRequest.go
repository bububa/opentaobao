package tmallcar

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoAutofinanceLoanReceiveAPIRequest) Reset() {
	r._loanReceiveDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.autofinance.loan.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoAutofinanceLoanReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallAliautoAutofinanceLoanReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoAutofinanceLoanReceiveRequest()
	},
}

// GetTmallAliautoAutofinanceLoanReceiveRequest 从 sync.Pool 获取 TmallAliautoAutofinanceLoanReceiveAPIRequest
func GetTmallAliautoAutofinanceLoanReceiveAPIRequest() *TmallAliautoAutofinanceLoanReceiveAPIRequest {
	return poolTmallAliautoAutofinanceLoanReceiveAPIRequest.Get().(*TmallAliautoAutofinanceLoanReceiveAPIRequest)
}

// ReleaseTmallAliautoAutofinanceLoanReceiveAPIRequest 将 TmallAliautoAutofinanceLoanReceiveAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoAutofinanceLoanReceiveAPIRequest(v *TmallAliautoAutofinanceLoanReceiveAPIRequest) {
	v.Reset()
	poolTmallAliautoAutofinanceLoanReceiveAPIRequest.Put(v)
}
