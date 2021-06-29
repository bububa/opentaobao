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
    _phoneNumber   string
    // 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
    _email   string
    // 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
    _userId   string
    // 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
    _idType   int64
    // 登录时候的IP地址
    _ip   string
    // 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
    _source   int64
    // 当前操作的页面URL。Source为1，2时，该参数必选
    _currentUrl   string
    // 发送HTTP请求的代理
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
    _userName   string
    // 账号绑定的公司名字
    _companyName   string
    // 账号绑定的注册的地址
    _address   string
    // 账号绑定的身份证号
    _idNumber   string
    // 账号绑定的银行卡号
    _bankCardNumber   string
    // 接入JS后获取的token
    _jsToken   string
    // 接入无线保镖安全组件后从服务端获取的token
    _sdkToken   string
    // 协议版本号。现在的值是1.0
    _protocolVersion   string
    // 扩展字段。json格式的字符串，根据具体情况而定 。
    _extendData   string
    // 账号在系统里面是否存在。0：不存在；1：存在
    _accountExist   int64
    // 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
    _loginType   int64
    // 密码是否正确。0：不正确；1：正确
    _passwordCorrect   int64
    // 将密码加盐hash后传递，用于弱密码检测
    _passwordHash   string
    // 注册的时间（秒）
    _registerDate   int64
    // 注册时候的ip
    _registerIp   string
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
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPhoneNumber(_phoneNumber string) error {
    r._phoneNumber = _phoneNumber
    r.Set("phone_number", _phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPhoneNumber() string {
    return r._phoneNumber
}
// Email Setter
// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetEmail(_email string) error {
    r._email = _email
    r.Set("email", _email)
    return nil
}

// Email Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetEmail() string {
    return r._email
}
// UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetUserId() string {
    return r._userId
}
// IdType Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIdType(_idType int64) error {
    r._idType = _idType
    r.Set("id_type", _idType)
    return nil
}

// IdType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIdType() int64 {
    return r._idType
}
// Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIp() string {
    return r._ip
}
// Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSource() int64 {
    return r._source
}
// CurrentUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCurrentUrl(_currentUrl string) error {
    r._currentUrl = _currentUrl
    r.Set("current_url", _currentUrl)
    return nil
}

// CurrentUrl Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCurrentUrl() string {
    return r._currentUrl
}
// Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAgent(_agent string) error {
    r._agent = _agent
    r.Set("agent", _agent)
    return nil
}

// Agent Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAgent() string {
    return r._agent
}
// Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCookie(_cookie string) error {
    r._cookie = _cookie
    r.Set("cookie", _cookie)
    return nil
}

// Cookie Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCookie() string {
    return r._cookie
}
// SessionId Setter
// Session id
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSessionId(_sessionId string) error {
    r._sessionId = _sessionId
    r.Set("session_id", _sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSessionId() string {
    return r._sessionId
}
// MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetMacAddress(_macAddress string) error {
    r._macAddress = _macAddress
    r.Set("mac_address", _macAddress)
    return nil
}

// MacAddress Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetMacAddress() string {
    return r._macAddress
}
// Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetReferer(_referer string) error {
    r._referer = _referer
    r.Set("referer", _referer)
    return nil
}

// Referer Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetReferer() string {
    return r._referer
}
// UserName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetUserName(_userName string) error {
    r._userName = _userName
    r.Set("user_name", _userName)
    return nil
}

// UserName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetUserName() string {
    return r._userName
}
// CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetCompanyName(_companyName string) error {
    r._companyName = _companyName
    r.Set("company_name", _companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetCompanyName() string {
    return r._companyName
}
// Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAddress() string {
    return r._address
}
// IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetIdNumber(_idNumber string) error {
    r._idNumber = _idNumber
    r.Set("id_number", _idNumber)
    return nil
}

// IdNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetIdNumber() string {
    return r._idNumber
}
// BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetBankCardNumber(_bankCardNumber string) error {
    r._bankCardNumber = _bankCardNumber
    r.Set("bank_card_number", _bankCardNumber)
    return nil
}

// BankCardNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetBankCardNumber() string {
    return r._bankCardNumber
}
// JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetJsToken(_jsToken string) error {
    r._jsToken = _jsToken
    r.Set("js_token", _jsToken)
    return nil
}

// JsToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetJsToken() string {
    return r._jsToken
}
// SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetSdkToken(_sdkToken string) error {
    r._sdkToken = _sdkToken
    r.Set("sdk_token", _sdkToken)
    return nil
}

// SdkToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetSdkToken() string {
    return r._sdkToken
}
// ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetProtocolVersion(_protocolVersion string) error {
    r._protocolVersion = _protocolVersion
    r.Set("protocol_version", _protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetProtocolVersion() string {
    return r._protocolVersion
}
// ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定 。
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetExtendData(_extendData string) error {
    r._extendData = _extendData
    r.Set("extend_data", _extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetExtendData() string {
    return r._extendData
}
// AccountExist Setter
// 账号在系统里面是否存在。0：不存在；1：存在
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetAccountExist(_accountExist int64) error {
    r._accountExist = _accountExist
    r.Set("account_exist", _accountExist)
    return nil
}

// AccountExist Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetAccountExist() int64 {
    return r._accountExist
}
// LoginType Setter
// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetLoginType(_loginType int64) error {
    r._loginType = _loginType
    r.Set("login_type", _loginType)
    return nil
}

// LoginType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetLoginType() int64 {
    return r._loginType
}
// PasswordCorrect Setter
// 密码是否正确。0：不正确；1：正确
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPasswordCorrect(_passwordCorrect int64) error {
    r._passwordCorrect = _passwordCorrect
    r.Set("password_correct", _passwordCorrect)
    return nil
}

// PasswordCorrect Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPasswordCorrect() int64 {
    return r._passwordCorrect
}
// PasswordHash Setter
// 将密码加盐hash后传递，用于弱密码检测
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetPasswordHash(_passwordHash string) error {
    r._passwordHash = _passwordHash
    r.Set("password_hash", _passwordHash)
    return nil
}

// PasswordHash Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetPasswordHash() string {
    return r._passwordHash
}
// RegisterDate Setter
// 注册的时间（秒）
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetRegisterDate(_registerDate int64) error {
    r._registerDate = _registerDate
    r.Set("register_date", _registerDate)
    return nil
}

// RegisterDate Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetRegisterDate() int64 {
    return r._registerDate
}
// RegisterIp Setter
// 注册时候的ip
func (r *AlibabaSecurityJaqLoginpreventionResultFetchRequest) SetRegisterIp(_registerIp string) error {
    r._registerIp = _registerIp
    r.Set("register_ip", _registerIp)
    return nil
}

// RegisterIp Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchRequest) GetRegisterIp() string {
    return r._registerIp
}
