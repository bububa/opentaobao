package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登录保护结果 APIRequest
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

func NewAlibabaSecurityJaqLoginpreventionResultFetchRequest() *AlibabaSecurityJaqLoginpreventionResultFetchRequest{
    return &AlibabaSecurityJaqLoginpreventionResultFetchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.loginprevention.result.fetch"
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPhoneNumber() string {
    return r.phoneNumber
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetEmail() string {
    return r.email
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIdType() int64 {
    return r.idType
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIp() string {
    return r.ip
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSource() int64 {
    return r.source
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCurrentUrl(currentUrl string) error {
    r.currentUrl = currentUrl
    r.Set("current_url", currentUrl)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCurrentUrl() string {
    return r.currentUrl
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAgent(agent string) error {
    r.agent = agent
    r.Set("agent", agent)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAgent() string {
    return r.agent
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCookie(cookie string) error {
    r.cookie = cookie
    r.Set("cookie", cookie)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCookie() string {
    return r.cookie
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSessionId() string {
    return r.sessionId
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetMacAddress(macAddress string) error {
    r.macAddress = macAddress
    r.Set("mac_address", macAddress)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetMacAddress() string {
    return r.macAddress
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetReferer(referer string) error {
    r.referer = referer
    r.Set("referer", referer)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetReferer() string {
    return r.referer
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetUserName() string {
    return r.userName
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCompanyName() string {
    return r.companyName
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIdNumber() string {
    return r.idNumber
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetBankCardNumber(bankCardNumber string) error {
    r.bankCardNumber = bankCardNumber
    r.Set("bank_card_number", bankCardNumber)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetBankCardNumber() string {
    return r.bankCardNumber
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetJsToken(jsToken string) error {
    r.jsToken = jsToken
    r.Set("js_token", jsToken)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetJsToken() string {
    return r.jsToken
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSdkToken(sdkToken string) error {
    r.sdkToken = sdkToken
    r.Set("sdk_token", sdkToken)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSdkToken() string {
    return r.sdkToken
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetProtocolVersion() string {
    return r.protocolVersion
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetExtendData(extendData string) error {
    r.extendData = extendData
    r.Set("extend_data", extendData)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetExtendData() string {
    return r.extendData
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAccountExist(accountExist int64) error {
    r.accountExist = accountExist
    r.Set("account_exist", accountExist)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAccountExist() int64 {
    return r.accountExist
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetLoginType(loginType int64) error {
    r.loginType = loginType
    r.Set("login_type", loginType)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetLoginType() int64 {
    return r.loginType
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPasswordCorrect(passwordCorrect int64) error {
    r.passwordCorrect = passwordCorrect
    r.Set("password_correct", passwordCorrect)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPasswordCorrect() int64 {
    return r.passwordCorrect
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPasswordHash(passwordHash string) error {
    r.passwordHash = passwordHash
    r.Set("password_hash", passwordHash)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPasswordHash() string {
    return r.passwordHash
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetRegisterDate(registerDate int64) error {
    r.registerDate = registerDate
    r.Set("register_date", registerDate)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetRegisterDate() int64 {
    return r.registerDate
}

func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetRegisterIp(registerIp string) error {
    r.registerIp = registerIp
    r.Set("register_ip", registerIp)
    return nil
}

func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetRegisterIp() string {
    return r.registerIp
}

