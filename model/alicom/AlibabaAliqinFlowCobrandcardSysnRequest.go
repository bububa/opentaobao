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
    tbUserNick   string
    // 手机号码
    phone   string
    // 联名卡类型cardType:1-大喵卡,2-小喵卡
    cardType   string
    // 1-激活，0-失效
    action   string
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
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetTbUserNick(tbUserNick string) error {
    r.tbUserNick = tbUserNick
    r.Set("tb_user_nick", tbUserNick)
    return nil
}

// TbUserNick Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetTbUserNick() string {
    return r.tbUserNick
}
// Phone Setter
// 手机号码
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetPhone() string {
    return r.phone
}
// CardType Setter
// 联名卡类型cardType:1-大喵卡,2-小喵卡
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetCardType(cardType string) error {
    r.cardType = cardType
    r.Set("card_type", cardType)
    return nil
}

// CardType Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetCardType() string {
    return r.cardType
}
// Action Setter
// 1-激活，0-失效
func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetAction() string {
    return r.action
}
