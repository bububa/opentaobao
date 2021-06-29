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
    _phoneNumber   string
    // 关联账号的email。 【phone_number ,email, (user_id,id_type)三种必选其一】
    _email   string
    // 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
    _userId   string
    // 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
    _idType   int64
    // 登录时候的IP地址
    _ip   string
    // 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
    _source   int64
    // 当前操作的页面URL。Source为1，2时，该参数必选
    _registerUrl   string
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
    _nickName   string
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
    // 扩展字段。json格式的字符串，根据具体情况而定
    _extendData   string
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
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetPhoneNumber(_phoneNumber string) error {
    r._phoneNumber = _phoneNumber
    r.Set("phone_number", _phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetPhoneNumber() string {
    return r._phoneNumber
}
// Email Setter
// 关联账号的email。 【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetEmail(_email string) error {
    r._email = _email
    r.Set("email", _email)
    return nil
}

// Email Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetEmail() string {
    return r._email
}
// UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetUserId() string {
    return r._userId
}
// IdType Setter
// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetIdType(_idType int64) error {
    r._idType = _idType
    r.Set("id_type", _idType)
    return nil
}

// IdType Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetIdType() int64 {
    return r._idType
}
// Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetIp() string {
    return r._ip
}
// Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetSource() int64 {
    return r._source
}
// RegisterUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetRegisterUrl(_registerUrl string) error {
    r._registerUrl = _registerUrl
    r.Set("register_url", _registerUrl)
    return nil
}

// RegisterUrl Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetRegisterUrl() string {
    return r._registerUrl
}
// Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetAgent(_agent string) error {
    r._agent = _agent
    r.Set("agent", _agent)
    return nil
}

// Agent Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetAgent() string {
    return r._agent
}
// Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetCookie(_cookie string) error {
    r._cookie = _cookie
    r.Set("cookie", _cookie)
    return nil
}

// Cookie Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetCookie() string {
    return r._cookie
}
// SessionId Setter
// Session id
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetSessionId(_sessionId string) error {
    r._sessionId = _sessionId
    r.Set("session_id", _sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetSessionId() string {
    return r._sessionId
}
// MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetMacAddress(_macAddress string) error {
    r._macAddress = _macAddress
    r.Set("mac_address", _macAddress)
    return nil
}

// MacAddress Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetMacAddress() string {
    return r._macAddress
}
// Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetReferer(_referer string) error {
    r._referer = _referer
    r.Set("referer", _referer)
    return nil
}

// Referer Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetReferer() string {
    return r._referer
}
// NickName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetNickName(_nickName string) error {
    r._nickName = _nickName
    r.Set("nick_name", _nickName)
    return nil
}

// NickName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetNickName() string {
    return r._nickName
}
// CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetCompanyName(_companyName string) error {
    r._companyName = _companyName
    r.Set("company_name", _companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetCompanyName() string {
    return r._companyName
}
// Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetAddress() string {
    return r._address
}
// IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetIdNumber(_idNumber string) error {
    r._idNumber = _idNumber
    r.Set("id_number", _idNumber)
    return nil
}

// IdNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetIdNumber() string {
    return r._idNumber
}
// BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetBankCardNumber(_bankCardNumber string) error {
    r._bankCardNumber = _bankCardNumber
    r.Set("bank_card_number", _bankCardNumber)
    return nil
}

// BankCardNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetBankCardNumber() string {
    return r._bankCardNumber
}
// JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetJsToken(_jsToken string) error {
    r._jsToken = _jsToken
    r.Set("js_token", _jsToken)
    return nil
}

// JsToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetJsToken() string {
    return r._jsToken
}
// SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetSdkToken(_sdkToken string) error {
    r._sdkToken = _sdkToken
    r.Set("sdk_token", _sdkToken)
    return nil
}

// SdkToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetSdkToken() string {
    return r._sdkToken
}
// ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetProtocolVersion(_protocolVersion string) error {
    r._protocolVersion = _protocolVersion
    r.Set("protocol_version", _protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetProtocolVersion() string {
    return r._protocolVersion
}
// ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) SetExtendData(_extendData string) error {
    r._extendData = _extendData
    r.Set("extend_data", _extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest) GetExtendData() string {
    return r._extendData
}
