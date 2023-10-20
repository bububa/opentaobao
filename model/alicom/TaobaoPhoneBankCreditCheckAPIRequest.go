package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneBankCreditCheckAPIRequest 虚拟话费任务银行信用卡办理检查 API请求
// taobao.phone.bank.credit.check
//
// 虚拟话费任务银行信用卡办理检查
type TaobaoPhoneBankCreditCheckAPIRequest struct {
	model.Params
	// 能否办理的请求
	_bankCreditCheckRequest *BankCreditCheckRequest
}

// NewTaobaoPhoneBankCreditCheckRequest 初始化TaobaoPhoneBankCreditCheckAPIRequest对象
func NewTaobaoPhoneBankCreditCheckRequest() *TaobaoPhoneBankCreditCheckAPIRequest {
	return &TaobaoPhoneBankCreditCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPhoneBankCreditCheckAPIRequest) Reset() {
	r._bankCreditCheckRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPhoneBankCreditCheckAPIRequest) GetApiMethodName() string {
	return "taobao.phone.bank.credit.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPhoneBankCreditCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPhoneBankCreditCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBankCreditCheckRequest is BankCreditCheckRequest Setter
// 能否办理的请求
func (r *TaobaoPhoneBankCreditCheckAPIRequest) SetBankCreditCheckRequest(_bankCreditCheckRequest *BankCreditCheckRequest) error {
	r._bankCreditCheckRequest = _bankCreditCheckRequest
	r.Set("bank_credit_check_request", _bankCreditCheckRequest)
	return nil
}

// GetBankCreditCheckRequest BankCreditCheckRequest Getter
func (r TaobaoPhoneBankCreditCheckAPIRequest) GetBankCreditCheckRequest() *BankCreditCheckRequest {
	return r._bankCreditCheckRequest
}

var poolTaobaoPhoneBankCreditCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPhoneBankCreditCheckRequest()
	},
}

// GetTaobaoPhoneBankCreditCheckRequest 从 sync.Pool 获取 TaobaoPhoneBankCreditCheckAPIRequest
func GetTaobaoPhoneBankCreditCheckAPIRequest() *TaobaoPhoneBankCreditCheckAPIRequest {
	return poolTaobaoPhoneBankCreditCheckAPIRequest.Get().(*TaobaoPhoneBankCreditCheckAPIRequest)
}

// ReleaseTaobaoPhoneBankCreditCheckAPIRequest 将 TaobaoPhoneBankCreditCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoPhoneBankCreditCheckAPIRequest(v *TaobaoPhoneBankCreditCheckAPIRequest) {
	v.Reset()
	poolTaobaoPhoneBankCreditCheckAPIRequest.Put(v)
}
