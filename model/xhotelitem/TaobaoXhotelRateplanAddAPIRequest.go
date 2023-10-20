package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelrateplanaddAPIRequest 酒店产品库rateplan添加 API请求
// taobao.xhotel.rateplan.add
//
// 酒店产品库rateplan
type TaobaoxhotelrateplanaddAPIRequest struct {
	model.Params
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 在淘宝搜索页面展示的房价名称。请注意名称里不要维护早餐信息，如果想设置早餐信息，请设置breakfast_count字段即可
	_name string
	// RP的英文名称
	_englishName string
	// 不推荐使用
	_extendFee string
	// 产品每日开始销售时间，start_time一定为当天时间
	_startTime string
	// 产品每日结束销售时间,当end_time<start_time时，表示end_time为第二天，此时附加限制end_time<=06:00:00并且start_time>=12:00:00,表明可售时间从当天12点到次日的凌晨6点（扩展此信息主要为了描述尾房的rp）注意start_time一定是当天的时间。尾房18：00起可售
	_endTime string
	// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间,多间夜扣款。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
	_cancelPolicy string
	// 个性化定制扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
	_extend string
	// 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
	_guaranteeStartTime string
	// 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
	_memberLevel string
	// 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪全渠道（选择H，可实现飞猪、高德、支付宝、手淘均可售卖） O:钉钉商旅 。如果有多个用","分开，比如H,O。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
	_channel string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 不推荐使用，使用改规则
	_cancelBeforeHour string
	// 在添加rateplan时，同时新增早餐日历。date：说明这条记录的早餐政策breakfast_count：这一天早餐的数量。>=-1,<=99。如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_breakfastCal string
	// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
	_cancelPolicyCal string
	// 在新增rateplan的同时，新增担保日历。date：担保日历的某一天。guarantee:担保政策。其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm"。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_guaranteeCal string
	// 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
	_effectiveTime string
	// 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
	_deadlineTime string
	// rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
	_rpType string
	// 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
	_hourage string
	// 最早可选入住时间，小时房特有字段。格式为HH:mm
	_canCheckinEnd string
	// 最晚可选入住时间，小时房特有字段。格式为HH:mm
	_canCheckinStart string
	// 餐食描述
	_dinningDesc string
	// 外部酒店id
	_outHid string
	// 外部房型id
	_outRid string
	// 父rpid,使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
	_parentRpCode string
	// 操作rateplan的来源
	_operator string
	// 新增RP时的 打标和去标 需求,
	_tagJson string
	// 保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
	_allotmentReleaseTime string
	// 普通保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
	_commonAllotReleaseTime string
	// 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
	_resourceType string
	// 最晚可选离店时间，小时房特有字段。格式为HH:mm
	_canCheckoutEnd string
	// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)
	_memberDiscountCal string
	// RP入住人限制信息。JSON格式
	_guestLimit string
	// RP参与的活动，3为尾房,4超级房券,8直连免房
	_activityType string
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
	// rp的权益信息, 调用该字段请优先联系对接业务同学。type枚举: eeo,meo, value取值:1. 额外积分 2. 优惠价格 3. 套餐权益 4.行政礼遇。
	_rights string
	// 商业化充值类型 seller充值到卖家 hotel充值到门店
	_freeRoomChargeDstRole string
	// 支付类型，只支持：1：预付5：现付6: 信用住7:预付在线预约8:信用住在线预约。其中5,6,7,8四种类型需要申请权限
	_paymentType int64
	// -1：状态早餐,有具体几人价有关系，几人价是几份早餐;0：不含早1：含单早2：含双早N：含N早（-1-99可选）
	_breakfastCount int64
	// 废弃
	_feeBreakfastCount int64
	// 不推荐使用
	_feeBreakfastAmount int64
	// 不推荐使用
	_feeGovTaxAmount int64
	// 不推荐使用
	_feeGovTaxPercent int64
	// 不推荐使用
	_feeServiceAmount int64
	// 不推荐使用
	_feeServicePercent int64
	// 最小入住天数（1-90）。默认1,小时房RP请设置为1
	_minDays int64
	// 最大入住天数（1-90）。默认90 信用住不超过9天,小时房RP请设置为1,特殊商家支持180天
	_maxDays int64
	// 首日入住房间数（1-99）。默认1。【废弃】
	_minAmount int64
	// 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
	_minAdvHours int64
	// 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
	_maxAdvHours int64
	// 1：开启（默认）2：关闭。如果没传值那么默认默认值为1
	_status int64
	// 担保类型，只支持： 0  无担保  1  峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
	_guaranteeType int64
	// 不推送则默认2人，如有低于2人的RP限制请推送该字段。
	_occupancy int64
	// 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0或者不填写代表否
	_firstStay int64
	// 废弃。价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
	_agreement int64
	// 不推荐使用，使用改规则
	_cancelBeforeDay int64
	// 儿童最大年龄(0-18)
	_maxChildAge int64
	// 儿童最小年龄(0-18)
	_minChildAge int64
	// 婴儿最大年龄(0-18)
	_maxInfantAge int64
	// 婴儿最小年龄(0-18)
	_minInfantAge int64
	// 是否学生价，0：否；1：是。
	_isStudent int64
	// 酒店id
	_hid int64
	// 房型id
	_rid int64
	// super rp标记，1是；0否
	_superRpFlag int64
	// base rp标记，1是；0否
	_baseRpFlag int64
	// 2 VCC担保 1 PCI担保 0 支付宝担保(信用住产品担保方式只能为支付宝担保)
	_guaranteeMode int64
	// 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
	_parentRpid int64
	// 来源
	_source int64
	// 是否底价加价，1是底价加价,0 非底价加价rp
	_bottomPriceFlag int64
	// 会员价支持标识,1表示支持会员价规则
	_memDiscFlag int64
}

