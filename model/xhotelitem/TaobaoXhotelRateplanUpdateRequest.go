package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan更新或添加 APIRequest
taobao.xhotel.rateplan.update

酒店产品库rateplan更新或添加
*/
type TaobaoXhotelRateplanUpdateRequest struct {
    model.Params

    // 在淘宝搜索页面展示的房价名称；（添加RP时为必须）。注意该名称不要包含早餐相关信息，如果想维护早餐信息，请设置breakfast_count字段即可。
    name   string 

    // -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）
    breakfastCount   int64 

    // 最小入住天数（1-90）。默认1,小时房RP请设置为1
    minDays   int64 

    // 最大入住天数（1-90）。默认90,信用住最大入住天数不超过9天,小时房RP请设置为1
    maxDays   int64 

    // 首日入住房间数（1-99）。默认1。【废弃】
    minAmount   int64 

    // 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
    minAdvHours   int64 

    // 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
    maxAdvHours   int64 

    // 产品每日开始销售时间，start_time一定为当天时间
    startTime   string 

    // 产品每日开始销售时间，start_time一定为当天时间
    endTime   string 

    // 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款
    cancelPolicy   string 

    // 1：开启（默认）2：关闭。如果没传值那么默认默认值为1；（添加RP时为必须）
    status   int64 

    // RP的英文名称
    englishName   string 

    // 担保类型，只支持： 0 无担保 1 峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
    guaranteeType   int64 

    // 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
    guaranteeStartTime   string 

    // 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
    memberLevel   string 

    // 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪 O:钉钉商旅 A:集团内部商旅 M:无线专享价。如果只投放飞猪，改字段不用填写或者只填H；如果有多个用","分开。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
    channel   string 

    // 系统商，一般不用填写，使用须申请
    vendor   string 

    // 商家价格政策编码
    rateplanCode   string 

    // 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0-代表否。不填写代表不更新该字段。
    firstStay   int64 

    // 价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
    agreement   int64 

    // 在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。>=0,<=99 如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
    breakfastCal   string 

    // 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
    cancelPolicyCal   string 

    // 在更新rateplan的同时，新增或更新担保日历。 date：担保日历的某一天。 guarantee:担保政策。 其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm" 。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点） 请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
    guaranteeCal   string 

    // 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
    effectiveTime   string 

    // 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
    deadlineTime   string 

    // 支付类型，只支持：1：预付5：现付6: 信用住7:在线预约8:在线预约信用住。其中5,6,7,8三种类型需要申请权限
    paymentType   int64 

    // rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
    rpType   string 

    // 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
    hourage   string 

    // 最晚可选入住时间，小时房特有字段。格式为HH:mm
    canCheckinEnd   string 

    // 最早可选入住时间，小时房特有字段。格式为HH:mm
    canCheckinStart   string 

    // 学生价，1：是；0：否
    isStudent   int64 

    // 酒店id
    hid   int64 

    // 房型id
    rid   int64 

    // 外部房型id
    outRid   string 

    // 外部酒店id
    outHid   string 

    // super rp标记，1是；0否
    superRpFlag   int64 

    // base rp标记，1是；0否
    baseRpFlag   int64 

    // -9999 清空担保, 2 VCC担保, 1 PCI担保，0 支付宝担保(信用住产品担保方式只能为支付宝担保)
    guaranteeMode   int64 

    // operator
    operator   string 

    // 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
    parentRpCode   string 

    // 父rpid，使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
    parentRpid   int64 

    // 更新RP时的 打标和去标 需求, 0 就是 去标, 1 就是打标,    key的含义:    non-direct-RP 表示非直连RP,,     super-could-price-change-RP 表示rp的super标，打上这个tag，表明这个rateplan下单的时候支持变价，商家系统直接放开价格校验。   base-could-derived-RP 表示base rateplan标签，打上了这个tag，表明这是一个base的rateplan，基于该rateplan可以衍生出子rateplan,       ebk-tail-room-RP 表示 ebk尾房rate plan级别标
    tagJson   string 

    // 协议保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
    allotmentReleaseTime   string 

    // 是否包房RP 1包房RP,0 非包房rp
    packRoomFlag   string 

    // 是否底价加价，1是底价加价,0 非底价加价rp
    bottomPriceFlag   string 

    // 价格计划名称name通过加工处理以后的值
    displayName   string 

    // 来源
    source   int64 

    // 普通保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
    commonAllotReleaseTime   string 

    // 是否企业托管RP 0-普通rp,1-企业托管rp
    companyAssist   int64 

    // 酒店-企业-rp映射实体集合
    hotelCompanyMappingDOS   string 

    // 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
    resourceType   string 

    // 最晚可选离店时间，小时房特有字段。格式为HH:mm
    canCheckoutEnd   string 

    // 会员价支持标识,1表示支持会员价规则
    memDiscFlag   int64 

    // 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)。该字段使用需要联系小二
    memberDiscountCal   string 

    // 卖点
    benefits   string 

    // 活动类型。1通兑, 4:超级房券
    activityType   string 

    // 入住人限制
    guestLimit   string 

    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    onlineBookingBindingInfo   string 

    // 不推荐使用，使用ratePlanCode来标识要修改的RP
    rpid   int64 

}

