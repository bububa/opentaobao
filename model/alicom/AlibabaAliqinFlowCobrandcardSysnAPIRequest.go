package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowCobrandcardSysnAPIRequest 联名卡信息同步 API请求
// alibaba.aliqin.flow.cobrandcard.sysn
//
// 提供给浙江移动同步联名卡信息接口。
type AlibabaAliqinFlowCobrandcardSysnAPIRequest struct {
	model.Params
	// 淘宝nick
	_tbUserNick string
	// 手机号码
	_phone string
	// 联名卡类型cardType:1-大喵卡,2-小喵卡
	_cardType string
	// 1-激活，0-失效
	_action string
}

// NewAlibabaAliqinFlowCobrandcardSysnRequest 初始化AlibabaAliqinFlowCobrandcardSysnAPIRequest对象
func NewAlibabaAliqinFlowCobrandcardSysnRequest() *AlibabaAliqinFlowCobrandcardSysnAPIRequest {
	return &AlibabaAliqinFlowCobrandcardSysnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowCobrandcardSysnAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.cobrandcard.sysn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowCobrandcardSysnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTbUserNick is TbUserNick Setter
// 淘宝nick
func (r *AlibabaAliqinFlowCobrandcardSysnAPIRequest) SetTbUserNick(_tbUserNick string) error {
	r._tbUserNick = _tbUserNick
	r.Set("tb_user_nick", _tbUserNick)
	return nil
}

// GetTbUserNick TbUserNick Getter
func (r AlibabaAliqinFlowCobrandcardSysnAPIRequest) GetTbUserNick() string {
	return r._tbUserNick
}

// SetPhone is Phone Setter
// 手机号码
func (r *AlibabaAliqinFlowCobrandcardSysnAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaAliqinFlowCobrandcardSysnAPIRequest) GetPhone() string {
	return r._phone
}

// SetCardType is CardType Setter
// 联名卡类型cardType:1-大喵卡,2-小喵卡
func (r *AlibabaAliqinFlowCobrandcardSysnAPIRequest) SetCardType(_cardType string) error {
	r._cardType = _cardType
	r.Set("card_type", _cardType)
	return nil
}

// GetCardType CardType Getter
func (r AlibabaAliqinFlowCobrandcardSysnAPIRequest) GetCardType() string {
	return r._cardType
}

// SetAction is Action Setter
// 1-激活，0-失效
func (r *AlibabaAliqinFlowCobrandcardSysnAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaAliqinFlowCobrandcardSysnAPIRequest) GetAction() string {
	return r._action
}