// NewTaobaoxhotelrateplanaddRequest 初始化TaobaoxhotelrateplanaddAPIRequest对象
func NewTaobaoxhotelrateplanaddRequest() *TaobaoxhotelrateplanaddAPIRequest {
	return &TaobaoxhotelrateplanaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelrateplanaddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rateplan.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelrateplanaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelrateplanaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRateplanCode is RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoxhotelrateplanaddAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetName is Name Setter
// 在淘宝搜索页面展示的房价名称。请注意名称里不要维护早餐信息，如果想设置早餐信息，请设置breakfast_count字段即可
func (r *TaobaoxhotelrateplanaddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetName() string {
	return r._name
}

// SetEnglishName is EnglishName Setter
// RP的英文名称
func (r *TaobaoxhotelrateplanaddAPIRequest) SetEnglishName(_englishName string) error {
	r._englishName = _englishName
	r.Set("english_name", _englishName)
	return nil
}

// GetEnglishName EnglishName Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetEnglishName() string {
	return r._englishName
}

// SetExtendFee is ExtendFee Setter
// 不推荐使用
func (r *TaobaoxhotelrateplanaddAPIRequest) SetExtendFee(_extendFee string) error {
	r._extendFee = _extendFee
	r.Set("extend_fee", _extendFee)
	return nil
}

// GetExtendFee ExtendFee Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetExtendFee() string {
	return r._extendFee
}

// SetStartTime is StartTime Setter
// 产品每日开始销售时间，start_time一定为当天时间
func (r *TaobaoxhotelrateplanaddAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 产品每日结束销售时间,当end_time&lt;start_time时，表示end_time为第二天，此时附加限制end_time&lt;=06:00:00并且start_time&gt;=12:00:00,表明可售时间从当天12点到次日的凌晨6点（扩展此信息主要为了描述尾房的rp）注意start_time一定是当天的时间。尾房18：00起可售
func (r *TaobaoxhotelrateplanaddAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCancelPolicy is CancelPolicy Setter
// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100&gt;Y&gt;=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100&gt;Y&gt;=0（不超过10条）。1.表示任意退{&#34;cancelPolicyType&#34;:1};2.表示不能退{&#34;cancelPolicyType&#34;:2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{&#34;cancelPolicyType&#34;:4,&#34;policyInfo&#34;:{&#34;48&#34;:10,&#34;24&#34;:20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{&#34;cancelPolicyType&#34;:5,&#34;policyInfo&#34;:{&#34;timeBefore&#34;:6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{&#34;cancelPolicyType&#34;:6,&#34;policyInfo&#34;:{&#34;14&#34;:1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间,多间夜扣款。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCancelPolicy(_cancelPolicy string) error {
	r._cancelPolicy = _cancelPolicy
	r.Set("cancel_policy", _cancelPolicy)
	return nil
}

// GetCancelPolicy CancelPolicy Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCancelPolicy() string {
	return r._cancelPolicy
}

