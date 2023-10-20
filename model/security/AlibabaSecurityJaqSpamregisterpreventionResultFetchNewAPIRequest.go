package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest 获取虚假注册保护结果 API请求
// alibaba.security.jaq.spamregisterprevention.result.fetch.new
//
// 获取虚假注册保护结果
type AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest struct {
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
	_registerUrl string
	// 关联账号的email。 【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 扩展字段。json格式的字符串，根据具体情况而定
	_extendData string
	// 登录时候的IP地址
	_ip string
	// 接入JS后获取的token
	_jsToken string
	// 硬件信息
	_macAddress string
	// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 协议版本号。现在的值是1.0
	_protocolVersion string
	// 上一跳信息
	_referer string
	// Session id
	_sessionId string
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 账号绑定的呢称
	_nickName string
	// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
	_idType int64
	// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
	_source int64
}

// NewAlibabasecurityjaqspamregisterpreventionresultfetchnewRequest 初始化AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest对象
func NewAlibabasecurityjaqspamregisterpreventionresultfetchnewRequest() *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest {
	return &AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.spamregisterprevention.result.fetch.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdNumber is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetSdkToken is SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// GetSdkToken SdkToken Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// SetAddress is Address Setter
// 账号绑定的注册的地址
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetAddress() string {
	return r._address
}

// SetAgent is Agent Setter
// 发送HTTP请求的代理
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// GetAgent Agent Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetAgent() string {
	return r._agent
}

// SetBankCardNumber is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// GetBankCardNumber BankCardNumber Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// SetCompanyName is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetCookie is Cookie Setter
// Cookie
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// GetCookie Cookie Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetCookie() string {
	return r._cookie
}

// SetRegisterUrl is RegisterUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetRegisterUrl(_registerUrl string) error {
	r._registerUrl = _registerUrl
	r.Set("register_url", _registerUrl)
	return nil
}

// GetRegisterUrl RegisterUrl Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetRegisterUrl() string {
	return r._registerUrl
}

// SetEmail is Email Setter
// 关联账号的email。 【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetEmail() string {
	return r._email
}

// SetExtendData is ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// GetExtendData ExtendData Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetExtendData() string {
	return r._extendData
}

// SetIp is Ip Setter
// 登录时候的IP地址
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetIp() string {
	return r._ip
}

// SetJsToken is JsToken Setter
// 接入JS后获取的token
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// GetJsToken JsToken Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetJsToken() string {
	return r._jsToken
}

// SetMacAddress is MacAddress Setter
// 硬件信息
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// GetMacAddress MacAddress Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// SetPhoneNumber is PhoneNumber Setter
// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// GetPhoneNumber PhoneNumber Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetReferer is Referer Setter
// 上一跳信息
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// GetReferer Referer Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetReferer() string {
	return r._referer
}

// SetSessionId is SessionId Setter
// Session id
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetSessionId() string {
	return r._sessionId
}

// SetUserId is UserId Setter
// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetUserId() string {
	return r._userId
}

// SetNickName is NickName Setter
// 账号绑定的呢称
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetNickName() string {
	return r._nickName
}

// SetIdType is IdType Setter
// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// GetIdType IdType Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetIdType() int64 {
	return r._idType
}

// SetSource is Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabasecurityjaqspamregisterpreventionresultfetchnewAPIRequest) GetSource() int64 {
	return r._source
}
