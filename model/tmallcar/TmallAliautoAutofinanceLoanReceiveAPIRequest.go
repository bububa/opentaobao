package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoautofinanceloanreceiveAPIRequest 接收放款结果通知 API请求
// tmall.aliauto.autofinance.loan.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
type TmallaliautoautofinanceloanreceiveAPIRequest struct {
	model.Params
	// 接收外部金融结构的放款结果通知参数
	_loanReceiveDto *LoanReceiveDto
}

// NewTmallaliautoautofinanceloanreceiveRequest 初始化TmallaliautoautofinanceloanreceiveAPIRequest对象
func NewTmallaliautoautofinanceloanreceiveRequest() *TmallaliautoautofinanceloanreceiveAPIRequest {
	return &TmallaliautoautofinanceloanreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoautofinanceloanreceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.autofinance.loan.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoautofinanceloanreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoautofinanceloanreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoanReceiveDto is LoanReceiveDto Setter
// 接收外部金融结构的放款结果通知参数
func (r *TmallaliautoautofinanceloanreceiveAPIRequest) SetLoanReceiveDto(_loanReceiveDto *LoanReceiveDto) error {
	r._loanReceiveDto = _loanReceiveDto
	r.Set("loan_receive_dto", _loanReceiveDto)
	return nil
}

// GetLoanReceiveDto LoanReceiveDto Getter
func (r TmallaliautoautofinanceloanreceiveAPIRequest) GetLoanReceiveDto() *LoanReceiveDto {
	return r._loanReceiveDto
}
