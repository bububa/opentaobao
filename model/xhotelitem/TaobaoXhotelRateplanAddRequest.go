package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库rateplan添加 API请求
taobao.xhotel.rateplan.add

酒店产品库rateplan
*/
type TaobaoXhotelRateplanAddRequest struct {
    model.Params
    // 卖家自己系统的Code，简称RateCode
    _rateplanCode   string
    // 在淘宝搜索页面展示的房价名称。请注意名称里不要维护早餐信息，如果想设置早餐信息，请设置breakfast_count字段即可
    _name   string
    // 支付类型，只支持：1：预付5：现付6: 信用住7:预付在线预约8:信用住在线预约。其中5,6,7,8四种类型需要申请权限
    _paymentType   int64
    // -1：状态早餐,有具体几人价有关系，几人价是几份早餐;0：不含早1：含单早2：含双早N：含N早（-1-99可选）
    _breakfastCount   int64
    // 不推荐使用
    _feeServicePercent   int64
    // 最小入住天数（1-90）。默认1,小时房RP请设置为1
    _minDays   int64
    // 最大入住天数（1-90）。默认90 信用住不超过9天,小时房RP请设置为1
    _maxDays   int64
    // 首日入住房间数（1-99）。默认1。【废弃】
    _minAmount   int64
    // 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
    _minAdvHours   int64
    // 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
    _maxAdvHours   int64
    // 产品每日开始销售时间，start_time一定为当天时间
    _startTime   string
    // 产品每日结束销售时间,当end_time<start_time时，表示end_time为第二天，此时附加限制end_time<=06:00:00并且start_time>=12:00:00,表明可售时间从当天12点到次日的凌晨6点（扩展此信息主要为了描述尾房的rp）注意start_time一定是当天的时间。尾房18：00起可售
    _endTime   string
    // 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间,多间夜扣款。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
    _cancelPolicy   string
    // 1：开启（默认）2：关闭。如果没传值那么默认默认值为1
    _status   int64
    // RP的英文名称
    _englishName   string
    // 担保类型，只支持： 0  无担保  1  峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
    _guaranteeType   int64
    // 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
    _guaranteeStartTime   string
    // 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
    _memberLevel   string
    // 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪 O:钉钉商旅 A:集团内部商旅。如果只投放飞猪，改字段不用填写或者只填H；如果有多个用","分开。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
    _channel   string
    // 不推送则默认2人，如有低于2人的RP限制请推送该字段。
    _occupancy   int64
    // 系统商，一般不填写，使用须申请
    _vendor   string
    // 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0或者不填写代表否
    _firstStay   int64
    // 废弃。价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
    _agreement   int64
    // 在添加rateplan时，同时新增早餐日历。date：说明这条记录的早餐政策breakfast_count：这一天早餐的数量。>=-1,<=99。如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
    _breakfastCal   string
    // 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
    _cancelPolicyCal   string
    // 在新增rateplan的同时，新增担保日历。date：担保日历的某一天。guarantee:担保政策。其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm"。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
    _guaranteeCal   string
    // 不推荐使用，使用改规则
    _cancelBeforeHour   string
    // 不推荐使用，使用改规则
    _cancelBeforeDay   int64
    // 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
    _effectiveTime   string
    // 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
    _deadlineTime   string
    // rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
    _rpType   string
    // 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
    _hourage   string
    // 最早可选入住时间，小时房特有字段。格式为HH:mm
    _canCheckinEnd   string
    // 最晚可选入住时间，小时房特有字段。格式为HH:mm
    _canCheckinStart   string
    // 是否学生价，0：否；1：是。
    _isStudent   int64
    // 酒店id
    _hid   int64
    // 房型id
    _rid   int64
    // 外部房型id
    _outRid   string
    // 外部酒店id
    _outHid   string
    // super rp标记，1是；0否
    _superRpFlag   int64
    // base rp标记，1是；0否
    _baseRpFlag   int64
    // 2 VCC担保 1 PCI担保 0 支付宝担保(信用住产品担保方式只能为支付宝担保)
    _guaranteeMode   int64
    // 父rpid,使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
    _parentRpCode   string
    // 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
    _parentRpid   int64
    // 操作rateplan的来源
    _operator   string
    // 新增RP时的 打标和去标 需求,
    _tagJson   string
    // 来源
    _source   int64
    // 保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
    _allotmentReleaseTime   string
    // 普通保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
    _commonAllotReleaseTime   string
    // 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
    _resourceType   string
    // 是否底价加价，1是底价加价,0 非底价加价rp
    _bottomPriceFlag   int64
    // 最晚可选离店时间，小时房特有字段。格式为HH:mm
    _canCheckoutEnd   string
    // 会员价支持标识,1表示支持会员价规则
    _memDiscFlag   int64
    // 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)
    _memberDiscountCal   string
    // RP入住人限制信息。JSON格式
    _guestLimit   string
    // RP参与的活动，3为尾房,4超级房券
    _activityType   string
    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    _onlineBookingBindingInfo   string
}

