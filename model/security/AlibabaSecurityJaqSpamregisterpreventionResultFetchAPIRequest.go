package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest 获取垃圾注册防控结果 API请求
// alibaba.security.jaq.spamregisterprevention.result.fetch
//
// 获取垃圾注册防控结果
type AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest struct {
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

// NewAlibabasecurityjaqspamregisterpreventionresultfetchRequest 初始化AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest对象
func NewAlibabasecurityjaqspamregisterpreventionresultfetchRequest() *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest {
	return &AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.spamregisterprevention.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNumber is PhoneNumber Setter
// 账号绑定的手机号码
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// GetPhoneNumber PhoneNumber Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// SetIp is Ip Setter
// 账号绑定的IP地址
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetIp() string {
	return r._ip
}

// SetContext is Context Setter
// 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetContext(_context string) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetContext() string {
	return r._context
}

// SetEmail is Email Setter
// 账号绑定的email地址
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetEmail() string {
	return r._email
}

// SetUserId is UserId Setter
// 账号的全局唯一标识
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetUserId() string {
	return r._userId
}

// SetRegisterUrl is RegisterUrl Setter
// 当前注册的页面URL，Source为1，2时，该参数必选
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetRegisterUrl(_registerUrl string) error {
	r._registerUrl = _registerUrl
	r.Set("register_url", _registerUrl)
	return nil
}

// GetRegisterUrl RegisterUrl Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetRegisterUrl() string {
	return r._registerUrl
}

// SetAgent is Agent Setter
// agent，发送HTTP请求的代理
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// GetAgent Agent Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetAgent() string {
	return r._agent
}

// SetCookie is Cookie Setter
// Cookie
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// GetCookie Cookie Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetCookie() string {
	return r._cookie
}

// SetSessionId is SessionId Setter
// Session id
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetSessionId() string {
	return r._sessionId
}

// SetMacAddress is MacAddress Setter
// 硬件信息
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// GetMacAddress MacAddress Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// SetReferer is Referer Setter
// 上一跳信息
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// GetReferer Referer Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetReferer() string {
	return r._referer
}

// SetNickName is NickName Setter
// 账号绑定的呢称
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetNickName() string {
	return r._nickName
}

// SetCompanyName is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetAddress is Address Setter
// 账号绑定的注册的地址
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetAddress() string {
	return r._address
}

// SetIdNumber is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetBankCardNumber is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// GetBankCardNumber BankCardNumber Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// SetJsToken is JsToken Setter
// 接入JS后从服务端获取的token
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// GetJsToken JsToken Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetJsToken() string {
	return r._jsToken
}

// SetSdkToken is SdkToken Setter
// 接入SDK后从服务端获取的token
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// GetSdkToken SdkToken Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// SetSource is Source Setter
// 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchAPIRequest) GetSource() int64 {
	return r._source
}