// SetExtend is Extend Setter
// 个性化定制扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
func (r *TaobaoxhotelrateplanaddAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetExtend() string {
	return r._extend
}

// SetGuaranteeStartTime is GuaranteeStartTime Setter
// 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
func (r *TaobaoxhotelrateplanaddAPIRequest) SetGuaranteeStartTime(_guaranteeStartTime string) error {
	r._guaranteeStartTime = _guaranteeStartTime
	r.Set("guarantee_start_time", _guaranteeStartTime)
	return nil
}

// GetGuaranteeStartTime GuaranteeStartTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetGuaranteeStartTime() string {
	return r._guaranteeStartTime
}

// SetMemberLevel is MemberLevel Setter
// 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMemberLevel(_memberLevel string) error {
	r._memberLevel = _memberLevel
	r.Set("member_level", _memberLevel)
	return nil
}

// GetMemberLevel MemberLevel Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMemberLevel() string {
	return r._memberLevel
}

// SetChannel is Channel Setter
// 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪全渠道（选择H，可实现飞猪、高德、支付宝、手淘均可售卖） O:钉钉商旅 。如果有多个用&#34;,&#34;分开，比如H,O。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetChannel() string {
	return r._channel
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoxhotelrateplanaddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetVendor() string {
	return r._vendor
}

// SetCancelBeforeHour is CancelBeforeHour Setter
// 不推荐使用，使用改规则
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCancelBeforeHour(_cancelBeforeHour string) error {
	r._cancelBeforeHour = _cancelBeforeHour
	r.Set("cancel_before_hour", _cancelBeforeHour)
	return nil
}

// GetCancelBeforeHour CancelBeforeHour Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCancelBeforeHour() string {
	return r._cancelBeforeHour
}

// SetBreakfastCal is BreakfastCal Setter
// 在添加rateplan时，同时新增早餐日历。date：说明这条记录的早餐政策breakfast_count：这一天早餐的数量。&gt;=-1,&lt;=99。如果date为空，那么会去读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetBreakfastCal(_breakfastCal string) error {
	r._breakfastCal = _breakfastCal
	r.Set("breakfast_cal", _breakfastCal)
	return nil
}

// GetBreakfastCal BreakfastCal Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetBreakfastCal() string {
	return r._breakfastCal
}

// SetCancelPolicyCal is CancelPolicyCal Setter
// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为&#34;yyyy-MM-dd&#34; cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCancelPolicyCal(_cancelPolicyCal string) error {
	r._cancelPolicyCal = _cancelPolicyCal
	r.Set("cancel_policy_cal", _cancelPolicyCal)
	return nil
}

// GetCancelPolicyCal CancelPolicyCal Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCancelPolicyCal() string {
	return r._cancelPolicyCal
}

// SetGuaranteeCal is GuaranteeCal Setter
// 在新增rateplan的同时，新增担保日历。date：担保日历的某一天。guarantee:担保政策。其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为&#34;HH:mm&#34;。如果date为空，那么会读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetGuaranteeCal(_guaranteeCal string) error {
	r._guaranteeCal = _guaranteeCal
	r.Set("guarantee_cal", _guaranteeCal)
	return nil
}

// GetGuaranteeCal GuaranteeCal Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetGuaranteeCal() string {
	return r._guaranteeCal
}

// SetEffectiveTime is EffectiveTime Setter
// 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
func (r *TaobaoxhotelrateplanaddAPIRequest) SetEffectiveTime(_effectiveTime string) error {
	r._effectiveTime = _effectiveTime
	r.Set("effective_time", _effectiveTime)
	return nil
}

// GetEffectiveTime EffectiveTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetEffectiveTime() string {
	return r._effectiveTime
}

