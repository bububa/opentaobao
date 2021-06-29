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
    _phoneNumber   string
    // 账号绑定的IP地址
    _ip   string
    // 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
    _context   string
    // 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
    _source   int64
    // 账号绑定的email地址
    _email   string
    // 账号的全局唯一标识
    _userId   string
    // 当前注册的页面URL，Source为1，2时，该参数必选
    _registerUrl   string
    // agent，发送HTTP请求的代理
    _agent   string
    // Cookie
    _cookie   string
    // Session id
    _sessionId   string
    // 硬件信息
    _macAddress   string
    // 上一跳信息
    _referer   string
    // 账号绑定的呢称
    _nickName   string
    // 账号绑定的公司名字
    _companyName   string
    // 账号绑定的注册的地址
    _address   string
    // 账号绑定的身份证号
    _idNumber   string
    // 账号绑定的银行卡号
    _bankCardNumber   string
    // 接入JS后从服务端获取的token
    _jsToken   string
    // 接入SDK后从服务端获取的token
    _sdkToken   string
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
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetPhoneNumber(_phoneNumber string) error {
    r._phoneNumber = _phoneNumber
    r.Set("phone_number", _phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetPhoneNumber() string {
    return r._phoneNumber
}
// Ip Setter
// 账号绑定的IP地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetIp() string {
    return r._ip
}
// Context Setter
// 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetContext(_context string) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetContext() string {
    return r._context
}
// Source Setter
// 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSource() int64 {
    return r._source
}
// Email Setter
// 账号绑定的email地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetEmail(_email string) error {
    r._email = _email
    r.Set("email", _email)
    return nil
}

// Email Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetEmail() string {
    return r._email
}
// UserId Setter
// 账号的全局唯一标识
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetUserId() string {
    return r._userId
}
// RegisterUrl Setter
// 当前注册的页面URL，Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetRegisterUrl(_registerUrl string) error {
    r._registerUrl = _registerUrl
    r.Set("register_url", _registerUrl)
    return nil
}

// RegisterUrl Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetRegisterUrl() string {
    return r._registerUrl
}
// Agent Setter
// agent，发送HTTP请求的代理
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetAgent(_agent string) error {
    r._agent = _agent
    r.Set("agent", _agent)
    return nil
}

// Agent Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetAgent() string {
    return r._agent
}
// Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetCookie(_cookie string) error {
    r._cookie = _cookie
    r.Set("cookie", _cookie)
    return nil
}

// Cookie Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetCookie() string {
    return r._cookie
}
// SessionId Setter
// Session id
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSessionId(_sessionId string) error {
    r._sessionId = _sessionId
    r.Set("session_id", _sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSessionId() string {
    return r._sessionId
}
// MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetMacAddress(_macAddress string) error {
    r._macAddress = _macAddress
    r.Set("mac_address", _macAddress)
    return nil
}

// MacAddress Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetMacAddress() string {
    return r._macAddress
}
// Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetReferer(_referer string) error {
    r._referer = _referer
    r.Set("referer", _referer)
    return nil
}

// Referer Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetReferer() string {
    return r._referer
}
// NickName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetNickName(_nickName string) error {
    r._nickName = _nickName
    r.Set("nick_name", _nickName)
    return nil
}

// NickName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetNickName() string {
    return r._nickName
}
// CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetCompanyName(_companyName string) error {
    r._companyName = _companyName
    r.Set("company_name", _companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetCompanyName() string {
    return r._companyName
}
// Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetAddress() string {
    return r._address
}
// IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetIdNumber(_idNumber string) error {
    r._idNumber = _idNumber
    r.Set("id_number", _idNumber)
    return nil
}

// IdNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetIdNumber() string {
    return r._idNumber
}
// BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetBankCardNumber(_bankCardNumber string) error {
    r._bankCardNumber = _bankCardNumber
    r.Set("bank_card_number", _bankCardNumber)
    return nil
}

// BankCardNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetBankCardNumber() string {
    return r._bankCardNumber
}
// JsToken Setter
// 接入JS后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetJsToken(_jsToken string) error {
    r._jsToken = _jsToken
    r.Set("js_token", _jsToken)
    return nil
}

// JsToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetJsToken() string {
    return r._jsToken
}
// SdkToken Setter
// 接入SDK后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) SetSdkToken(_sdkToken string) error {
    r._sdkToken = _sdkToken
    r.Set("sdk_token", _sdkToken)
    return nil
}

// SdkToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest) GetSdkToken() string {
    return r._sdkToken
}