// 初始化TaobaoXhotelRateplanAddRequest对象
func NewTaobaoXhotelRateplanAddRequest() *TaobaoXhotelRateplanAddRequest{
    return &TaobaoXhotelRateplanAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoXhotelRateplanAddRequest) SetRateplanCode(_rateplanCode string) error {
    r._rateplanCode = _rateplanCode
    r.Set("rateplan_code", _rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateplanAddRequest) GetRateplanCode() string {
    return r._rateplanCode
}
// Name Setter
// 在淘宝搜索页面展示的房价名称。请注意名称里不要维护早餐信息，如果想设置早餐信息，请设置breakfast_count字段即可
func (r *TaobaoXhotelRateplanAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoXhotelRateplanAddRequest) GetName() string {
    return r._name
}
// PaymentType Setter
// 支付类型，只支持：1：预付5：现付6: 信用住7:预付在线预约8:信用住在线预约。其中5,6,7,8四种类型需要申请权限
func (r *TaobaoXhotelRateplanAddRequest) SetPaymentType(_paymentType int64) error {
    r._paymentType = _paymentType
    r.Set("payment_type", _paymentType)
    return nil
}

// PaymentType Getter
func (r TaobaoXhotelRateplanAddRequest) GetPaymentType() int64 {
    return r._paymentType
}
// BreakfastCount Setter
// -1：状态早餐,有具体几人价有关系，几人价是几份早餐;0：不含早1：含单早2：含双早N：含N早（-1-99可选）
func (r *TaobaoXhotelRateplanAddRequest) SetBreakfastCount(_breakfastCount int64) error {
    r._breakfastCount = _breakfastCount
    r.Set("breakfast_count", _breakfastCount)
    return nil
}

// BreakfastCount Getter
func (r TaobaoXhotelRateplanAddRequest) GetBreakfastCount() int64 {
    return r._breakfastCount
}
// FeeServicePercent Setter
// 不推荐使用
func (r *TaobaoXhotelRateplanAddRequest) SetFeeServicePercent(_feeServicePercent int64) error {
    r._feeServicePercent = _feeServicePercent
    r.Set("fee_service_percent", _feeServicePercent)
    return nil
}

// FeeServicePercent Getter
func (r TaobaoXhotelRateplanAddRequest) GetFeeServicePercent() int64 {
    return r._feeServicePercent
}
// MinDays Setter
// 最小入住天数（1-90）。默认1,小时房RP请设置为1
func (r *TaobaoXhotelRateplanAddRequest) SetMinDays(_minDays int64) error {
    r._minDays = _minDays
    r.Set("min_days", _minDays)
    return nil
}

// MinDays Getter
func (r TaobaoXhotelRateplanAddRequest) GetMinDays() int64 {
    return r._minDays
}
// MaxDays Setter
// 最大入住天数（1-90）。默认90 信用住不超过9天,小时房RP请设置为1
func (r *TaobaoXhotelRateplanAddRequest) SetMaxDays(_maxDays int64) error {
    r._maxDays = _maxDays
    r.Set("max_days", _maxDays)
    return nil
}

// MaxDays Getter
func (r TaobaoXhotelRateplanAddRequest) GetMaxDays() int64 {
    return r._maxDays
}
// MinAmount Setter
// 首日入住房间数（1-99）。默认1。【废弃】
func (r *TaobaoXhotelRateplanAddRequest) SetMinAmount(_minAmount int64) error {
    r._minAmount = _minAmount
    r.Set("min_amount", _minAmount)
    return nil
}

// MinAmount Getter
func (r TaobaoXhotelRateplanAddRequest) GetMinAmount() int64 {
    return r._minAmount
}
// MinAdvHours Setter
// 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
func (r *TaobaoXhotelRateplanAddRequest) SetMinAdvHours(_minAdvHours int64) error {
    r._minAdvHours = _minAdvHours
    r.Set("min_adv_hours", _minAdvHours)
    return nil
}

// MinAdvHours Getter
func (r TaobaoXhotelRateplanAddRequest) GetMinAdvHours() int64 {
    return r._minAdvHours
}
// MaxAdvHours Setter
// 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
func (r *TaobaoXhotelRateplanAddRequest) SetMaxAdvHours(_maxAdvHours int64) error {
    r._maxAdvHours = _maxAdvHours
    r.Set("max_adv_hours", _maxAdvHours)
    return nil
}

// MaxAdvHours Getter
func (r TaobaoXhotelRateplanAddRequest) GetMaxAdvHours() int64 {
    return r._maxAdvHours
}
// StartTime Setter
// 产品每日开始销售时间，start_time一定为当天时间
func (r *TaobaoXhotelRateplanAddRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 产品每日结束销售时间,当end_time<start_time时，表示end_time为第二天，此时附加限制end_time<=06:00:00并且start_time>=12:00:00,表明可售时间从当天12点到次日的凌晨6点（扩展此信息主要为了描述尾房的rp）注意start_time一定是当天的时间。尾房18：00起可售
func (r *TaobaoXhotelRateplanAddRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetEndTime() string {
    return r._endTime
}
// CancelPolicy Setter
// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间,多间夜扣款。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
func (r *TaobaoXhotelRateplanAddRequest) SetCancelPolicy(_cancelPolicy string) error {
    r._cancelPolicy = _cancelPolicy
    r.Set("cancel_policy", _cancelPolicy)
    return nil
}

// CancelPolicy Getter
func (r TaobaoXhotelRateplanAddRequest) GetCancelPolicy() string {
    return r._cancelPolicy
}
// Status Setter
// 1：开启（默认）2：关闭。如果没传值那么默认默认值为1
func (r *TaobaoXhotelRateplanAddRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelRateplanAddRequest) GetStatus() int64 {
    return r._status
}
// EnglishName Setter
// RP的英文名称
func (r *TaobaoXhotelRateplanAddRequest) SetEnglishName(_englishName string) error {
    r._englishName = _englishName
    r.Set("english_name", _englishName)
    return nil
}

// EnglishName Getter
func (r TaobaoXhotelRateplanAddRequest) GetEnglishName() string {
    return r._englishName
}
// GuaranteeType Setter
// 担保类型，只支持： 0  无担保  1  峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
func (r *TaobaoXhotelRateplanAddRequest) SetGuaranteeType(_guaranteeType int64) error {
    r._guaranteeType = _guaranteeType
    r.Set("guarantee_type", _guaranteeType)
    return nil
}

// GuaranteeType Getter
func (r TaobaoXhotelRateplanAddRequest) GetGuaranteeType() int64 {
    return r._guaranteeType
}
// GuaranteeStartTime Setter
// 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
func (r *TaobaoXhotelRateplanAddRequest) SetGuaranteeStartTime(_guaranteeStartTime string) error {
    r._guaranteeStartTime = _guaranteeStartTime
    r.Set("guarantee_start_time", _guaranteeStartTime)
    return nil
}

// GuaranteeStartTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetGuaranteeStartTime() string {
    return r._guaranteeStartTime
}
// MemberLevel Setter
// 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
func (r *TaobaoXhotelRateplanAddRequest) SetMemberLevel(_memberLevel string) error {
    r._memberLevel = _memberLevel
    r.Set("member_level", _memberLevel)
    return nil
}

// MemberLevel Getter
func (r TaobaoXhotelRateplanAddRequest) GetMemberLevel() string {
    return r._memberLevel
}
// Channel Setter
// 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪 O:钉钉商旅 A:集团内部商旅。如果只投放飞猪，改字段不用填写或者只填H；如果有多个用","分开。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
func (r *TaobaoXhotelRateplanAddRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoXhotelRateplanAddRequest) GetChannel() string {
    return r._channel
}
// Occupancy Setter
// 不推送则默认2人，如有低于2人的RP限制请推送该字段。
func (r *TaobaoXhotelRateplanAddRequest) SetOccupancy(_occupancy int64) error {
    r._occupancy = _occupancy
    r.Set("occupancy", _occupancy)
    return nil
}

// Occupancy Getter
func (r TaobaoXhotelRateplanAddRequest) GetOccupancy() int64 {
    return r._occupancy
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRateplanAddRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateplanAddRequest) GetVendor() string {
    return r._vendor
}
// FirstStay Setter
// 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0或者不填写代表否
func (r *TaobaoXhotelRateplanAddRequest) SetFirstStay(_firstStay int64) error {
    r._firstStay = _firstStay
    r.Set("first_stay", _firstStay)
    return nil
}

// FirstStay Getter
func (r TaobaoXhotelRateplanAddRequest) GetFirstStay() int64 {
    return r._firstStay
}
// Agreement Setter
// 废弃。价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
func (r *TaobaoXhotelRateplanAddRequest) SetAgreement(_agreement int64) error {
    r._agreement = _agreement
    r.Set("agreement", _agreement)
    return nil
}

// Agreement Getter
func (r TaobaoXhotelRateplanAddRequest) GetAgreement() int64 {
    return r._agreement
}
// BreakfastCal Setter
// 在添加rateplan时，同时新增早餐日历。date：说明这条记录的早餐政策breakfast_count：这一天早餐的数量。>=-1,<=99。如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelRateplanAddRequest) SetBreakfastCal(_breakfastCal string) error {
    r._breakfastCal = _breakfastCal
    r.Set("breakfast_cal", _breakfastCal)
    return nil
}

// BreakfastCal Getter
func (r TaobaoXhotelRateplanAddRequest) GetBreakfastCal() string {
    return r._breakfastCal
}
// CancelPolicyCal Setter
// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
func (r *TaobaoXhotelRateplanAddRequest) SetCancelPolicyCal(_cancelPolicyCal string) error {
    r._cancelPolicyCal = _cancelPolicyCal
    r.Set("cancel_policy_cal", _cancelPolicyCal)
    return nil
}

// CancelPolicyCal Getter
func (r TaobaoXhotelRateplanAddRequest) GetCancelPolicyCal() string {
    return r._cancelPolicyCal
}
// GuaranteeCal Setter
// 在新增rateplan的同时，新增担保日历。date：担保日历的某一天。guarantee:担保政策。其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm"。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelRateplanAddRequest) SetGuaranteeCal(_guaranteeCal string) error {
    r._guaranteeCal = _guaranteeCal
    r.Set("guarantee_cal", _guaranteeCal)
    return nil
}

