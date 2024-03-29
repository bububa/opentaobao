package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest 获取垃圾注册防控结果 API请求
// alibaba.security.jaq.spamregisterprevention.result.fetch
//
// 获取垃圾注册防控结果
type AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest struct {
	model.Params
	// 账号绑定的手机号码
	_phoneNumber string
	// 账号绑定的IP地址
	_ip string
	// 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
	_context string
	// 账号绑定的email地址
	_email string
	// 账号的全局唯一标识
	_userId string
	// 当前注册的页面URL，Source为1，2时，该参数必选
	_registerUrl string
	// agent，发送HTTP请求的代理
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
	_nickName string
	// 账号绑定的公司名字
	_companyName string
	// 账号绑定的注册的地址
	_address string
	// 账号绑定的身份证号
	_idNumber string
	// 账号绑定的银行卡号
	_bankCardNumber string
	// 接入JS后从服务端获取的token
	_jsToken string
	// 接入SDK后从服务端获取的token
	_sdkToken string
	// 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
	_source int64
}

// NewAlibabaSecurityJaqSpamregisterpreventionResultFetchRequest 初始化AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest对象
func NewAlibabaSecurityJaqSpamregisterpreventionResultFetchRequest() *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest {
	return &AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest{
		Params: model.NewParams(19),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) Reset() {
	r._phoneNumber = ""
	r._ip = ""
	r._context = ""
	r._email = ""
	r._userId = ""
	r._registerUrl = ""
	r._agent = ""
	r._cookie = ""
	r._sessionId = ""
	r._macAddress = ""
	r._referer = ""
	r._nickName = ""
	r._companyName = ""
	r._address = ""
	r._idNumber = ""
	r._bankCardNumber = ""
	r._jsToken = ""
	r._sdkToken = ""
	r._source = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.spamregisterprevention.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNumber is PhoneNumber Setter
// 账号绑定的手机号码
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// GetPhoneNumber PhoneNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// SetIp is Ip Setter
// 账号绑定的IP地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetIp() string {
	return r._ip
}

// SetContext is Context Setter
// 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetContext(_context string) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetContext() string {
	return r._context
}

// SetEmail is Email Setter
// 账号绑定的email地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetEmail() string {
	return r._email
}

// SetUserId is UserId Setter
// 账号的全局唯一标识
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetUserId() string {
	return r._userId
}

// SetRegisterUrl is RegisterUrl Setter
// 当前注册的页面URL，Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetRegisterUrl(_registerUrl string) error {
	r._registerUrl = _registerUrl
	r.Set("register_url", _registerUrl)
	return nil
}

// GetRegisterUrl RegisterUrl Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetRegisterUrl() string {
	return r._registerUrl
}

// SetAgent is Agent Setter
// agent，发送HTTP请求的代理
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// GetAgent Agent Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetAgent() string {
	return r._agent
}

// SetCookie is Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// GetCookie Cookie Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetCookie() string {
	return r._cookie
}

// SetSessionId is SessionId Setter
// Session id
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetSessionId() string {
	return r._sessionId
}

// SetMacAddress is MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// GetMacAddress MacAddress Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// SetReferer is Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// GetReferer Referer Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetReferer() string {
	return r._referer
}

// SetNickName is NickName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetNickName() string {
	return r._nickName
}

// SetCompanyName is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetAddress is Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetAddress() string {
	return r._address
}

// SetIdNumber is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetBankCardNumber is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// GetBankCardNumber BankCardNumber Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// SetJsToken is JsToken Setter
// 接入JS后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// GetJsToken JsToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetJsToken() string {
	return r._jsToken
}

// SetSdkToken is SdkToken Setter
// 接入SDK后从服务端获取的token
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// GetSdkToken SdkToken Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// SetSource is Source Setter
// 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
func (r *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) GetSource() int64 {
	return r._source
}

var poolAlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqSpamregisterpreventionResultFetchRequest()
	},
}

// GetAlibabaSecurityJaqSpamregisterpreventionResultFetchRequest 从 sync.Pool 获取 AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest
func GetAlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest() *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest {
	return poolAlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest.Get().(*AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest)
}

// ReleaseAlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest 将 AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest(v *AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest.Put(v)
}