// SetDeadlineTime is DeadlineTime Setter
// 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
func (r *TaobaoxhotelrateplanaddAPIRequest) SetDeadlineTime(_deadlineTime string) error {
	r._deadlineTime = _deadlineTime
	r.Set("deadline_time", _deadlineTime)
	return nil
}

// GetDeadlineTime DeadlineTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetDeadlineTime() string {
	return r._deadlineTime
}

// SetRpType is RpType Setter
// rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
func (r *TaobaoxhotelrateplanaddAPIRequest) SetRpType(_rpType string) error {
	r._rpType = _rpType
	r.Set("rp_type", _rpType)
	return nil
}

// GetRpType RpType Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetRpType() string {
	return r._rpType
}

// SetHourage is Hourage Setter
// 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
func (r *TaobaoxhotelrateplanaddAPIRequest) SetHourage(_hourage string) error {
	r._hourage = _hourage
	r.Set("hourage", _hourage)
	return nil
}

// GetHourage Hourage Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetHourage() string {
	return r._hourage
}

// SetCanCheckinEnd is CanCheckinEnd Setter
// 最早可选入住时间，小时房特有字段。格式为HH:mm
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCanCheckinEnd(_canCheckinEnd string) error {
	r._canCheckinEnd = _canCheckinEnd
	r.Set("can_checkin_end", _canCheckinEnd)
	return nil
}

// GetCanCheckinEnd CanCheckinEnd Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCanCheckinEnd() string {
	return r._canCheckinEnd
}

// SetCanCheckinStart is CanCheckinStart Setter
// 最晚可选入住时间，小时房特有字段。格式为HH:mm
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCanCheckinStart(_canCheckinStart string) error {
	r._canCheckinStart = _canCheckinStart
	r.Set("can_checkin_start", _canCheckinStart)
	return nil
}

// GetCanCheckinStart CanCheckinStart Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCanCheckinStart() string {
	return r._canCheckinStart
}

// SetDinningDesc is DinningDesc Setter
// 餐食描述
func (r *TaobaoxhotelrateplanaddAPIRequest) SetDinningDesc(_dinningDesc string) error {
	r._dinningDesc = _dinningDesc
	r.Set("dinning_desc", _dinningDesc)
	return nil
}

// GetDinningDesc DinningDesc Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetDinningDesc() string {
	return r._dinningDesc
}

// SetOutHid is OutHid Setter
// 外部酒店id
func (r *TaobaoxhotelrateplanaddAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetOutRid is OutRid Setter
// 外部房型id
func (r *TaobaoxhotelrateplanaddAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetParentRpCode is ParentRpCode Setter
// 父rpid,使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
func (r *TaobaoxhotelrateplanaddAPIRequest) SetParentRpCode(_parentRpCode string) error {
	r._parentRpCode = _parentRpCode
	r.Set("parent_rp_code", _parentRpCode)
	return nil
}

// GetParentRpCode ParentRpCode Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetParentRpCode() string {
	return r._parentRpCode
}

// SetOperator is Operator Setter
// 操作rateplan的来源
func (r *TaobaoxhotelrateplanaddAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetOperator() string {
	return r._operator
}

// SetTagJson is TagJson Setter
// 新增RP时的 打标和去标 需求,
func (r *TaobaoxhotelrateplanaddAPIRequest) SetTagJson(_tagJson string) error {
	r._tagJson = _tagJson
	r.Set("tag_json", _tagJson)
	return nil
}

// GetTagJson TagJson Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetTagJson() string {
	return r._tagJson
}

// SetAllotmentReleaseTime is AllotmentReleaseTime Setter
// 保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
func (r *TaobaoxhotelrateplanaddAPIRequest) SetAllotmentReleaseTime(_allotmentReleaseTime string) error {
	r._allotmentReleaseTime = _allotmentReleaseTime
	r.Set("allotment_release_time", _allotmentReleaseTime)
	return nil
}

// GetAllotmentReleaseTime AllotmentReleaseTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetAllotmentReleaseTime() string {
	return r._allotmentReleaseTime
}

// SetCommonAllotReleaseTime is CommonAllotReleaseTime Setter
// 普通保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCommonAllotReleaseTime(_commonAllotReleaseTime string) error {
	r._commonAllotReleaseTime = _commonAllotReleaseTime
	r.Set("common_allot_release_time", _commonAllotReleaseTime)
	return nil
}

