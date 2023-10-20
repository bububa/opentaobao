package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocardexpandcardqueryAPIRequest 购物金卡查询 API请求
// taobao.card.expandcard.query
//
// 购物金充值信息查询接口，会返回余额等信息。
type TaobaocardexpandcardqueryAPIRequest struct {
	model.Params
	// 支付宝accountNo
	_accountNo string
	// 卡使用范围，不传则会查询所有
	_usedScopeCode string
}

// NewTaobaocardexpandcardqueryRequest 初始化TaobaocardexpandcardqueryAPIRequest对象
func NewTaobaocardexpandcardqueryRequest() *TaobaocardexpandcardqueryAPIRequest {
	return &TaobaocardexpandcardqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocardexpandcardqueryAPIRequest) GetApiMethodName() string {
	return "taobao.card.expandcard.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocardexpandcardqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocardexpandcardqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountNo is AccountNo Setter
// 支付宝accountNo
func (r *TaobaocardexpandcardqueryAPIRequest) SetAccountNo(_accountNo string) error {
	r._accountNo = _accountNo
	r.Set("account_no", _accountNo)
	return nil
}

// GetAccountNo AccountNo Getter
func (r TaobaocardexpandcardqueryAPIRequest) GetAccountNo() string {
	return r._accountNo
}

// SetUsedScopeCode is UsedScopeCode Setter
// 卡使用范围，不传则会查询所有
func (r *TaobaocardexpandcardqueryAPIRequest) SetUsedScopeCode(_usedScopeCode string) error {
	r._usedScopeCode = _usedScopeCode
	r.Set("used_scope_code", _usedScopeCode)
	return nil
}

// GetUsedScopeCode UsedScopeCode Getter
func (r TaobaocardexpandcardqueryAPIRequest) GetUsedScopeCode() string {
	return r._usedScopeCode
}
