package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取虚假注册保护结果 API请求
alibaba.security.jaq.spamregisterprevention.result.fetch.new

获取虚假注册保护结果
*/
type AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest struct {
    model.Params
    // 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
    phoneNumber   string
    // 关联账号的email。 【phone_number ,email, (user_id,id_type)三种必选其一】
    email   string
    // 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
    userId   string
    // 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
    idType   int64
    // 登录时候的IP地址
    ip   string
    // 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
    source   int64
    // 当前操作的页面URL。Source为1，2时，该参数必选
    registerUrl   string
    // 发送HTTP请求的代理
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
    // 接入JS后获取的token
    jsToken   string
    // 接入无线保镖安全组件后从服务端获取的token
    sdkToken   string
    // 协议版本号。现在的值是1.0
    protocolVersion   string
    // 扩展字段。json格式的字符串，根据具体情况而定
    extendData   string
}

// 初始化AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest对象
func NewAlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest() *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest{
    return &AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.spamregisterprevention.result.fetch.new"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhoneNumber Setter
// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetPhoneNumber() string {
    return r.phoneNumber
}
// Email Setter
// 关联账号的email。 【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

// Email Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetEmail() string {
    return r.email
}
// UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetUserId() string {
    return r.userId
}
// IdType Setter
// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

// IdType Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetIdType() int64 {
    return r.idType
}
// Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetIp() string {
    return r.ip
}
// Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetSource() int64 {
    return r.source
}
// RegisterUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetRegisterUrl(registerUrl string) error {
    r.registerUrl = registerUrl
    r.Set("register_url", registerUrl)
    return nil
}

// RegisterUrl Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetRegisterUrl() string {
    return r.registerUrl
}
// Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetAgent(agent string) error {
    r.agent = agent
    r.Set("agent", agent)
    return nil
}

// Agent Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetAgent() string {
    return r.agent
}
// Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetCookie(cookie string) error {
    r.cookie = cookie
    r.Set("cookie", cookie)
    return nil
}

// Cookie Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetCookie() string {
    return r.cookie
}
// SessionId Setter
// Session id
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetSessionId() string {
    return r.sessionId
}
// MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetMacAddress(macAddress string) error {
    r.macAddress = macAddress
    r.Set("mac_address", macAddress)
    return nil
}

// MacAddress Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetMacAddress() string {
    return r.macAddress
}
// Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetReferer(referer string) error {
    r.referer = referer
    r.Set("referer", referer)
    return nil
}

// Referer Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetReferer() string {
    return r.referer
}
// NickName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetNickName(nickName string) error {
    r.nickName = nickName
    r.Set("nick_name", nickName)
    return nil
}

// NickName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetNickName() string {
    return r.nickName
}
// CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetCompanyName() string {
    return r.companyName
}
// Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetAddress() string {
    return r.address
}
// IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

// IdNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetIdNumber() string {
    return r.idNumber
}
// BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetBankCardNumber(bankCardNumber string) error {
    r.bankCardNumber = bankCardNumber
    r.Set("bank_card_number", bankCardNumber)
    return nil
}

// BankCardNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetBankCardNumber() string {
    return r.bankCardNumber
}
// JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetJsToken(jsToken string) error {
    r.jsToken = jsToken
    r.Set("js_token", jsToken)
    return nil
}

// JsToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetJsToken() string {
    return r.jsToken
}
// SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetSdkToken(sdkToken string) error {
    r.sdkToken = sdkToken
    r.Set("sdk_token", sdkToken)
    return nil
}

// SdkToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetSdkToken() string {
    return r.sdkToken
}
// ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetProtocolVersion() string {
    return r.protocolVersion
}
// ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetExtendData(extendData string) error {
    r.extendData = extendData
    r.Set("extend_data", extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetExtendData() string {
    return r.extendData
}
