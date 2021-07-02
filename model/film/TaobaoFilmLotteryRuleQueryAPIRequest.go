package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmLotteryRuleQueryAPIRequest 淘票票抽奖活动查询API(渠道) API请求
// taobao.film.lottery.rule.query
//
// 淘票票抽奖活动查询API，渠道维度查询
type TaobaoFilmLotteryRuleQueryAPIRequest struct {
	model.Params
	// 账号类型（TAOBAO\ALIPAY\PHONE）
	_accountType string
	// 渠道来源
	_channel string
	// 账号ID
	_accountNo string
}

// NewTaobaoFilmLotteryRuleQueryRequest 初始化TaobaoFilmLotteryRuleQueryAPIRequest对象
func NewTaobaoFilmLotteryRuleQueryRequest() *TaobaoFilmLotteryRuleQueryAPIRequest {
	return &TaobaoFilmLotteryRuleQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotteryRuleQueryAPIRequest) GetApiMethodName() string {
	return "taobao.film.lottery.rule.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotteryRuleQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAccountType is AccountType Setter
// 账号类型（TAOBAO\ALIPAY\PHONE）
func (r *TaobaoFilmLotteryRuleQueryAPIRequest) SetAccountType(_accountType string) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r TaobaoFilmLotteryRuleQueryAPIRequest) GetAccountType() string {
	return r._accountType
}

// SetChannel is Channel Setter
// 渠道来源
func (r *TaobaoFilmLotteryRuleQueryAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoFilmLotteryRuleQueryAPIRequest) GetChannel() string {
	return r._channel
}

// SetAccountNo is AccountNo Setter
// 账号ID
func (r *TaobaoFilmLotteryRuleQueryAPIRequest) SetAccountNo(_accountNo string) error {
	r._accountNo = _accountNo
	r.Set("account_no", _accountNo)
	return nil
}

// GetAccountNo AccountNo Getter
func (r TaobaoFilmLotteryRuleQueryAPIRequest) GetAccountNo() string {
	return r._accountNo
}
