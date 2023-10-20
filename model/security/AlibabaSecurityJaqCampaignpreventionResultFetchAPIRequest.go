package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest 获取活动保护结果 API请求
// alibaba.security.jaq.campaignprevention.result.fetch
//
// 获取活动保护结果
type AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest struct {
	model.Params
	// 账号绑定的身份证号
	_idNumber string
	// 接入无线保镖安全组件后从服务端获取的token
	_sdkToken string
	// 活动描述，场景为活动时提供，活动基本规则描述，以评估活动流程的风险
	_activityDescription string
	// 活动编号
	_activityId string
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
	// 邮箱地址。【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 扩展字段。json格式的字符串，根据具体情况而定 。
	_extendData string
	// 登录时候的IP地址
	_ip string
	// 接入JS后获取的token
	_jsToken string
	// 硬件信息
	_macAddress string
	// 电话号码。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 活动奖品面值，场景为活动时提供，活动抽中的价值
	_prize string
	// 协议版本号。现在的值是1.0
	_protocolVersion string
	// 上一跳信息
	_referer string
	// 注册时候的ip
	_registerIp string
	// Session id
	_sessionId string
	// 账号的全局唯一标识。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 账号绑定的呢称
	_userName string
	// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
	_idType int64
	// 活动奖品类型。1：事物；2：红包；3：优惠券；4：购物券；5：积分；6：代币；0：其它
	_prizeType int64
	// 注册的时间（秒）
	_registerDate int64
	// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
	_source int64
}

// NewAlibabasecurityjaqcampaignpreventionresultfetchRequest 初始化AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest对象
func NewAlibabasecurityjaqcampaignpreventionresultfetchRequest() *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest {
	return &AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.campaignprevention.result.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdNumber is IdNumber Setter
// 账号绑定的身份证号
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetSdkToken is SdkToken Setter
// 接入无线保镖安全组件后从服务端获取的token
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetSdkToken(_sdkToken string) error {
	r._sdkToken = _sdkToken
	r.Set("sdk_token", _sdkToken)
	return nil
}

// GetSdkToken SdkToken Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetSdkToken() string {
	return r._sdkToken
}

// SetActivityDescription is ActivityDescription Setter
// 活动描述，场景为活动时提供，活动基本规则描述，以评估活动流程的风险
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetActivityDescription(_activityDescription string) error {
	r._activityDescription = _activityDescription
	r.Set("activity_description", _activityDescription)
	return nil
}

// GetActivityDescription ActivityDescription Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetActivityDescription() string {
	return r._activityDescription
}

// SetActivityId is ActivityId Setter
// 活动编号
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetAddress is Address Setter
// 账号绑定的注册的地址
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetAddress() string {
	return r._address
}

// SetAgent is Agent Setter
// 发送HTTP请求的代理
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetAgent(_agent string) error {
	r._agent = _agent
	r.Set("agent", _agent)
	return nil
}

// GetAgent Agent Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetAgent() string {
	return r._agent
}

// SetBankCardNumber is BankCardNumber Setter
// 账号绑定的银行卡号
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetBankCardNumber(_bankCardNumber string) error {
	r._bankCardNumber = _bankCardNumber
	r.Set("bank_card_number", _bankCardNumber)
	return nil
}

// GetBankCardNumber BankCardNumber Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetBankCardNumber() string {
	return r._bankCardNumber
}

// SetCompanyName is CompanyName Setter
// 账号绑定的公司名字
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetCookie is Cookie Setter
// Cookie
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetCookie(_cookie string) error {
	r._cookie = _cookie
	r.Set("cookie", _cookie)
	return nil
}

// GetCookie Cookie Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetCookie() string {
	return r._cookie
}

// SetCurrentUrl is CurrentUrl Setter
// 当前操作的页面URL。Source为1，2时，该参数必选
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetCurrentUrl(_currentUrl string) error {
	r._currentUrl = _currentUrl
	r.Set("current_url", _currentUrl)
	return nil
}

// GetCurrentUrl CurrentUrl Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetCurrentUrl() string {
	return r._currentUrl
}

// SetEmail is Email Setter
// 邮箱地址。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetEmail() string {
	return r._email
}

// SetExtendData is ExtendData Setter
// 扩展字段。json格式的字符串，根据具体情况而定 。
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetExtendData(_extendData string) error {
	r._extendData = _extendData
	r.Set("extend_data", _extendData)
	return nil
}

// GetExtendData ExtendData Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetExtendData() string {
	return r._extendData
}

// SetIp is Ip Setter
// 登录时候的IP地址
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetIp() string {
	return r._ip
}

// SetJsToken is JsToken Setter
// 接入JS后获取的token
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetJsToken(_jsToken string) error {
	r._jsToken = _jsToken
	r.Set("js_token", _jsToken)
	return nil
}

// GetJsToken JsToken Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetJsToken() string {
	return r._jsToken
}

// SetMacAddress is MacAddress Setter
// 硬件信息
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetMacAddress(_macAddress string) error {
	r._macAddress = _macAddress
	r.Set("mac_address", _macAddress)
	return nil
}

// GetMacAddress MacAddress Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetMacAddress() string {
	return r._macAddress
}

// SetPhoneNumber is PhoneNumber Setter
// 电话号码。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// GetPhoneNumber PhoneNumber Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// SetPrize is Prize Setter
// 活动奖品面值，场景为活动时提供，活动抽中的价值
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetPrize(_prize string) error {
	r._prize = _prize
	r.Set("prize", _prize)
	return nil
}

// GetPrize Prize Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetPrize() string {
	return r._prize
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本号。现在的值是1.0
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetReferer is Referer Setter
// 上一跳信息
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetReferer(_referer string) error {
	r._referer = _referer
	r.Set("referer", _referer)
	return nil
}

// GetReferer Referer Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetReferer() string {
	return r._referer
}

// SetRegisterIp is RegisterIp Setter
// 注册时候的ip
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetRegisterIp(_registerIp string) error {
	r._registerIp = _registerIp
	r.Set("register_ip", _registerIp)
	return nil
}

// GetRegisterIp RegisterIp Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetRegisterIp() string {
	return r._registerIp
}

// SetSessionId is SessionId Setter
// Session id
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetSessionId(_sessionId string) error {
	r._sessionId = _sessionId
	r.Set("session_id", _sessionId)
	return nil
}

// GetSessionId SessionId Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetSessionId() string {
	return r._sessionId
}

// SetUserId is UserId Setter
// 账号的全局唯一标识。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetUserId() string {
	return r._userId
}

// SetUserName is UserName Setter
// 账号绑定的呢称
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetUserName() string {
	return r._userName
}

// SetIdType is IdType Setter
// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// GetIdType IdType Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetIdType() int64 {
	return r._idType
}

// SetPrizeType is PrizeType Setter
// 活动奖品类型。1：事物；2：红包；3：优惠券；4：购物券；5：积分；6：代币；0：其它
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetPrizeType(_prizeType int64) error {
	r._prizeType = _prizeType
	r.Set("prize_type", _prizeType)
	return nil
}

// GetPrizeType PrizeType Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetPrizeType() int64 {
	return r._prizeType
}

// SetRegisterDate is RegisterDate Setter
// 注册的时间（秒）
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetRegisterDate(_registerDate int64) error {
	r._registerDate = _registerDate
	r.Set("register_date", _registerDate)
	return nil
}

// GetRegisterDate RegisterDate Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetRegisterDate() int64 {
	return r._registerDate
}

// SetSource is Source Setter
// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
func (r *AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabasecurityjaqcampaignpreventionresultfetchAPIRequest) GetSource() int64 {
	return r._source
}
