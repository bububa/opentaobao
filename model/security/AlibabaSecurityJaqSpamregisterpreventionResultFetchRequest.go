package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取垃圾注册防控结果 API请求
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

// 初始化AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest对象
func NewAlibabaSecurityJaqSpamregisterpreventionResultFetchRequest() *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest{
    return &AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.spamregisterprevention.result.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhoneNumber Setter
// 账号绑定的手机号码
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetPhoneNumber() string {
    return r.phoneNumber
}
// Ip Setter
// 账号绑定的IP地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetIp() string {
    return r.ip
}
// Context Setter
// 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetContext(context string) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetContext() string {
    return r.context
}
// Source Setter
// 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSource() int64 {
    return r.source
}
// Email Setter
// 账号绑定的email地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

// Email Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetEmail() string {
    return r.email
}
// UserId Setter
// 账号的全局唯一标识
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetUserId() string {
    return r.userId
}
// RegisterUrl Setter
// 当前注册的页面URL，Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetRegisterUrl(registerUrl string) error {
    r.registerUrl = registerUrl
    r.Set("register_url", registerUrl)
    return nil
}

// RegisterUrl Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetRegisterUrl() string {
    return r.registerUrl
}
// Agent Setter
// agent，发送HTTP请求的代理
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetAgent(agent string) error {
    r.agent = agent
    r.Set("agent", agent)
    return nil
}

// Agent Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetAgent() string {
    return r.agent
}
// Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetCookie(cookie string) error {
    r.cookie = cookie
    r.Set("cookie", cookie)
    return nil
}

// Cookie Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetCookie() string {
    return r.cookie
}
// SessionId Setter
// Session id
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSessionId() string {
    return r.sessionId
}
// MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetMacAddress(macAddress string) error {
    r.macAddress = macAddress
    r.Set("mac_address", macAddress)
    return nil
}

// MacAddress Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetMacAddress() string {
    return r.macAddress
}
// Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetReferer(referer string) error {
    r.referer = referer
    r.Set("referer", referer)
    return nil
}

// Referer Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetReferer() string {
    return r.referer
}
// NickName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetNickName(nickName string) error {
    r.nickName = nickName
    r.Set("nick_name", nickName)
    return nil
}

// NickName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetNickName() string {
    return r.nickName
}
// CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetCompanyName() string {
    return r.companyName
}
// Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetAddress() string {
    return r.address
}
// IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

// IdNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetIdNumber() string {
    return r.idNumber
}
// BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetBankCardNumber(bankCardNumber string) error {
    r.bankCardNumber = bankCardNumber
    r.Set("bank_card_number", bankCardNumber)
    return nil
}

// BankCardNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetBankCardNumber() string {
    return r.bankCardNumber
}
// JsToken Setter
// 接入JS后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetJsToken(jsToken string) error {
    r.jsToken = jsToken
    r.Set("js_token", jsToken)
    return nil
}

// JsToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetJsToken() string {
    return r.jsToken
}
// SdkToken Setter
// 接入SDK后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSdkToken(sdkToken string) error {
    r.sdkToken = sdkToken
    r.Set("sdk_token", sdkToken)
    return nil
}

// SdkToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSdkToken() string {
    return r.sdkToken
}
