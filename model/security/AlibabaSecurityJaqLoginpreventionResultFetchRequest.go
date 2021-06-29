package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登录保护结果 API请求
alibaba.security.jaq.loginprevention.result.fetch

获取登录保护结果
*/
type AlibabaSecurityJaqLoginpreventionResultFetchRequest struct {
    model.Params
    // 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
    phoneNumber   string
    // 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
    email   string
    // 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
    userId   string
    // 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
    idType   int64
    // 登录时候的IP地址
    ip   string
    // 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
    source   int64
    // 当前操作的页面URL。Source为1，2时，该参数必选
    currentUrl   string
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
    userName   string
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
    // 扩展字段。json格式的字符串，根据具体情况而定 。
    extendData   string
    // 账号在系统里面是否存在。0：不存在；1：存在
    accountExist   int64
    // 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
    loginType   int64
    // 密码是否正确。0：不正确；1：正确
    passwordCorrect   int64
    // 将密码加盐hash后传递，用于弱密码检测
    passwordHash   string
    // 注册的时间（秒）
    registerDate   int64
    // 注册时候的ip
    registerIp   string
}

// 初始化AlibabaSecurityJaqLoginpreventionResultFetchRequest对象
func NewAlibabaSecurityJaqLoginpreventionResultFetchRequest() *AlibabaSecurityJaqLoginpreventionResultFetchRequest{
    return &AlibabaSecurityJaqLoginpreventionResultFetchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.loginprevention.result.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhoneNumber Setter
// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPhoneNumber() string {
    return r.phoneNumber
}
// Email Setter
// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

// Email Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetEmail() string {
    return r.email
}
// UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetUserId() string {
    return r.userId
}
// IdType Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

// IdType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIdType() int64 {
    return r.idType
}
// Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIp() string {
    return r.ip
}
// Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSource() int64 {
    return r.source
}
// CurrentUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCurrentUrl(currentUrl string) error {
    r.currentUrl = currentUrl
    r.Set("current_url", currentUrl)
    return nil
}

// CurrentUrl Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCurrentUrl() string {
    return r.currentUrl
}
// Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAgent(agent string) error {
    r.agent = agent
    r.Set("agent", agent)
    return nil
}

// Agent Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAgent() string {
    return r.agent
}
// Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCookie(cookie string) error {
    r.cookie = cookie
    r.Set("cookie", cookie)
    return nil
}

// Cookie Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCookie() string {
    return r.cookie
}
// SessionId Setter
// Session id
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSessionId() string {
    return r.sessionId
}
// MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetMacAddress(macAddress string) error {
    r.macAddress = macAddress
    r.Set("mac_address", macAddress)
    return nil
}

// MacAddress Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetMacAddress() string {
    return r.macAddress
}
// Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetReferer(referer string) error {
    r.referer = referer
    r.Set("referer", referer)
    return nil
}

// Referer Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetReferer() string {
    return r.referer
}
// UserName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

// UserName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetUserName() string {
    return r.userName
}
// CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCompanyName() string {
    return r.companyName
}
// Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAddress() string {
    return r.address
}
// IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

// IdNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIdNumber() string {
    return r.idNumber
}
// BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetBankCardNumber(bankCardNumber string) error {
    r.bankCardNumber = bankCardNumber
    r.Set("bank_card_number", bankCardNumber)
    return nil
}

// BankCardNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetBankCardNumber() string {
    return r.bankCardNumber
}
// JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetJsToken(jsToken string) error {
    r.jsToken = jsToken
    r.Set("js_token", jsToken)
    return nil
}

// JsToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetJsToken() string {
    return r.jsToken
}
// SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSdkToken(sdkToken string) error {
    r.sdkToken = sdkToken
    r.Set("sdk_token", sdkToken)
    return nil
}

// SdkToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSdkToken() string {
    return r.sdkToken
}
// ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetProtocolVersion() string {
    return r.protocolVersion
}
// ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定 。
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetExtendData(extendData string) error {
    r.extendData = extendData
    r.Set("extend_data", extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetExtendData() string {
    return r.extendData
}
// AccountExist Setter
// 账号在系统里面是否存在。0：不存在；1：存在
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAccountExist(accountExist int64) error {
    r.accountExist = accountExist
    r.Set("account_exist", accountExist)
    return nil
}

// AccountExist Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAccountExist() int64 {
    return r.accountExist
}
// LoginType Setter
// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetLoginType(loginType int64) error {
    r.loginType = loginType
    r.Set("login_type", loginType)
    return nil
}

// LoginType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetLoginType() int64 {
    return r.loginType
}
// PasswordCorrect Setter
// 密码是否正确。0：不正确；1：正确
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPasswordCorrect(passwordCorrect int64) error {
    r.passwordCorrect = passwordCorrect
    r.Set("password_correct", passwordCorrect)
    return nil
}

// PasswordCorrect Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPasswordCorrect() int64 {
    return r.passwordCorrect
}
// PasswordHash Setter
// 将密码加盐hash后传递，用于弱密码检测
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPasswordHash(passwordHash string) error {
    r.passwordHash = passwordHash
    r.Set("password_hash", passwordHash)
    return nil
}

// PasswordHash Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPasswordHash() string {
    return r.passwordHash
}
// RegisterDate Setter
// 注册的时间（秒）
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetRegisterDate(registerDate int64) error {
    r.registerDate = registerDate
    r.Set("register_date", registerDate)
    return nil
}

// RegisterDate Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetRegisterDate() int64 {
    return r.registerDate
}
// RegisterIp Setter
// 注册时候的ip
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetRegisterIp(registerIp string) error {
    r.registerIp = registerIp
    r.Set("register_ip", registerIp)
    return nil
}

// RegisterIp Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetRegisterIp() string {
    return r.registerIp
}
