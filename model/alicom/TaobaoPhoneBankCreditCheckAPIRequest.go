package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaophonebankcreditcheckAPIRequest 虚拟话费任务银行信用卡办理检查 API请求
// taobao.phone.bank.credit.check
//
// 虚拟话费任务银行信用卡办理检查
type TaobaophonebankcreditcheckAPIRequest struct {
	model.Params
	// 能否办理的请求
	_bankCreditCheckRequest *BankCreditCheckRequest
}

// NewTaobaophonebankcreditcheckRequest 初始化TaobaophonebankcreditcheckAPIRequest对象
func NewTaobaophonebankcreditcheckRequest() *TaobaophonebankcreditcheckAPIRequest {
	return &TaobaophonebankcreditcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaophonebankcreditcheckAPIRequest) GetApiMethodName() string {
	return "taobao.phone.bank.credit.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaophonebankcreditcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaophonebankcreditcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBankCreditCheckRequest is BankCreditCheckRequest Setter
// 能否办理的请求
func (r *TaobaophonebankcreditcheckAPIRequest) SetBankCreditCheckRequest(_bankCreditCheckRequest *BankCreditCheckRequest) error {
	r._bankCreditCheckRequest = _bankCreditCheckRequest
	r.Set("bank_credit_check_request", _bankCreditCheckRequest)
	return nil
}

// GetBankCreditCheckRequest BankCreditCheckRequest Getter
func (r TaobaophonebankcreditcheckAPIRequest) GetBankCreditCheckRequest() *BankCreditCheckRequest {
	return r._bankCreditCheckRequest
}