// GetCommonAllotReleaseTime CommonAllotReleaseTime Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCommonAllotReleaseTime() string {
	return r._commonAllotReleaseTime
}

// SetResourceType is ResourceType Setter
// 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetResourceType(_resourceType string) error {
	r._resourceType = _resourceType
	r.Set("resource_type", _resourceType)
	return nil
}

// GetResourceType ResourceType Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetResourceType() string {
	return r._resourceType
}

// SetCanCheckoutEnd is CanCheckoutEnd Setter
// 最晚可选离店时间，小时房特有字段。格式为HH:mm
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCanCheckoutEnd(_canCheckoutEnd string) error {
	r._canCheckoutEnd = _canCheckoutEnd
	r.Set("can_checkout_end", _canCheckoutEnd)
	return nil
}

// GetCanCheckoutEnd CanCheckoutEnd Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCanCheckoutEnd() string {
	return r._canCheckoutEnd
}

// SetMemberDiscountCal is MemberDiscountCal Setter
// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMemberDiscountCal(_memberDiscountCal string) error {
	r._memberDiscountCal = _memberDiscountCal
	r.Set("member_discount_cal", _memberDiscountCal)
	return nil
}

// GetMemberDiscountCal MemberDiscountCal Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMemberDiscountCal() string {
	return r._memberDiscountCal
}

// SetGuestLimit is GuestLimit Setter
// RP入住人限制信息。JSON格式
func (r *TaobaoxhotelrateplanaddAPIRequest) SetGuestLimit(_guestLimit string) error {
	r._guestLimit = _guestLimit
	r.Set("guest_limit", _guestLimit)
	return nil
}

// GetGuestLimit GuestLimit Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetGuestLimit() string {
	return r._guestLimit
}

// SetActivityType is ActivityType Setter
// RP参与的活动，3为尾房,4超级房券,8直连免房
func (r *TaobaoxhotelrateplanaddAPIRequest) SetActivityType(_activityType string) error {
	r._activityType = _activityType
	r.Set("activity_type", _activityType)
	return nil
}

// GetActivityType ActivityType Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetActivityType() string {
	return r._activityType
}

// SetOnlineBookingBindingInfo is OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoxhotelrateplanaddAPIRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
	r._onlineBookingBindingInfo = _onlineBookingBindingInfo
	r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
	return nil
}

// GetOnlineBookingBindingInfo OnlineBookingBindingInfo Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetOnlineBookingBindingInfo() string {
	return r._onlineBookingBindingInfo
}

// SetRights is Rights Setter
// rp的权益信息, 调用该字段请优先联系对接业务同学。type枚举: eeo,meo, value取值:1. 额外积分 2. 优惠价格 3. 套餐权益 4.行政礼遇。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetRights(_rights string) error {
	r._rights = _rights
	r.Set("rights", _rights)
	return nil
}

// GetRights Rights Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetRights() string {
	return r._rights
}

// SetFreeRoomChargeDstRole is FreeRoomChargeDstRole Setter
// 商业化充值类型 seller充值到卖家 hotel充值到门店
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFreeRoomChargeDstRole(_freeRoomChargeDstRole string) error {
	r._freeRoomChargeDstRole = _freeRoomChargeDstRole
	r.Set("free_room_charge_dst_role", _freeRoomChargeDstRole)
	return nil
}

// GetFreeRoomChargeDstRole FreeRoomChargeDstRole Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFreeRoomChargeDstRole() string {
	return r._freeRoomChargeDstRole
}

// SetPaymentType is PaymentType Setter
// 支付类型，只支持：1：预付5：现付6: 信用住7:预付在线预约8:信用住在线预约。其中5,6,7,8四种类型需要申请权限
func (r *TaobaoxhotelrateplanaddAPIRequest) SetPaymentType(_paymentType int64) error {
	r._paymentType = _paymentType
	r.Set("payment_type", _paymentType)
	return nil
}

// GetPaymentType PaymentType Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetPaymentType() int64 {
	return r._paymentType
}

