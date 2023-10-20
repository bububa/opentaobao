package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneBankCreditProcessAPIRequest 虚拟话费任务银行信用卡办理进度回传 API请求
// taobao.phone.bank.credit.process
//
// 虚拟话费任务银行信用卡办理进度回传
type TaobaoPhoneBankCreditProcessAPIRequest struct {
	model.Params
	// 办卡进度的回调
	_bankCreditRequest *BankCreditProcessRequest
}

// NewTaobaoPhoneBankCreditProcessRequest 初始化TaobaoPhoneBankCreditProcessAPIRequest对象
func NewTaobaoPhoneBankCreditProcessRequest() *TaobaoPhoneBankCreditProcessAPIRequest {
	return &TaobaoPhoneBankCreditProcessAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPhoneBankCreditProcessAPIRequest) Reset() {
	r._bankCreditRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPhoneBankCreditProcessAPIRequest) GetApiMethodName() string {
	return "taobao.phone.bank.credit.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPhoneBankCreditProcessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPhoneBankCreditProcessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBankCreditRequest is BankCreditRequest Setter
// 办卡进度的回调
func (r *TaobaoPhoneBankCreditProcessAPIRequest) SetBankCreditRequest(_bankCreditRequest *BankCreditProcessRequest) error {
	r._bankCreditRequest = _bankCreditRequest
	r.Set("bank_credit_request", _bankCreditRequest)
	return nil
}

// GetBankCreditRequest BankCreditRequest Getter
func (r TaobaoPhoneBankCreditProcessAPIRequest) GetBankCreditRequest() *BankCreditProcessRequest {
	return r._bankCreditRequest
}

var poolTaobaoPhoneBankCreditProcessAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPhoneBankCreditProcessRequest()
	},
}

// GetTaobaoPhoneBankCreditProcessRequest 从 sync.Pool 获取 TaobaoPhoneBankCreditProcessAPIRequest
func GetTaobaoPhoneBankCreditProcessAPIRequest() *TaobaoPhoneBankCreditProcessAPIRequest {
	return poolTaobaoPhoneBankCreditProcessAPIRequest.Get().(*TaobaoPhoneBankCreditProcessAPIRequest)
}

// ReleaseTaobaoPhoneBankCreditProcessAPIRequest 将 TaobaoPhoneBankCreditProcessAPIRequest 放入 sync.Pool
func ReleaseTaobaoPhoneBankCreditProcessAPIRequest(v *TaobaoPhoneBankCreditProcessAPIRequest) {
	v.Reset()
	poolTaobaoPhoneBankCreditProcessAPIRequest.Put(v)
}
