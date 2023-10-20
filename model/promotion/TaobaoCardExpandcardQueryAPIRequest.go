package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCardExpandcardQueryAPIRequest 购物金卡查询 API请求
// taobao.card.expandcard.query
//
// 购物金充值信息查询接口，会返回余额等信息。
type TaobaoCardExpandcardQueryAPIRequest struct {
	model.Params
	// 支付宝accountNo
	_accountNo string
	// 卡使用范围，不传则会查询所有
	_usedScopeCode string
}

// NewTaobaoCardExpandcardQueryRequest 初始化TaobaoCardExpandcardQueryAPIRequest对象
func NewTaobaoCardExpandcardQueryRequest() *TaobaoCardExpandcardQueryAPIRequest {
	return &TaobaoCardExpandcardQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCardExpandcardQueryAPIRequest) Reset() {
	r._accountNo = ""
	r._usedScopeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCardExpandcardQueryAPIRequest) GetApiMethodName() string {
	return "taobao.card.expandcard.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCardExpandcardQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCardExpandcardQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountNo is AccountNo Setter
// 支付宝accountNo
func (r *TaobaoCardExpandcardQueryAPIRequest) SetAccountNo(_accountNo string) error {
	r._accountNo = _accountNo
	r.Set("account_no", _accountNo)
	return nil
}

// GetAccountNo AccountNo Getter
func (r TaobaoCardExpandcardQueryAPIRequest) GetAccountNo() string {
	return r._accountNo
}

// SetUsedScopeCode is UsedScopeCode Setter
// 卡使用范围，不传则会查询所有
func (r *TaobaoCardExpandcardQueryAPIRequest) SetUsedScopeCode(_usedScopeCode string) error {
	r._usedScopeCode = _usedScopeCode
	r.Set("used_scope_code", _usedScopeCode)
	return nil
}

// GetUsedScopeCode UsedScopeCode Getter
func (r TaobaoCardExpandcardQueryAPIRequest) GetUsedScopeCode() string {
	return r._usedScopeCode
}

var poolTaobaoCardExpandcardQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCardExpandcardQueryRequest()
	},
}

// GetTaobaoCardExpandcardQueryRequest 从 sync.Pool 获取 TaobaoCardExpandcardQueryAPIRequest
func GetTaobaoCardExpandcardQueryAPIRequest() *TaobaoCardExpandcardQueryAPIRequest {
	return poolTaobaoCardExpandcardQueryAPIRequest.Get().(*TaobaoCardExpandcardQueryAPIRequest)
}

// ReleaseTaobaoCardExpandcardQueryAPIRequest 将 TaobaoCardExpandcardQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoCardExpandcardQueryAPIRequest(v *TaobaoCardExpandcardQueryAPIRequest) {
	v.Reset()
	poolTaobaoCardExpandcardQueryAPIRequest.Put(v)
}