// GuaranteeCal Getter
func (r TaobaoXhotelRateplanAddRequest) GetGuaranteeCal() string {
    return r._guaranteeCal
}
// CancelBeforeHour Setter
// 不推荐使用，使用改规则
func (r *TaobaoXhotelRateplanAddRequest) SetCancelBeforeHour(_cancelBeforeHour string) error {
    r._cancelBeforeHour = _cancelBeforeHour
    r.Set("cancel_before_hour", _cancelBeforeHour)
    return nil
}

// CancelBeforeHour Getter
func (r TaobaoXhotelRateplanAddRequest) GetCancelBeforeHour() string {
    return r._cancelBeforeHour
}
// CancelBeforeDay Setter
// 不推荐使用，使用改规则
func (r *TaobaoXhotelRateplanAddRequest) SetCancelBeforeDay(_cancelBeforeDay int64) error {
    r._cancelBeforeDay = _cancelBeforeDay
    r.Set("cancel_before_day", _cancelBeforeDay)
    return nil
}

// CancelBeforeDay Getter
func (r TaobaoXhotelRateplanAddRequest) GetCancelBeforeDay() int64 {
    return r._cancelBeforeDay
}
// EffectiveTime Setter
// 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
func (r *TaobaoXhotelRateplanAddRequest) SetEffectiveTime(_effectiveTime string) error {
    r._effectiveTime = _effectiveTime
    r.Set("effective_time", _effectiveTime)
    return nil
}

// EffectiveTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetEffectiveTime() string {
    return r._effectiveTime
}
// DeadlineTime Setter
// 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
func (r *TaobaoXhotelRateplanAddRequest) SetDeadlineTime(_deadlineTime string) error {
    r._deadlineTime = _deadlineTime
    r.Set("deadline_time", _deadlineTime)
    return nil
}

// DeadlineTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetDeadlineTime() string {
    return r._deadlineTime
}
// RpType Setter
// rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
func (r *TaobaoXhotelRateplanAddRequest) SetRpType(_rpType string) error {
    r._rpType = _rpType
    r.Set("rp_type", _rpType)
    return nil
}

// RpType Getter
func (r TaobaoXhotelRateplanAddRequest) GetRpType() string {
    return r._rpType
}
// Hourage Setter
// 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
func (r *TaobaoXhotelRateplanAddRequest) SetHourage(_hourage string) error {
    r._hourage = _hourage
    r.Set("hourage", _hourage)
    return nil
}

// Hourage Getter
func (r TaobaoXhotelRateplanAddRequest) GetHourage() string {
    return r._hourage
}
// CanCheckinEnd Setter
// 最早可选入住时间，小时房特有字段。格式为HH:mm
func (r *TaobaoXhotelRateplanAddRequest) SetCanCheckinEnd(_canCheckinEnd string) error {
    r._canCheckinEnd = _canCheckinEnd
    r.Set("can_checkin_end", _canCheckinEnd)
    return nil
}

