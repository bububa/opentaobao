package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动保护结果 APIRequest
alibaba.security.jaq.campaignprevention.result.fetch

获取活动保护结果
*/
type AlibabaSecurityJaqCampaignpreventionResultFetchRequest struct {
    model.Params

    // 电话号码。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
    phoneNumber   string 

    // 邮箱地址。【phone_number ,email, (user_id,id_type)三种必选其一】
    email   string 

    // 账号的全局唯一标识。【phone_number ,email, (user_id,id_type)三种必选其一】
    userId   string 

    // 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
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

    // 活动描述，场景为活动时提供，活动基本规则描述，以评估活动流程的风险
    activityDescription   string 

    // 活动编号
    activityId   string 

    // 活动奖品面值，场景为活动时提供，活动抽中的价值
    prize   string 

    // 活动奖品类型。1：事物；2：红包；3：优惠券；4：购物券；5：积分；6：代币；0：其它
    prizeType   int64 

    // 注册的时间（秒）
    registerDate   int64 

    // 注册时候的ip
    registerIp   string 

}

func NewAlibabaSecurityJaqCampaignpreventionResultFetchRequest() *AlibabaSecurityJaqCampaignpreventionResultFetchRequest{
    return &AlibabaSecurityJaqCampaignpreventionResultFetchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.campaignprevention.result.fetch"
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetPhoneNumber() string {
    return r.phoneNumber
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetEmail() string {
    return r.email
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetIdType() int64 {
    return r.idType
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetIp() string {
    return r.ip
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetSource() int64 {
    return r.source
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetCurrentUrl(currentUrl string) error {
    r.currentUrl = currentUrl
    r.Set("current_url", currentUrl)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetCurrentUrl() string {
    return r.currentUrl
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetAgent(agent string) error {
    r.agent = agent
    r.Set("agent", agent)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetAgent() string {
    return r.agent
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetCookie(cookie string) error {
    r.cookie = cookie
    r.Set("cookie", cookie)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetCookie() string {
    return r.cookie
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetSessionId() string {
    return r.sessionId
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetMacAddress(macAddress string) error {
    r.macAddress = macAddress
    r.Set("mac_address", macAddress)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetMacAddress() string {
    return r.macAddress
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetReferer(referer string) error {
    r.referer = referer
    r.Set("referer", referer)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetReferer() string {
    return r.referer
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetUserName() string {
    return r.userName
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetCompanyName() string {
    return r.companyName
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetIdNumber() string {
    return r.idNumber
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetBankCardNumber(bankCardNumber string) error {
    r.bankCardNumber = bankCardNumber
    r.Set("bank_card_number", bankCardNumber)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetBankCardNumber() string {
    return r.bankCardNumber
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetJsToken(jsToken string) error {
    r.jsToken = jsToken
    r.Set("js_token", jsToken)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetJsToken() string {
    return r.jsToken
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetSdkToken(sdkToken string) error {
    r.sdkToken = sdkToken
    r.Set("sdk_token", sdkToken)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetSdkToken() string {
    return r.sdkToken
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetProtocolVersion() string {
    return r.protocolVersion
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetExtendData(extendData string) error {
    r.extendData = extendData
    r.Set("extend_data", extendData)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetExtendData() string {
    return r.extendData
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetActivityDescription(activityDescription string) error {
    r.activityDescription = activityDescription
    r.Set("activity_description", activityDescription)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetActivityDescription() string {
    return r.activityDescription
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetActivityId() string {
    return r.activityId
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetPrize(prize string) error {
    r.prize = prize
    r.Set("prize", prize)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetPrize() string {
    return r.prize
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetPrizeType(prizeType int64) error {
    r.prizeType = prizeType
    r.Set("prize_type", prizeType)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetPrizeType() int64 {
    return r.prizeType
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetRegisterDate(registerDate int64) error {
    r.registerDate = registerDate
    r.Set("register_date", registerDate)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetRegisterDate() int64 {
    return r.registerDate
}

func (r *AlibabaSecurityJaqCampaignpreventionResultFetchRequest) SetRegisterIp(registerIp string) error {
    r.registerIp = registerIp
    r.Set("register_ip", registerIp)
    return nil
}

func (r AlibabaSecurityJaqCampaignpreventionResultFetchRequest) GetRegisterIp() string {
    return r.registerIp
}

