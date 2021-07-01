package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest
获取登录保护结果 API请求
alibaba.security.jaq.loginprevention.result.fetch

获取登录保护结果 */
type AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest struct {
	model.Params
	// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
	_idType int64
	// 登录时候的IP地址
	_ip string
	// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
	_source int64
	// 当前操作的页面URL。Source为1，2时，该参数必选
	_currentUrl string
	// 发送HTTP请求的代理
	_agent string
	// Cookie
	_cookie string
	// Session id
	_sessionId string
	// 硬件信息
	_macAddress string
	// 上一跳信息
	_referer string
	// 账号绑定的呢称
	_userName string
	// 账号绑定的公司名字
	_companyName string
	// 账号绑定的注册的地址
	_address string
	// 账号绑定的身份证号
	_idNumber string
	// 账号绑定的银行卡号
	_bankCardNumber string
	// 接入JS后获取的token
	_jsToken string
	// 接入无线保镖安全组件后从服务端获取的token
	_sdkToken string
	// 协议版本号。现在的值是1.0
	_protocolVersion string
	// 扩展字段。json格式的字符串，根据具体情况而定 。
	_extendData string
	// 账号在系统里面是否存在。0：不存在；1：存在
	_accountExist int64
	// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
	_loginType int64
	// 密码是否正确。0：不正确；1：正确
	_passwordCorrect int64
	// 将密码加盐hash后传递，用于弱密码检测
	_passwordHash string
	// 注册的时间（秒）
	_registerDate int64
	// 注册时候的ip
	_registerIp string
}

// NewAlibabaSecurityJaqLoginpreventionResultFetchRequest 初始化AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest对象
func NewAlibabaSecurityJaqLoginpreventionResultFetchRequest() *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest {
	return &AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.loginprevention.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PhoneNumber Setter
// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// Get PhoneNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// Set is Email Setter
// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// Get Email Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetEmail() string {
	return r._email
}

// Set is UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetUserId() string {
	return r._userId
}

// Set is IdType Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// Get IdType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetIdType() int64 {
	return r._idType
}

// Set is Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// Get Ip Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetIp() string {
	return r._ip
}

// Set is Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetSource() int64 {
	return r._source
}

// Set is CurrentUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetCurrentUrl(_currentUrl string) error {
	r._currentUrl = _currentUrl
	r.Set("current_url", _currentUrl)
	return nil
}

// Get CurrentUrl Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetCurrentUrl() string {
	return r._currentUrl
}

// Set is Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// Get Agent Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetAgent() string {
	return r._agent
}

// Set is Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// Get Cookie Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetCookie() string {
	return r._cookie
}

// Set is SessionId Setter
// Session id
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// Get SessionId Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetSessionId() string {
	return r._sessionId
}

// Set is MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// Get MacAddress Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// Set is Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// Get Referer Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetReferer() string {
	return r._referer
}

// Set is UserName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// Get UserName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetUserName() string {
	return r._userName
}

// Set is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// Get CompanyName Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetCompanyName() string {
	return r._companyName
}

// Set is Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetAddress() string {
	return r._address
}

// Set is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// Get IdNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// Set is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// Get BankCardNumber Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// Set is JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// Get JsToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetJsToken() string {
	return r._jsToken
}

// Set is SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// Get SdkToken Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// Set is ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// Get ProtocolVersion Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// Set is ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定 。
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// Get ExtendData Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetExtendData() string {
	return r._extendData
}

// Set is AccountExist Setter
// 账号在系统里面是否存在。0：不存在；1：存在
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetAccountExist(_accountExist int64) error {
	r._accountExist = _accountExist
	r.Set("account_exist", _accountExist)
	return nil
}

// Get AccountExist Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetAccountExist() int64 {
	return r._accountExist
}

// Set is LoginType Setter
// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetLoginType(_loginType int64) error {
	r._loginType = _loginType
	r.Set("login_type", _loginType)
	return nil
}

// Get LoginType Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetLoginType() int64 {
	return r._loginType
}

// Set is PasswordCorrect Setter
// 密码是否正确。0：不正确；1：正确
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetPasswordCorrect(_passwordCorrect int64) error {
	r._passwordCorrect = _passwordCorrect
	r.Set("password_correct", _passwordCorrect)
	return nil
}

// Get PasswordCorrect Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetPasswordCorrect() int64 {
	return r._passwordCorrect
}

// Set is PasswordHash Setter
// 将密码加盐hash后传递，用于弱密码检测
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetPasswordHash(_passwordHash string) error {
	r._passwordHash = _passwordHash
	r.Set("password_hash", _passwordHash)
	return nil
}

// Get PasswordHash Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetPasswordHash() string {
	return r._passwordHash
}

// Set is RegisterDate Setter
// 注册的时间（秒）
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetRegisterDate(_registerDate int64) error {
	r._registerDate = _registerDate
	r.Set("register_date", _registerDate)
	return nil
}

// Get RegisterDate Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetRegisterDate() int64 {
	return r._registerDate
}

// Set is RegisterIp Setter
// 注册时候的ip
func (r *AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) SetRegisterIp(_registerIp string) error {
	r._registerIp = _registerIp
	r.Set("register_ip", _registerIp)
	return nil
}

// Get RegisterIp Getter
func (r AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest) GetRegisterIp() string {
	return r._registerIp
}