// CanCheckinEnd Getter
func (r TaobaoXhotelRateplanAddRequest) GetCanCheckinEnd() string {
    return r._canCheckinEnd
}
// CanCheckinStart Setter
// 最晚可选入住时间，小时房特有字段。格式为HH:mm
func (r *TaobaoXhotelRateplanAddRequest) SetCanCheckinStart(_canCheckinStart string) error {
    r._canCheckinStart = _canCheckinStart
    r.Set("can_checkin_start", _canCheckinStart)
    return nil
}

// CanCheckinStart Getter
func (r TaobaoXhotelRateplanAddRequest) GetCanCheckinStart() string {
    return r._canCheckinStart
}
// IsStudent Setter
// 是否学生价，0：否；1：是。
func (r *TaobaoXhotelRateplanAddRequest) SetIsStudent(_isStudent int64) error {
    r._isStudent = _isStudent
    r.Set("is_student", _isStudent)
    return nil
}

// IsStudent Getter
func (r TaobaoXhotelRateplanAddRequest) GetIsStudent() int64 {
    return r._isStudent
}
// Hid Setter
// 酒店id
func (r *TaobaoXhotelRateplanAddRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelRateplanAddRequest) GetHid() int64 {
    return r._hid
}
// Rid Setter
// 房型id
func (r *TaobaoXhotelRateplanAddRequest) SetRid(_rid int64) error {
    r._rid = _rid
    r.Set("rid", _rid)
    return nil
}