// SetBreakfastCount is BreakfastCount Setter
// -1：状态早餐,有具体几人价有关系，几人价是几份早餐;0：不含早1：含单早2：含双早N：含N早（-1-99可选）
func (r *TaobaoxhotelrateplanaddAPIRequest) SetBreakfastCount(_breakfastCount int64) error {
	r._breakfastCount = _breakfastCount
	r.Set("breakfast_count", _breakfastCount)
	return nil
}

// GetBreakfastCount BreakfastCount Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetBreakfastCount() int64 {
	return r._breakfastCount
}

// SetFeeBreakfastCount is FeeBreakfastCount Setter
// 废弃
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFeeBreakfastCount(_feeBreakfastCount int64) error {
	r._feeBreakfastCount = _feeBreakfastCount
	r.Set("fee_breakfast_count", _feeBreakfastCount)
	return nil
}

// GetFeeBreakfastCount FeeBreakfastCount Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFeeBreakfastCount() int64 {
	return r._feeBreakfastCount
}

// SetFeeBreakfastAmount is FeeBreakfastAmount Setter
// 不推荐使用
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFeeBreakfastAmount(_feeBreakfastAmount int64) error {
	r._feeBreakfastAmount = _feeBreakfastAmount
	r.Set("fee_breakfast_amount", _feeBreakfastAmount)
	return nil
}

// GetFeeBreakfastAmount FeeBreakfastAmount Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFeeBreakfastAmount() int64 {
	return r._feeBreakfastAmount
}

// SetFeeGovTaxAmount is FeeGovTaxAmount Setter
// 不推荐使用
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFeeGovTaxAmount(_feeGovTaxAmount int64) error {
	r._feeGovTaxAmount = _feeGovTaxAmount
	r.Set("fee_gov_tax_amount", _feeGovTaxAmount)
	return nil
}

// GetFeeGovTaxAmount FeeGovTaxAmount Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFeeGovTaxAmount() int64 {
	return r._feeGovTaxAmount
}

// SetFeeGovTaxPercent is FeeGovTaxPercent Setter
// 不推荐使用
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFeeGovTaxPercent(_feeGovTaxPercent int64) error {
	r._feeGovTaxPercent = _feeGovTaxPercent
	r.Set("fee_gov_tax_percent", _feeGovTaxPercent)
	return nil
}

// GetFeeGovTaxPercent FeeGovTaxPercent Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFeeGovTaxPercent() int64 {
	return r._feeGovTaxPercent
}

// SetFeeServiceAmount is FeeServiceAmount Setter
// 不推荐使用
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFeeServiceAmount(_feeServiceAmount int64) error {
	r._feeServiceAmount = _feeServiceAmount
	r.Set("fee_service_amount", _feeServiceAmount)
	return nil
}

// GetFeeServiceAmount FeeServiceAmount Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFeeServiceAmount() int64 {
	return r._feeServiceAmount
}

// SetFeeServicePercent is FeeServicePercent Setter
// 不推荐使用
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFeeServicePercent(_feeServicePercent int64) error {
	r._feeServicePercent = _feeServicePercent
	r.Set("fee_service_percent", _feeServicePercent)
	return nil
}

// GetFeeServicePercent FeeServicePercent Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFeeServicePercent() int64 {
	return r._feeServicePercent
}

// SetMinDays is MinDays Setter
// 最小入住天数（1-90）。默认1,小时房RP请设置为1
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMinDays(_minDays int64) error {
	r._minDays = _minDays
	r.Set("min_days", _minDays)
	return nil
}

// GetMinDays MinDays Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMinDays() int64 {
	return r._minDays
}

// SetMaxDays is MaxDays Setter
// 最大入住天数（1-90）。默认90 信用住不超过9天,小时房RP请设置为1,特殊商家支持180天
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMaxDays(_maxDays int64) error {
	r._maxDays = _maxDays
	r.Set("max_days", _maxDays)
	return nil
}

// GetMaxDays MaxDays Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMaxDays() int64 {
	return r._maxDays
}

// SetMinAmount is MinAmount Setter
// 首日入住房间数（1-99）。默认1。【废弃】
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMinAmount(_minAmount int64) error {
	r._minAmount = _minAmount
	r.Set("min_amount", _minAmount)
	return nil
}

// GetMinAmount MinAmount Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMinAmount() int64 {
	return r._minAmount
}