func NewTaobaoXhotelRateplanUpdateRequest() *TaobaoXhotelRateplanUpdateRequest{
    return &TaobaoXhotelRateplanUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateplanUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.update"
}

func (r TaobaoXhotelRateplanUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateplanUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetBreakfastCount(breakfastCount int64) error {
    r.breakfastCount = breakfastCount
    r.Set("breakfast_count", breakfastCount)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetBreakfastCount() int64 {
    return r.breakfastCount
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMinDays(minDays int64) error {
    r.minDays = minDays
    r.Set("min_days", minDays)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMinDays() int64 {
    return r.minDays
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMaxDays(maxDays int64) error {
    r.maxDays = maxDays
    r.Set("max_days", maxDays)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMaxDays() int64 {
    return r.maxDays
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMinAmount(minAmount int64) error {
    r.minAmount = minAmount
    r.Set("min_amount", minAmount)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMinAmount() int64 {
    return r.minAmount
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMinAdvHours(minAdvHours int64) error {
    r.minAdvHours = minAdvHours
    r.Set("min_adv_hours", minAdvHours)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMinAdvHours() int64 {
    return r.minAdvHours
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMaxAdvHours(maxAdvHours int64) error {
    r.maxAdvHours = maxAdvHours
    r.Set("max_adv_hours", maxAdvHours)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMaxAdvHours() int64 {
    return r.maxAdvHours
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCancelPolicy(cancelPolicy string) error {
    r.cancelPolicy = cancelPolicy
    r.Set("cancel_policy", cancelPolicy)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCancelPolicy() string {
    return r.cancelPolicy
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetEnglishName(englishName string) error {
    r.englishName = englishName
    r.Set("english_name", englishName)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetEnglishName() string {
    return r.englishName
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeType(guaranteeType int64) error {
    r.guaranteeType = guaranteeType
    r.Set("guarantee_type", guaranteeType)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetGuaranteeType() int64 {
    return r.guaranteeType
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeStartTime(guaranteeStartTime string) error {
    r.guaranteeStartTime = guaranteeStartTime
    r.Set("guarantee_start_time", guaranteeStartTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetGuaranteeStartTime() string {
    return r.guaranteeStartTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMemberLevel(memberLevel string) error {
    r.memberLevel = memberLevel
    r.Set("member_level", memberLevel)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMemberLevel() string {
    return r.memberLevel
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetChannel() string {
    return r.channel
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetRateplanCode() string {
    return r.rateplanCode
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetFirstStay(firstStay int64) error {
    r.firstStay = firstStay
    r.Set("first_stay", firstStay)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetFirstStay() int64 {
    return r.firstStay
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetAgreement(agreement int64) error {
    r.agreement = agreement
    r.Set("agreement", agreement)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetAgreement() int64 {
    return r.agreement
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetBreakfastCal(breakfastCal string) error {
    r.breakfastCal = breakfastCal
    r.Set("breakfast_cal", breakfastCal)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetBreakfastCal() string {
    return r.breakfastCal
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCancelPolicyCal(cancelPolicyCal string) error {
    r.cancelPolicyCal = cancelPolicyCal
    r.Set("cancel_policy_cal", cancelPolicyCal)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCancelPolicyCal() string {
    return r.cancelPolicyCal
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeCal(guaranteeCal string) error {
    r.guaranteeCal = guaranteeCal
    r.Set("guarantee_cal", guaranteeCal)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetGuaranteeCal() string {
    return r.guaranteeCal
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetEffectiveTime(effectiveTime string) error {
    r.effectiveTime = effectiveTime
    r.Set("effective_time", effectiveTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetEffectiveTime() string {
    return r.effectiveTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetDeadlineTime(deadlineTime string) error {
    r.deadlineTime = deadlineTime
    r.Set("deadline_time", deadlineTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetDeadlineTime() string {
    return r.deadlineTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetPaymentType(paymentType int64) error {
    r.paymentType = paymentType
    r.Set("payment_type", paymentType)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetPaymentType() int64 {
    return r.paymentType
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetRpType(rpType string) error {
    r.rpType = rpType
    r.Set("rp_type", rpType)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetRpType() string {
    return r.rpType
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetHourage(hourage string) error {
    r.hourage = hourage
    r.Set("hourage", hourage)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetHourage() string {
    return r.hourage
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCanCheckinEnd(canCheckinEnd string) error {
    r.canCheckinEnd = canCheckinEnd
    r.Set("can_checkin_end", canCheckinEnd)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCanCheckinEnd() string {
    return r.canCheckinEnd
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCanCheckinStart(canCheckinStart string) error {
    r.canCheckinStart = canCheckinStart
    r.Set("can_checkin_start", canCheckinStart)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCanCheckinStart() string {
    return r.canCheckinStart
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetIsStudent(isStudent int64) error {
    r.isStudent = isStudent
    r.Set("is_student", isStudent)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetIsStudent() int64 {
    return r.isStudent
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetHid() int64 {
    return r.hid
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetRid(rid int64) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetRid() int64 {
    return r.rid
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetOutHid() string {
    return r.outHid
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetSuperRpFlag(superRpFlag int64) error {
    r.superRpFlag = superRpFlag
    r.Set("super_rp_flag", superRpFlag)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetSuperRpFlag() int64 {
    return r.superRpFlag
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetBaseRpFlag(baseRpFlag int64) error {
    r.baseRpFlag = baseRpFlag
    r.Set("base_rp_flag", baseRpFlag)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetBaseRpFlag() int64 {
    return r.baseRpFlag
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeMode(guaranteeMode int64) error {
    r.guaranteeMode = guaranteeMode
    r.Set("guarantee_mode", guaranteeMode)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetGuaranteeMode() int64 {
    return r.guaranteeMode
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetOperator() string {
    return r.operator
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetParentRpCode(parentRpCode string) error {
    r.parentRpCode = parentRpCode
    r.Set("parent_rp_code", parentRpCode)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetParentRpCode() string {
    return r.parentRpCode
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetParentRpid(parentRpid int64) error {
    r.parentRpid = parentRpid
    r.Set("parent_rpid", parentRpid)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetParentRpid() int64 {
    return r.parentRpid
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetTagJson(tagJson string) error {
    r.tagJson = tagJson
    r.Set("tag_json", tagJson)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetTagJson() string {
    return r.tagJson
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetAllotmentReleaseTime(allotmentReleaseTime string) error {
    r.allotmentReleaseTime = allotmentReleaseTime
    r.Set("allotment_release_time", allotmentReleaseTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetAllotmentReleaseTime() string {
    return r.allotmentReleaseTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetPackRoomFlag(packRoomFlag string) error {
    r.packRoomFlag = packRoomFlag
    r.Set("pack_room_flag", packRoomFlag)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetPackRoomFlag() string {
    return r.packRoomFlag
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetBottomPriceFlag(bottomPriceFlag string) error {
    r.bottomPriceFlag = bottomPriceFlag
    r.Set("bottom_price_flag", bottomPriceFlag)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetBottomPriceFlag() string {
    return r.bottomPriceFlag
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetDisplayName(displayName string) error {
    r.displayName = displayName
    r.Set("display_name", displayName)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetDisplayName() string {
    return r.displayName
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetSource() int64 {
    return r.source
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCommonAllotReleaseTime(commonAllotReleaseTime string) error {
    r.commonAllotReleaseTime = commonAllotReleaseTime
    r.Set("common_allot_release_time", commonAllotReleaseTime)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCommonAllotReleaseTime() string {
    return r.commonAllotReleaseTime
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCompanyAssist(companyAssist int64) error {
    r.companyAssist = companyAssist
    r.Set("company_assist", companyAssist)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCompanyAssist() int64 {
    return r.companyAssist
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetHotelCompanyMappingDOS(hotelCompanyMappingDOS string) error {
    r.hotelCompanyMappingDOS = hotelCompanyMappingDOS
    r.Set("hotel_company_mapping_d_o_s", hotelCompanyMappingDOS)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetHotelCompanyMappingDOS() string {
    return r.hotelCompanyMappingDOS
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetResourceType(resourceType string) error {
    r.resourceType = resourceType
    r.Set("resource_type", resourceType)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetResourceType() string {
    return r.resourceType
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetCanCheckoutEnd(canCheckoutEnd string) error {
    r.canCheckoutEnd = canCheckoutEnd
    r.Set("can_checkout_end", canCheckoutEnd)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetCanCheckoutEnd() string {
    return r.canCheckoutEnd
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMemDiscFlag(memDiscFlag int64) error {
    r.memDiscFlag = memDiscFlag
    r.Set("mem_disc_flag", memDiscFlag)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMemDiscFlag() int64 {
    return r.memDiscFlag
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetMemberDiscountCal(memberDiscountCal string) error {
    r.memberDiscountCal = memberDiscountCal
    r.Set("member_discount_cal", memberDiscountCal)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetMemberDiscountCal() string {
    return r.memberDiscountCal
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetBenefits(benefits string) error {
    r.benefits = benefits
    r.Set("benefits", benefits)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetBenefits() string {
    return r.benefits
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetActivityType(activityType string) error {
    r.activityType = activityType
    r.Set("activity_type", activityType)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetActivityType() string {
    return r.activityType
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetGuestLimit(guestLimit string) error {
    r.guestLimit = guestLimit
    r.Set("guest_limit", guestLimit)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetGuestLimit() string {
    return r.guestLimit
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetOnlineBookingBindingInfo(onlineBookingBindingInfo string) error {
    r.onlineBookingBindingInfo = onlineBookingBindingInfo
    r.Set("online_booking_binding_info", onlineBookingBindingInfo)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetOnlineBookingBindingInfo() string {
    return r.onlineBookingBindingInfo
}

func (r *TaobaoXhotelRateplanUpdateRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

func (r TaobaoXhotelRateplanUpdateRequest) GetRpid() int64 {
    return r.rpid
}

