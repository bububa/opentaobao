package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联名卡信息同步 APIRequest
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

func NewAlibabaAliqinFlowCobrandcardSysnRequest() *AlibabaAliqinFlowCobrandcardSysnRequest{
    return &AlibabaAliqinFlowCobrandcardSysnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.cobrandcard.sysn"
}

func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetTbUserNick(tbUserNick string) error {
    r.tbUserNick = tbUserNick
    r.Set("tb_user_nick", tbUserNick)
    return nil
}

func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetTbUserNick() string {
    return r.tbUserNick
}

func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetPhone() string {
    return r.phone
}

func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetCardType(cardType string) error {
    r.cardType = cardType
    r.Set("card_type", cardType)
    return nil
}

func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetCardType() string {
    return r.cardType
}

func (r *AlibabaAliqinFlowCobrandcardSysnRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r AlibabaAliqinFlowCobrandcardSysnRequest) GetAction() string {
    return r.action
}

