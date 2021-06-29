package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联名卡信息同步 API请求
alibaba.aliqin.flow.cobrandcard.sysn

提供给浙江移动同步联名卡信息接口。
*/
type AlibabaAliqinFlowCobrandcardSysnRequest struct {
    model.Params
    // 淘宝nick
    _tbUserNick   string
    // 手机号码
    _phone   string
    // 联名卡类型cardType:1-大喵卡,2-小喵卡
    _cardType   string
    // 1-激活，0-失效
    _action   string
}

// 初始化AlibabaAliqinFlowCobrandcardSysnRequest对象
func NewAlibabaAliqinFlowCobrandcardSysnRequest() *AlibabaAliqinFlowCobrandcardSysnRequest{
    return &AlibabaAliqinFlowCobrandcardSysnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.cobrandcard.sysn"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbUserNick Setter
// 淘宝nick
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetTbUserNick(_tbUserNick string) error {
    r._tbUserNick = _tbUserNick
    r.Set("tb_user_nick", _tbUserNick)
    return nil
}

// TbUserNick Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetTbUserNick() string {
    return r._tbUserNick
}
// Phone Setter
// 手机号码
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetPhone() string {
    return r._phone
}
// CardType Setter
// 联名卡类型cardType:1-大喵卡,2-小喵卡
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetCardType(_cardType string) error {
    r._cardType = _cardType
    r.Set("card_type", _cardType)
    return nil
}

// CardType Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetCardType() string {
    return r._cardType
}
// Action Setter
// 1-激活，0-失效
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetAction() string {
    return r._action
}
