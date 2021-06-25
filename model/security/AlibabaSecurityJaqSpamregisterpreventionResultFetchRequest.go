package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取垃圾注册防控结果 APIRequest
alibaba.security.jaq.spamregisterprevention.result.fetch

获取垃圾注册防控结果
*/
type AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest struct {
    model.Params

    // 账号绑定的手机号码
    phoneNumber   string 

    // 账号绑定的IP地址
    ip   string 

    // 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
    context   string 

    // 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
    source   int64 

    // 账号绑定的email地址
    email   string 

    // 账号的全局唯一标识
    userId   string 

    // 当前注册的页面URL，Source为1，2时，该参数必选
    registerUrl   string 

    // agent，发送HTTP请求的代理
    agent   string 

    // Cookie
    cookie   string 

    // Session id
    sessionId   string 

    // 硬件信息
    macAddress   string 

    // 上一跳信息
    referer   string 

    // 账号绑定的呢称
    nickName   string 

    // 账号绑定的公司名字
    companyName   string 

    // 账号绑定的注册的地址
    address   string 

    // 账号绑定的身份证号
    idNumber   string 

    // 账号绑定的银行卡号
    bankCardNumber   string 

    // 接入JS后从服务端获取的token
    jsToken   string 

    // 接入SDK后从服务端获取的token
    sdkToken   string 

}

func NewAlibabaSecurityJaqSpamregisterpreventionResultFetchRequest() *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest{
    return &AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.spamregisterprevention.result.fetch"
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetPhoneNumber() string {
    return r.phoneNumber
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetIp() string {
    return r.ip
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetContext(context string) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetContext() string {
    return r.context
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSource() int64 {
    return r.source
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetEmail() string {
    return r.email
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetRegisterUrl(registerUrl string) error {
    r.registerUrl = registerUrl
    r.Set("register_url", registerUrl)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetRegisterUrl() string {
    return r.registerUrl
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetAgent(agent string) error {
    r.agent = agent
    r.Set("agent", agent)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetAgent() string {
    return r.agent
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetCookie(cookie string) error {
    r.cookie = cookie
    r.Set("cookie", cookie)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetCookie() string {
    return r.cookie
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSessionId() string {
    return r.sessionId
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetMacAddress(macAddress string) error {
    r.macAddress = macAddress
    r.Set("mac_address", macAddress)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetMacAddress() string {
    return r.macAddress
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetReferer(referer string) error {
    r.referer = referer
    r.Set("referer", referer)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetReferer() string {
    return r.referer
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetNickName(nickName string) error {
    r.nickName = nickName
    r.Set("nick_name", nickName)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetNickName() string {
    return r.nickName
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetCompanyName() string {
    return r.companyName
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetIdNumber() string {
    return r.idNumber
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetBankCardNumber(bankCardNumber string) error {
    r.bankCardNumber = bankCardNumber
    r.Set("bank_card_number", bankCardNumber)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetBankCardNumber() string {
    return r.bankCardNumber
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetJsToken(jsToken string) error {
    r.jsToken = jsToken
    r.Set("js_token", jsToken)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetJsToken() string {
    return r.jsToken
}

func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSdkToken(sdkToken string) error {
    r.sdkToken = sdkToken
    r.Set("sdk_token", sdkToken)
    return nil
}

func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSdkToken() string {
    return r.sdkToken
}