// SetMinAdvHours is MinAdvHours Setter
// 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMinAdvHours(_minAdvHours int64) error {
	r._minAdvHours = _minAdvHours
	r.Set("min_adv_hours", _minAdvHours)
	return nil
}

// GetMinAdvHours MinAdvHours Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMinAdvHours() int64 {
	return r._minAdvHours
}

// SetMaxAdvHours is MaxAdvHours Setter
// 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMaxAdvHours(_maxAdvHours int64) error {
	r._maxAdvHours = _maxAdvHours
	r.Set("max_adv_hours", _maxAdvHours)
	return nil
}

// GetMaxAdvHours MaxAdvHours Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMaxAdvHours() int64 {
	return r._maxAdvHours
}

// SetStatus is Status Setter
// 1：开启（默认）2：关闭。如果没传值那么默认默认值为1
func (r *TaobaoxhotelrateplanaddAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetStatus() int64 {
	return r._status
}

// SetGuaranteeType is GuaranteeType Setter
// 担保类型，只支持： 0  无担保  1  峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
func (r *TaobaoxhotelrateplanaddAPIRequest) SetGuaranteeType(_guaranteeType int64) error {
	r._guaranteeType = _guaranteeType
	r.Set("guarantee_type", _guaranteeType)
	return nil
}

// GetGuaranteeType GuaranteeType Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetGuaranteeType() int64 {
	return r._guaranteeType
}

// SetOccupancy is Occupancy Setter
// 不推送则默认2人，如有低于2人的RP限制请推送该字段。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetOccupancy(_occupancy int64) error {
	r._occupancy = _occupancy
	r.Set("occupancy", _occupancy)
	return nil
}

// GetOccupancy Occupancy Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetOccupancy() int64 {
	return r._occupancy
}

// SetFirstStay is FirstStay Setter
// 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0或者不填写代表否
func (r *TaobaoxhotelrateplanaddAPIRequest) SetFirstStay(_firstStay int64) error {
	r._firstStay = _firstStay
	r.Set("first_stay", _firstStay)
	return nil
}

// GetFirstStay FirstStay Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetFirstStay() int64 {
	return r._firstStay
}

// SetAgreement is Agreement Setter
// 废弃。价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
func (r *TaobaoxhotelrateplanaddAPIRequest) SetAgreement(_agreement int64) error {
	r._agreement = _agreement
	r.Set("agreement", _agreement)
	return nil
}

// GetAgreement Agreement Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetAgreement() int64 {
	return r._agreement
}

// SetCancelBeforeDay is CancelBeforeDay Setter
// 不推荐使用，使用改规则
func (r *TaobaoxhotelrateplanaddAPIRequest) SetCancelBeforeDay(_cancelBeforeDay int64) error {
	r._cancelBeforeDay = _cancelBeforeDay
	r.Set("cancel_before_day", _cancelBeforeDay)
	return nil
}

// GetCancelBeforeDay CancelBeforeDay Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetCancelBeforeDay() int64 {
	return r._cancelBeforeDay
}

// SetMaxChildAge is MaxChildAge Setter
// 儿童最大年龄(0-18)
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMaxChildAge(_maxChildAge int64) error {
	r._maxChildAge = _maxChildAge
	r.Set("max_child_age", _maxChildAge)
	return nil
}

// GetMaxChildAge MaxChildAge Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMaxChildAge() int64 {
	return r._maxChildAge
}

// SetMinChildAge is MinChildAge Setter
// 儿童最小年龄(0-18)
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMinChildAge(_minChildAge int64) error {
	r._minChildAge = _minChildAge
	r.Set("min_child_age", _minChildAge)
	return nil
}

// GetMinChildAge MinChildAge Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMinChildAge() int64 {
	return r._minChildAge
}

// SetMaxInfantAge is MaxInfantAge Setter
// 婴儿最大年龄(0-18)
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMaxInfantAge(_maxInfantAge int64) error {
	r._maxInfantAge = _maxInfantAge
	r.Set("max_infant_age", _maxInfantAge)
	return nil
}

// GetMaxInfantAge MaxInfantAge Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMaxInfantAge() int64 {
	return r._maxInfantAge
}