// Rid Getter
func (r TaobaoXhotelRateplanAddRequest) GetRid() int64 {
    return r._rid
}
// OutRid Setter
// 外部房型id
func (r *TaobaoXhotelRateplanAddRequest) SetOutRid(_outRid string) error {
    r._outRid = _outRid
    r.Set("out_rid", _outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRateplanAddRequest) GetOutRid() string {
    return r._outRid
}
// OutHid Setter
// 外部酒店id
func (r *TaobaoXhotelRateplanAddRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelRateplanAddRequest) GetOutHid() string {
    return r._outHid
}
// SuperRpFlag Setter
// super rp标记，1是；0否
func (r *TaobaoXhotelRateplanAddRequest) SetSuperRpFlag(_superRpFlag int64) error {
    r._superRpFlag = _superRpFlag
    r.Set("super_rp_flag", _superRpFlag)
    return nil
}

// SuperRpFlag Getter
func (r TaobaoXhotelRateplanAddRequest) GetSuperRpFlag() int64 {
    return r._superRpFlag
}
// BaseRpFlag Setter
// base rp标记，1是；0否
func (r *TaobaoXhotelRateplanAddRequest) SetBaseRpFlag(_baseRpFlag int64) error {
    r._baseRpFlag = _baseRpFlag
    r.Set("base_rp_flag", _baseRpFlag)
    return nil
}

// BaseRpFlag Getter
func (r TaobaoXhotelRateplanAddRequest) GetBaseRpFlag() int64 {
    return r._baseRpFlag
}
// GuaranteeMode Setter
// 2 VCC担保 1 PCI担保 0 支付宝担保(信用住产品担保方式只能为支付宝担保)
func (r *TaobaoXhotelRateplanAddRequest) SetGuaranteeMode(_guaranteeMode int64) error {
    r._guaranteeMode = _guaranteeMode
    r.Set("guarantee_mode", _guaranteeMode)
    return nil
}

// GuaranteeMode Getter
func (r TaobaoXhotelRateplanAddRequest) GetGuaranteeMode() int64 {
    return r._guaranteeMode
}
// ParentRpCode Setter
// 父rpid,使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
func (r *TaobaoXhotelRateplanAddRequest) SetParentRpCode(_parentRpCode string) error {
    r._parentRpCode = _parentRpCode
    r.Set("parent_rp_code", _parentRpCode)
    return nil
}

// ParentRpCode Getter
func (r TaobaoXhotelRateplanAddRequest) GetParentRpCode() string {
    return r._parentRpCode
}
// ParentRpid Setter
// 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
func (r *TaobaoXhotelRateplanAddRequest) SetParentRpid(_parentRpid int64) error {
    r._parentRpid = _parentRpid
    r.Set("parent_rpid", _parentRpid)
    return nil
}

// ParentRpid Getter
func (r TaobaoXhotelRateplanAddRequest) GetParentRpid() int64 {
    return r._parentRpid
}
// Operator Setter
// 操作rateplan的来源
func (r *TaobaoXhotelRateplanAddRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelRateplanAddRequest) GetOperator() string {
    return r._operator
}
// TagJson Setter
// 新增RP时的 打标和去标 需求,
func (r *TaobaoXhotelRateplanAddRequest) SetTagJson(_tagJson string) error {
    r._tagJson = _tagJson
    r.Set("tag_json", _tagJson)
    return nil
}

// TagJson Getter
func (r TaobaoXhotelRateplanAddRequest) GetTagJson() string {
    return r._tagJson
}
// Source Setter
// 来源
func (r *TaobaoXhotelRateplanAddRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoXhotelRateplanAddRequest) GetSource() int64 {
    return r._source
}
// AllotmentReleaseTime Setter
// 保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
func (r *TaobaoXhotelRateplanAddRequest) SetAllotmentReleaseTime(_allotmentReleaseTime string) error {
    r._allotmentReleaseTime = _allotmentReleaseTime
    r.Set("allotment_release_time", _allotmentReleaseTime)
    return nil
}

// AllotmentReleaseTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetAllotmentReleaseTime() string {
    return r._allotmentReleaseTime
}
// CommonAllotReleaseTime Setter
// 普通保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
func (r *TaobaoXhotelRateplanAddRequest) SetCommonAllotReleaseTime(_commonAllotReleaseTime string) error {
    r._commonAllotReleaseTime = _commonAllotReleaseTime
    r.Set("common_allot_release_time", _commonAllotReleaseTime)
    return nil
}

