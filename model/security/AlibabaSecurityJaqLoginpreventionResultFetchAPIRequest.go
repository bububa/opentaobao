package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest 获取登录保护结果 API请求
// alibaba.security.jaq.loginprevention.result.fetch
//
// 获取登录保护结果
type AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest struct {
	model.Params
	// 账号绑定的身份证号
	_idNumber string
	// 接入无线保镖安全组件后从服务端获取的token
	_sdkToken string
	// 账号绑定的注册的地址
	_address string
	// 发送HTTP请求的代理
	_agent string
	// 账号绑定的银行卡号
	_bankCardNumber string
	// 账号绑定的公司名字
	_companyName string
	// Cookie
	_cookie string
	// 当前操作的页面URL。Source为1，2时，该参数必选
	_currentUrl string
	// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 扩展字段。json格式的字符串，根据具体情况而定 。
	_extendData string
	// 登录时候的IP地址
	_ip string
	// 接入JS后获取的token
	_jsToken string
	// 硬件信息
	_macAddress string
	// 将密码加盐hash后传递，用于弱密码检测
	_passwordHash string
	// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 协议版本号。现在的值是1.0
	_protocolVersion string
	// 上一跳信息
	_referer string
	// 注册时候的ip
	_registerIp string
	// Session id
	_sessionId string
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 账号绑定的呢称
	_userName string
	// 账号在系统里面是否存在。0：不存在；1：存在
	_accountExist int64
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
	_idType int64
	// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
	_loginType int64
	// 密码是否正确。0：不正确；1：正确
	_passwordCorrect int64
	// 注册的时间（秒）
	_registerDate int64
	// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
	_source int64
}

// NewAlibabaSecurityJaqLoginpreventionResultFetchRequest 初始化AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest对象
func NewAlibabaSecurityJaqLoginpreventionResultFetchRequest() *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest {
	return &AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest{
		Params: model.NewParams(27),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) Reset() {
	r._idNumber = ""
	r._sdkToken = ""
	r._address = ""
	r._agent = ""
	r._bankCardNumber = ""
	r._companyName = ""
	r._cookie = ""
	r._currentUrl = ""
	r._email = ""
	r._extendData = ""
	r._ip = ""
	r._jsToken = ""
	r._macAddress = ""
	r._passwordHash = ""
	r._phoneNumber = ""
	r._protocolVersion = ""
	r._referer = ""
	r._registerIp = ""
	r._sessionId = ""
	r._userId = ""
	r._userName = ""
	r._accountExist = 0
	r._idType = 0
	r._loginType = 0
	r._passwordCorrect = 0
	r._registerDate = 0
	r._source = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.loginprevention.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdNumber is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetSdkToken is SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// GetSdkToken SdkToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// SetAddress is Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetAddress() string {
	return r._address
}

// SetAgent is Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// GetAgent Agent Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetAgent() string {
	return r._agent
}

// SetBankCardNumber is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// GetBankCardNumber BankCardNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// SetCompanyName is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetCookie is Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// GetCookie Cookie Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetCookie() string {
	return r._cookie
}

// SetCurrentUrl is CurrentUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetCurrentUrl(_currentUrl string) error {
	r._currentUrl = _currentUrl
	r.Set("current_url", _currentUrl)
	return nil
}

// GetCurrentUrl CurrentUrl Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetCurrentUrl() string {
	return r._currentUrl
}

// SetEmail is Email Setter
// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetEmail() string {
	return r._email
}

// SetExtendData is ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定 。
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// GetExtendData ExtendData Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetExtendData() string {
	return r._extendData
}

// SetIp is Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetIp() string {
	return r._ip
}

// SetJsToken is JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// GetJsToken JsToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetJsToken() string {
	return r._jsToken
}

// SetMacAddress is MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// GetMacAddress MacAddress Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// SetPasswordHash is PasswordHash Setter
// 将密码加盐hash后传递，用于弱密码检测
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetPasswordHash(_passwordHash string) error {
	r._passwordHash = _passwordHash
	r.Set("password_hash", _passwordHash)
	return nil
}

// GetPasswordHash PasswordHash Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetPasswordHash() string {
	return r._passwordHash
}

// SetPhoneNumber is PhoneNumber Setter
// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// GetPhoneNumber PhoneNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetReferer is Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// GetReferer Referer Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetReferer() string {
	return r._referer
}

// SetRegisterIp is RegisterIp Setter
// 注册时候的ip
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetRegisterIp(_registerIp string) error {
	r._registerIp = _registerIp
	r.Set("register_ip", _registerIp)
	return nil
}

// GetRegisterIp RegisterIp Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetRegisterIp() string {
	return r._registerIp
}

// SetSessionId is SessionId Setter
// Session id
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetSessionId() string {
	return r._sessionId
}

// SetUserId is UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetUserId() string {
	return r._userId
}

// SetUserName is UserName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetUserName() string {
	return r._userName
}

// SetAccountExist is AccountExist Setter
// 账号在系统里面是否存在。0：不存在；1：存在
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetAccountExist(_accountExist int64) error {
	r._accountExist = _accountExist
	r.Set("account_exist", _accountExist)
	return nil
}

// GetAccountExist AccountExist Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetAccountExist() int64 {
	return r._accountExist
}

// SetIdType is IdType Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// GetIdType IdType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetIdType() int64 {
	return r._idType
}

// SetLoginType is LoginType Setter
// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetLoginType(_loginType int64) error {
	r._loginType = _loginType
	r.Set("login_type", _loginType)
	return nil
}

// GetLoginType LoginType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetLoginType() int64 {
	return r._loginType
}

// SetPasswordCorrect is PasswordCorrect Setter
// 密码是否正确。0：不正确；1：正确
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetPasswordCorrect(_passwordCorrect int64) error {
	r._passwordCorrect = _passwordCorrect
	r.Set("password_correct", _passwordCorrect)
	return nil
}

// GetPasswordCorrect PasswordCorrect Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetPasswordCorrect() int64 {
	return r._passwordCorrect
}

// SetRegisterDate is RegisterDate Setter
// 注册的时间（秒）
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetRegisterDate(_registerDate int64) error {
	r._registerDate = _registerDate
	r.Set("register_date", _registerDate)
	return nil
}

// GetRegisterDate RegisterDate Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetRegisterDate() int64 {
	return r._registerDate
}

// SetSource is Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetSource() int64 {
	return r._source
}

var poolAlibabaSecurityJaqLoginpreventionResultFetchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqLoginpreventionResultFetchRequest()
	},
}

// GetAlibabaSecurityJaqLoginpreventionResultFetchRequest 从 sync.Pool 获取 AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest
func GetAlibabaSecurityJaqLoginpreventionResultFetchAPIRequest() *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest {
	return poolAlibabaSecurityJaqLoginpreventionResultFetchAPIRequest.Get().(*AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest)
}

// ReleaseAlibabaSecurityJaqLoginpreventionResultFetchAPIRequest 将 AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqLoginpreventionResultFetchAPIRequest(v *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqLoginpreventionResultFetchAPIRequest.Put(v)
}