// SetMinInfantAge is MinInfantAge Setter
// 婴儿最小年龄(0-18)
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMinInfantAge(_minInfantAge int64) error {
	r._minInfantAge = _minInfantAge
	r.Set("min_infant_age", _minInfantAge)
	return nil
}

// GetMinInfantAge MinInfantAge Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMinInfantAge() int64 {
	return r._minInfantAge
}

// SetIsStudent is IsStudent Setter
// 是否学生价，0：否；1：是。
func (r *TaobaoxhotelrateplanaddAPIRequest) SetIsStudent(_isStudent int64) error {
	r._isStudent = _isStudent
	r.Set("is_student", _isStudent)
	return nil
}

// GetIsStudent IsStudent Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetIsStudent() int64 {
	return r._isStudent
}

// SetHid is Hid Setter
// 酒店id
func (r *TaobaoxhotelrateplanaddAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetHid() int64 {
	return r._hid
}

// SetRid is Rid Setter
// 房型id
func (r *TaobaoxhotelrateplanaddAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetRid() int64 {
	return r._rid
}

// SetSuperRpFlag is SuperRpFlag Setter
// super rp标记，1是；0否
func (r *TaobaoxhotelrateplanaddAPIRequest) SetSuperRpFlag(_superRpFlag int64) error {
	r._superRpFlag = _superRpFlag
	r.Set("super_rp_flag", _superRpFlag)
	return nil
}

// GetSuperRpFlag SuperRpFlag Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetSuperRpFlag() int64 {
	return r._superRpFlag
}

// SetBaseRpFlag is BaseRpFlag Setter
// base rp标记，1是；0否
func (r *TaobaoxhotelrateplanaddAPIRequest) SetBaseRpFlag(_baseRpFlag int64) error {
	r._baseRpFlag = _baseRpFlag
	r.Set("base_rp_flag", _baseRpFlag)
	return nil
}

// GetBaseRpFlag BaseRpFlag Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetBaseRpFlag() int64 {
	return r._baseRpFlag
}

// SetGuaranteeMode is GuaranteeMode Setter
// 2 VCC担保 1 PCI担保 0 支付宝担保(信用住产品担保方式只能为支付宝担保)
func (r *TaobaoxhotelrateplanaddAPIRequest) SetGuaranteeMode(_guaranteeMode int64) error {
	r._guaranteeMode = _guaranteeMode
	r.Set("guarantee_mode", _guaranteeMode)
	return nil
}

// GetGuaranteeMode GuaranteeMode Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetGuaranteeMode() int64 {
	return r._guaranteeMode
}

// SetParentRpid is ParentRpid Setter
// 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
func (r *TaobaoxhotelrateplanaddAPIRequest) SetParentRpid(_parentRpid int64) error {
	r._parentRpid = _parentRpid
	r.Set("parent_rpid", _parentRpid)
	return nil
}

// GetParentRpid ParentRpid Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetParentRpid() int64 {
	return r._parentRpid
}

// SetSource is Source Setter
// 来源
func (r *TaobaoxhotelrateplanaddAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetSource() int64 {
	return r._source
}

// SetBottomPriceFlag is BottomPriceFlag Setter
// 是否底价加价，1是底价加价,0 非底价加价rp
func (r *TaobaoxhotelrateplanaddAPIRequest) SetBottomPriceFlag(_bottomPriceFlag int64) error {
	r._bottomPriceFlag = _bottomPriceFlag
	r.Set("bottom_price_flag", _bottomPriceFlag)
	return nil
}

// GetBottomPriceFlag BottomPriceFlag Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetBottomPriceFlag() int64 {
	return r._bottomPriceFlag
}

// SetMemDiscFlag is MemDiscFlag Setter
// 会员价支持标识,1表示支持会员价规则
func (r *TaobaoxhotelrateplanaddAPIRequest) SetMemDiscFlag(_memDiscFlag int64) error {
	r._memDiscFlag = _memDiscFlag
	r.Set("mem_disc_flag", _memDiscFlag)
	return nil
}

// GetMemDiscFlag MemDiscFlag Getter
func (r TaobaoxhotelrateplanaddAPIRequest) GetMemDiscFlag() int64 {
	return r._memDiscFlag
}