// CommonAllotReleaseTime Getter
func (r TaobaoXhotelRateplanAddRequest) GetCommonAllotReleaseTime() string {
    return r._commonAllotReleaseTime
}
// ResourceType Setter
// 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
func (r *TaobaoXhotelRateplanAddRequest) SetResourceType(_resourceType string) error {
    r._resourceType = _resourceType
    r.Set("resource_type", _resourceType)
    return nil
}

// ResourceType Getter
func (r TaobaoXhotelRateplanAddRequest) GetResourceType() string {
    return r._resourceType
}
// BottomPriceFlag Setter
// 是否底价加价，1是底价加价,0 非底价加价rp
func (r *TaobaoXhotelRateplanAddRequest) SetBottomPriceFlag(_bottomPriceFlag int64) error {
    r._bottomPriceFlag = _bottomPriceFlag
    r.Set("bottom_price_flag", _bottomPriceFlag)
    return nil
}

// BottomPriceFlag Getter
func (r TaobaoXhotelRateplanAddRequest) GetBottomPriceFlag() int64 {
    return r._bottomPriceFlag
}
// CanCheckoutEnd Setter
// 最晚可选离店时间，小时房特有字段。格式为HH:mm
func (r *TaobaoXhotelRateplanAddRequest) SetCanCheckoutEnd(_canCheckoutEnd string) error {
    r._canCheckoutEnd = _canCheckoutEnd
    r.Set("can_checkout_end", _canCheckoutEnd)
    return nil
}

// CanCheckoutEnd Getter
func (r TaobaoXhotelRateplanAddRequest) GetCanCheckoutEnd() string {
    return r._canCheckoutEnd
}
// MemDiscFlag Setter
// 会员价支持标识,1表示支持会员价规则
func (r *TaobaoXhotelRateplanAddRequest) SetMemDiscFlag(_memDiscFlag int64) error {
    r._memDiscFlag = _memDiscFlag
    r.Set("mem_disc_flag", _memDiscFlag)
    return nil
}

// MemDiscFlag Getter
func (r TaobaoXhotelRateplanAddRequest) GetMemDiscFlag() int64 {
    return r._memDiscFlag
}
// MemberDiscountCal Setter
// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)
func (r *TaobaoXhotelRateplanAddRequest) SetMemberDiscountCal(_memberDiscountCal string) error {
    r._memberDiscountCal = _memberDiscountCal
    r.Set("member_discount_cal", _memberDiscountCal)
    return nil
}

// MemberDiscountCal Getter
func (r TaobaoXhotelRateplanAddRequest) GetMemberDiscountCal() string {
    return r._memberDiscountCal
}
// GuestLimit Setter
// RP入住人限制信息。JSON格式
func (r *TaobaoXhotelRateplanAddRequest) SetGuestLimit(_guestLimit string) error {
    r._guestLimit = _guestLimit
    r.Set("guest_limit", _guestLimit)
    return nil
}

// GuestLimit Getter
func (r TaobaoXhotelRateplanAddRequest) GetGuestLimit() string {
    return r._guestLimit
}
// ActivityType Setter
// RP参与的活动，3为尾房,4超级房券
func (r *TaobaoXhotelRateplanAddRequest) SetActivityType(_activityType string) error {
    r._activityType = _activityType
    r.Set("activity_type", _activityType)
    return nil
}

// ActivityType Getter
func (r TaobaoXhotelRateplanAddRequest) GetActivityType() string {
    return r._activityType
}
// OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoXhotelRateplanAddRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
    r._onlineBookingBindingInfo = _onlineBookingBindingInfo
    r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
    return nil
}

// OnlineBookingBindingInfo Getter
func (r TaobaoXhotelRateplanAddRequest) GetOnlineBookingBindingInfo() string {
    return r._onlineBookingBindingInfo
}
