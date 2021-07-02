package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest 获取活动保护结果 API请求
// alibaba.security.jaq.campaignprevention.result.fetch
//
// 获取活动保护结果
type AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest struct {
	model.Params
	// 电话号码。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 邮箱地址。【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 账号的全局唯一标识。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
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
	// 活动描述，场景为活动时提供，活动基本规则描述，以评估活动流程的风险
	_activityDescription string
	// 活动编号
	_activityId string
	// 活动奖品面值，场景为活动时提供，活动抽中的价值
	_prize string
	// 活动奖品类型。1：事物；2：红包；3：优惠券；4：购物券；5：积分；6：代币；0：其它
	_prizeType int64
	// 注册的时间（秒）
	_registerDate int64
	// 注册时候的ip
	_registerIp string
}

// NewAlibabaSecurityJaqCampaignpreventionResultFetchRequest 初始化AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest对象
func NewAlibabaSecurityJaqCampaignpreventionResultFetchRequest() *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest {
	return &AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.campaignprevention.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PhoneNumber Setter
// 电话号码。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// Get PhoneNumber Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// Set is Email Setter
// 邮箱地址。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// Get Email Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetEmail() string {
	return r._email
}

// Set is UserId Setter
// 账号的全局唯一标识。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetUserId() string {
	return r._userId
}

// Set is IdType Setter
// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// Get IdType Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetIdType() int64 {
	return r._idType
}

// Set is Ip Setter
// 登录时候的IP地址
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// Get Ip Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetIp() string {
	return r._ip
}

// Set is Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetSource() int64 {
	return r._source
}

// Set is CurrentUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetCurrentUrl(_currentUrl string) error {
	r._currentUrl = _currentUrl
	r.Set("current_url", _currentUrl)
	return nil
}

// Get CurrentUrl Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetCurrentUrl() string {
	return r._currentUrl
}

// Set is Agent Setter
// 发送HTTP请求的代理
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// Get Agent Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetAgent() string {
	return r._agent
}

// Set is Cookie Setter
// Cookie
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// Get Cookie Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetCookie() string {
	return r._cookie
}

// Set is SessionId Setter
// Session id
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// Get SessionId Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetSessionId() string {
	return r._sessionId
}

// Set is MacAddress Setter
// 硬件信息
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// Get MacAddress Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// Set is Referer Setter
// 上一跳信息
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// Get Referer Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetReferer() string {
	return r._referer
}

// Set is UserName Setter
// 账号绑定的呢称
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// Get UserName Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetUserName() string {
	return r._userName
}

// Set is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// Get CompanyName Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetCompanyName() string {
	return r._companyName
}

// Set is Address Setter
// 账号绑定的注册的地址
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetAddress() string {
	return r._address
}

// Set is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// Get IdNumber Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// Set is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// Get BankCardNumber Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// Set is JsToken Setter
// 接入JS后获取的token
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// Get JsToken Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetJsToken() string {
	return r._jsToken
}

// Set is SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// Get SdkToken Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// Set is ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// Get ProtocolVersion Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// Set is ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定 。
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// Get ExtendData Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetExtendData() string {
	return r._extendData
}

// Set is ActivityDescription Setter
// 活动描述，场景为活动时提供，活动基本规则描述，以评估活动流程的风险
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetActivityDescription(_activityDescription string) error {
	r._activityDescription = _activityDescription
	r.Set("activity_description", _activityDescription)
	return nil
}

// Get ActivityDescription Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetActivityDescription() string {
	return r._activityDescription
}

// Set is ActivityId Setter
// 活动编号
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetActivityId() string {
	return r._activityId
}

// Set is Prize Setter
// 活动奖品面值，场景为活动时提供，活动抽中的价值
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetPrize(_prize string) error {
	r._prize = _prize
	r.Set("prize", _prize)
	return nil
}

// Get Prize Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetPrize() string {
	return r._prize
}

// Set is PrizeType Setter
// 活动奖品类型。1：事物；2：红包；3：优惠券；4：购物券；5：积分；6：代币；0：其它
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetPrizeType(_prizeType int64) error {
	r._prizeType = _prizeType
	r.Set("prize_type", _prizeType)
	return nil
}

// Get PrizeType Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetPrizeType() int64 {
	return r._prizeType
}

// Set is RegisterDate Setter
// 注册的时间（秒）
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetRegisterDate(_registerDate int64) error {
	r._registerDate = _registerDate
	r.Set("register_date", _registerDate)
	return nil
}

// Get RegisterDate Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetRegisterDate() int64 {
	return r._registerDate
}

// Set is RegisterIp Setter
// 注册时候的ip
func (r *AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) SetRegisterIp(_registerIp string) error {
	r._registerIp = _registerIp
	r.Set("register_ip", _registerIp)
	return nil
}

// Get RegisterIp Getter
func (r AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest) GetRegisterIp() string {
	return r._registerIp
}
