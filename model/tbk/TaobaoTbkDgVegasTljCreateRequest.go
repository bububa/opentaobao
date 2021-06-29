package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-淘礼金创建 API请求
taobao.tbk.dg.vegas.tlj.create

创建淘礼金
*/
type TaobaoTbkDgVegasTljCreateRequest struct {
    model.Params
    // CPS佣金计划类型
    campaignType   string
    // 妈妈广告位Id
    adzoneId   int64
    // 宝贝id
    itemId   int64
    // 淘礼金总个数
    totalNum   int64
    // 淘礼金名称，最大10个字符
    name   string
    // 单用户累计中奖次数上限
    userTotalWinNumLimit   int64
    // 安全开关，若不进行安全校验，可能放大您的资损风险，请谨慎选择
    securitySwitch   bool
    // 单个淘礼金面额，支持两位小数，单位元
    perFace   string
    // 发放开始时间
    sendStartTime   string
    // 发放截止时间
    sendEndTime   string
    // 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
    useEndTime   string
    // 结束日期的模式,1:相对时间，2:绝对时间
    useEndTimeMode   int64
    // 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
    useStartTime   string
    // 安全等级，0：适用于常规淘礼金投放场景；1：更高安全级别，适用于淘礼金面额偏大等安全性较高的淘礼金投放场景，可能导致更多用户拦截。security_switch为true，此字段不填写时，使用0作为默认安全级别。如果security_switch为false，不进行安全判断。
    securityLevel   int64
}

// 初始化TaobaoTbkDgVegasTljCreateRequest对象
func NewTaobaoTbkDgVegasTljCreateRequest() *TaobaoTbkDgVegasTljCreateRequest{
    return &TaobaoTbkDgVegasTljCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljCreateRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.vegas.tlj.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignType Setter
// CPS佣金计划类型
func (r *TaobaoTbkDgVegasTljCreateRequest) SetCampaignType(campaignType string) error {
    r.campaignType = campaignType
    r.Set("campaign_type", campaignType)
    return nil
}

// CampaignType Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetCampaignType() string {
    return r.campaignType
}
// AdzoneId Setter
// 妈妈广告位Id
func (r *TaobaoTbkDgVegasTljCreateRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetAdzoneId() int64 {
    return r.adzoneId
}
// ItemId Setter
// 宝贝id
func (r *TaobaoTbkDgVegasTljCreateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetItemId() int64 {
    return r.itemId
}
// TotalNum Setter
// 淘礼金总个数
func (r *TaobaoTbkDgVegasTljCreateRequest) SetTotalNum(totalNum int64) error {
    r.totalNum = totalNum
    r.Set("total_num", totalNum)
    return nil
}

// TotalNum Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetTotalNum() int64 {
    return r.totalNum
}
// Name Setter
// 淘礼金名称，最大10个字符
func (r *TaobaoTbkDgVegasTljCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetName() string {
    return r.name
}
// UserTotalWinNumLimit Setter
// 单用户累计中奖次数上限
func (r *TaobaoTbkDgVegasTljCreateRequest) SetUserTotalWinNumLimit(userTotalWinNumLimit int64) error {
    r.userTotalWinNumLimit = userTotalWinNumLimit
    r.Set("user_total_win_num_limit", userTotalWinNumLimit)
    return nil
}

// UserTotalWinNumLimit Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetUserTotalWinNumLimit() int64 {
    return r.userTotalWinNumLimit
}
// SecuritySwitch Setter
// 安全开关，若不进行安全校验，可能放大您的资损风险，请谨慎选择
func (r *TaobaoTbkDgVegasTljCreateRequest) SetSecuritySwitch(securitySwitch bool) error {
    r.securitySwitch = securitySwitch
    r.Set("security_switch", securitySwitch)
    return nil
}

// SecuritySwitch Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetSecuritySwitch() bool {
    return r.securitySwitch
}
// PerFace Setter
// 单个淘礼金面额，支持两位小数，单位元
func (r *TaobaoTbkDgVegasTljCreateRequest) SetPerFace(perFace string) error {
    r.perFace = perFace
    r.Set("per_face", perFace)
    return nil
}

// PerFace Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetPerFace() string {
    return r.perFace
}
// SendStartTime Setter
// 发放开始时间
func (r *TaobaoTbkDgVegasTljCreateRequest) SetSendStartTime(sendStartTime string) error {
    r.sendStartTime = sendStartTime
    r.Set("send_start_time", sendStartTime)
    return nil
}

// SendStartTime Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetSendStartTime() string {
    return r.sendStartTime
}
// SendEndTime Setter
// 发放截止时间
func (r *TaobaoTbkDgVegasTljCreateRequest) SetSendEndTime(sendEndTime string) error {
    r.sendEndTime = sendEndTime
    r.Set("send_end_time", sendEndTime)
    return nil
}

// SendEndTime Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetSendEndTime() string {
    return r.sendEndTime
}
// UseEndTime Setter
// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
func (r *TaobaoTbkDgVegasTljCreateRequest) SetUseEndTime(useEndTime string) error {
    r.useEndTime = useEndTime
    r.Set("use_end_time", useEndTime)
    return nil
}

// UseEndTime Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetUseEndTime() string {
    return r.useEndTime
}
// UseEndTimeMode Setter
// 结束日期的模式,1:相对时间，2:绝对时间
func (r *TaobaoTbkDgVegasTljCreateRequest) SetUseEndTimeMode(useEndTimeMode int64) error {
    r.useEndTimeMode = useEndTimeMode
    r.Set("use_end_time_mode", useEndTimeMode)
    return nil
}

// UseEndTimeMode Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetUseEndTimeMode() int64 {
    return r.useEndTimeMode
}
// UseStartTime Setter
// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
func (r *TaobaoTbkDgVegasTljCreateRequest) SetUseStartTime(useStartTime string) error {
    r.useStartTime = useStartTime
    r.Set("use_start_time", useStartTime)
    return nil
}

// UseStartTime Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetUseStartTime() string {
    return r.useStartTime
}
// SecurityLevel Setter
// 安全等级，0：适用于常规淘礼金投放场景；1：更高安全级别，适用于淘礼金面额偏大等安全性较高的淘礼金投放场景，可能导致更多用户拦截。security_switch为true，此字段不填写时，使用0作为默认安全级别。如果security_switch为false，不进行安全判断。
func (r *TaobaoTbkDgVegasTljCreateRequest) SetSecurityLevel(securityLevel int64) error {
    r.securityLevel = securityLevel
    r.Set("security_level", securityLevel)
    return nil
}

// SecurityLevel Getter
func (r TaobaoTbkDgVegasTljCreateRequest) GetSecurityLevel() int64 {
    return r.securityLevel
}
