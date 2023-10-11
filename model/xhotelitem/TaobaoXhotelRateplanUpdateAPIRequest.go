package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanUpdateAPIRequest 价格计划rateplan更新或添加 API请求
// taobao.xhotel.rateplan.update
//
// 酒店产品库rateplan更新或添加
type TaobaoXhotelRateplanUpdateAPIRequest struct {
	model.Params
	// 在淘宝搜索页面展示的房价名称；（添加RP时为必须）。注意该名称不要包含早餐相关信息，如果想维护早餐信息，请设置breakfast_count字段即可。
	_name string
	// RP的英文名称
	_englishName string
	// 产品每日开始销售时间，start_time一定为当天时间
	_startTime string
	// 产品每日开始销售时间，start_time一定为当天时间
	_endTime string
	// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款
	_cancelPolicy string
	// 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
	_guaranteeStartTime string
	// 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
	_memberLevel string
	// 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪全渠道（选择H，可实现飞猪、高德、支付宝、手淘均可售卖） O:钉钉商旅 。如果有多个用","分开，比如H,O。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
	_channel string
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// 在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。>=0,<=99 如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_breakfastCal string
	// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_cancelPolicyCal string
	// 在更新rateplan的同时，新增或更新担保日历。 date：担保日历的某一天。 guarantee:担保政策。 其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm" 。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点） 请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_guaranteeCal string
	// 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
	_effectiveTime string
	// 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
	_deadlineTime string
	// rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
	_rpType string
	// 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
	_hourage string
	// 最晚可选入住时间，小时房特有字段。格式为HH:mm
	_canCheckinEnd string
	// 最早可选入住时间，小时房特有字段。格式为HH:mm
	_canCheckinStart string
	// 外部房型id
	_outRid string
	// 外部酒店id
	_outHid string
	// operator
	_operator string
	// 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
	_parentRpCode string
	// 更新RP时的 打标和去标 需求, 0 就是 去标, 1 就是打标,    key的含义:    non-direct-RP 表示非直连RP,    super-could-price-change-RP 表示rp的super标，打上这个tag，表明这个rateplan下单的时候支持变价，商家系统直接放开价格校验。   base-could-derived-RP 表示base rateplan标签，打上了这个tag，表明这是一个base的rateplan，基于该rateplan可以衍生出子rateplan,       ebk-tail-room-RP 表示 ebk尾房rate plan级别标, free-room 表示免房商品
	_tagJson string
	// 协议保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
	_allotmentReleaseTime string
	// 是否包房RP 1包房RP,0 非包房rp
	_packRoomFlag string
	// 是否底价加价，1是底价加价,0 非底价加价rp
	_bottomPriceFlag string
	// 价格计划名称name通过加工处理以后的值
	_displayName string
	// 普通保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
	_commonAllotReleaseTime string
	// 酒店-企业-rp映射实体集合
	_hotelCompanyMappingDOS string
	// 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
	_resourceType string
	// 最晚可选离店时间，小时房特有字段。格式为HH:mm
	_canCheckoutEnd string
	// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)。该字段使用需要联系小二
	_memberDiscountCal string
	// 卖点
	_benefits string
	// 活动类型。1通兑 4超级房券 8直连免房
	_activityType string
	// 入住人限制
	_guestLimit string
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
	// rp的权益信息。1. 额外积分 2. 优惠价格 3. 套餐权益 4.行政礼遇
	_rights string
	// 商业化充值类型 seller充值到卖家 hotel充值到门店
	_freeRoomChargeDstRole string
	// 不推荐使用，使用ratePlanCode来标识要修改的RP
	_rpid int64
	// -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）
	_breakfastCount int64
	// 最小入住天数（1-90）。默认1,小时房RP请设置为1
	_minDays int64
	// 最大入住天数（1-90）。默认90,信用住最大入住天数不超过9天,小时房RP请设置为1,个别卖家支持180
	_maxDays int64
	// 首日入住房间数（1-99）。默认1。【废弃】
	_minAmount int64
	// 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
	_minAdvHours int64
	// 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
	_maxAdvHours int64
	// 1：开启（默认）2：关闭。如果没传值那么默认默认值为1；（添加RP时为必须）
	_status int64
	// 担保类型，只支持： 0 无担保 1 峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
	_guaranteeType int64
	// 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0-代表否。不填写代表不更新该字段。
	_firstStay int64
	// 价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
	_agreement int64
	// 支付类型，只支持：1：预付5：现付6: 信用住7:在线预约8:在线预约信用住。其中5,6,7,8三种类型需要申请权限
	_paymentType int64
	// 学生价，1：是；0：否
	_isStudent int64
	// 酒店id
	_hid int64
	// 房型id
	_rid int64
	// super rp标记，1是；0否
	_superRpFlag int64
	// base rp标记，1是；0否
	_baseRpFlag int64
	// -9999 清空担保, 2 VCC担保, 1 PCI担保，0 支付宝担保(信用住产品担保方式只能为支付宝担保)
	_guaranteeMode int64
	// 父rpid，使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
	_parentRpid int64
	// 来源
	_source int64
	// 是否企业托管RP 0-普通rp,1-企业托管rp
	_companyAssist int64
	// 会员价支持标识,1表示支持会员价规则
	_memDiscFlag int64
}

// NewTaobaoXhotelRateplanUpdateRequest 初始化TaobaoXhotelRateplanUpdateAPIRequest对象
func NewTaobaoXhotelRateplanUpdateRequest() *TaobaoXhotelRateplanUpdateAPIRequest {
	return &TaobaoXhotelRateplanUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rateplan.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 在淘宝搜索页面展示的房价名称；（添加RP时为必须）。注意该名称不要包含早餐相关信息，如果想维护早餐信息，请设置breakfast_count字段即可。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetName() string {
	return r._name
}

// SetEnglishName is EnglishName Setter
// RP的英文名称
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetEnglishName(_englishName string) error {
	r._englishName = _englishName
	r.Set("english_name", _englishName)
	return nil
}

// GetEnglishName EnglishName Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetEnglishName() string {
	return r._englishName
}

// SetStartTime is StartTime Setter
// 产品每日开始销售时间，start_time一定为当天时间
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 产品每日开始销售时间，start_time一定为当天时间
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCancelPolicy is CancelPolicy Setter
// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100&gt;Y&gt;=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100&gt;Y&gt;=0（不超过10条）。1.表示任意退{&#34;cancelPolicyType&#34;:1};2.表示不能退{&#34;cancelPolicyType&#34;:2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{&#34;cancelPolicyType&#34;:4,&#34;policyInfo&#34;:{&#34;48&#34;:10,&#34;24&#34;:20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{&#34;cancelPolicyType&#34;:5,&#34;policyInfo&#34;:{&#34;timeBefore&#34;:6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{&#34;cancelPolicyType&#34;:6,&#34;policyInfo&#34;:{&#34;14&#34;:1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCancelPolicy(_cancelPolicy string) error {
	r._cancelPolicy = _cancelPolicy
	r.Set("cancel_policy", _cancelPolicy)
	return nil
}

// GetCancelPolicy CancelPolicy Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCancelPolicy() string {
	return r._cancelPolicy
}

// SetGuaranteeStartTime is GuaranteeStartTime Setter
// 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetGuaranteeStartTime(_guaranteeStartTime string) error {
	r._guaranteeStartTime = _guaranteeStartTime
	r.Set("guarantee_start_time", _guaranteeStartTime)
	return nil
}

// GetGuaranteeStartTime GuaranteeStartTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetGuaranteeStartTime() string {
	return r._guaranteeStartTime
}

// SetMemberLevel is MemberLevel Setter
// 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMemberLevel(_memberLevel string) error {
	r._memberLevel = _memberLevel
	r.Set("member_level", _memberLevel)
	return nil
}

// GetMemberLevel MemberLevel Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMemberLevel() string {
	return r._memberLevel
}

// SetChannel is Channel Setter
// 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪全渠道（选择H，可实现飞猪、高德、支付宝、手淘均可售卖） O:钉钉商旅 。如果有多个用&#34;,&#34;分开，比如H,O。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetChannel() string {
	return r._channel
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetBreakfastCal is BreakfastCal Setter
// 在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。&gt;=0,&lt;=99 如果date为空，那么会去读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetBreakfastCal(_breakfastCal string) error {
	r._breakfastCal = _breakfastCal
	r.Set("breakfast_cal", _breakfastCal)
	return nil
}

// GetBreakfastCal BreakfastCal Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetBreakfastCal() string {
	return r._breakfastCal
}

// SetCancelPolicyCal is CancelPolicyCal Setter
// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为&#34;yyyy-MM-dd&#34; cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCancelPolicyCal(_cancelPolicyCal string) error {
	r._cancelPolicyCal = _cancelPolicyCal
	r.Set("cancel_policy_cal", _cancelPolicyCal)
	return nil
}

// GetCancelPolicyCal CancelPolicyCal Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCancelPolicyCal() string {
	return r._cancelPolicyCal
}

// SetGuaranteeCal is GuaranteeCal Setter
// 在更新rateplan的同时，新增或更新担保日历。 date：担保日历的某一天。 guarantee:担保政策。 其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为&#34;HH:mm&#34; 。如果date为空，那么会读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点） 请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetGuaranteeCal(_guaranteeCal string) error {
	r._guaranteeCal = _guaranteeCal
	r.Set("guarantee_cal", _guaranteeCal)
	return nil
}

// GetGuaranteeCal GuaranteeCal Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetGuaranteeCal() string {
	return r._guaranteeCal
}

// SetEffectiveTime is EffectiveTime Setter
// 生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetEffectiveTime(_effectiveTime string) error {
	r._effectiveTime = _effectiveTime
	r.Set("effective_time", _effectiveTime)
	return nil
}

// GetEffectiveTime EffectiveTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetEffectiveTime() string {
	return r._effectiveTime
}

// SetDeadlineTime is DeadlineTime Setter
// 生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetDeadlineTime(_deadlineTime string) error {
	r._deadlineTime = _deadlineTime
	r.Set("deadline_time", _deadlineTime)
	return nil
}

// GetDeadlineTime DeadlineTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetDeadlineTime() string {
	return r._deadlineTime
}

// SetRpType is RpType Setter
// rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetRpType(_rpType string) error {
	r._rpType = _rpType
	r.Set("rp_type", _rpType)
	return nil
}

// GetRpType RpType Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetRpType() string {
	return r._rpType
}

// SetHourage is Hourage Setter
// 小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetHourage(_hourage string) error {
	r._hourage = _hourage
	r.Set("hourage", _hourage)
	return nil
}

// GetHourage Hourage Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetHourage() string {
	return r._hourage
}

// SetCanCheckinEnd is CanCheckinEnd Setter
// 最晚可选入住时间，小时房特有字段。格式为HH:mm
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCanCheckinEnd(_canCheckinEnd string) error {
	r._canCheckinEnd = _canCheckinEnd
	r.Set("can_checkin_end", _canCheckinEnd)
	return nil
}

// GetCanCheckinEnd CanCheckinEnd Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCanCheckinEnd() string {
	return r._canCheckinEnd
}

// SetCanCheckinStart is CanCheckinStart Setter
// 最早可选入住时间，小时房特有字段。格式为HH:mm
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCanCheckinStart(_canCheckinStart string) error {
	r._canCheckinStart = _canCheckinStart
	r.Set("can_checkin_start", _canCheckinStart)
	return nil
}

// GetCanCheckinStart CanCheckinStart Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCanCheckinStart() string {
	return r._canCheckinStart
}

// SetOutRid is OutRid Setter
// 外部房型id
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetOutHid is OutHid Setter
// 外部酒店id
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetOperator is Operator Setter
// operator
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetOperator() string {
	return r._operator
}

// SetParentRpCode is ParentRpCode Setter
// 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetParentRpCode(_parentRpCode string) error {
	r._parentRpCode = _parentRpCode
	r.Set("parent_rp_code", _parentRpCode)
	return nil
}

// GetParentRpCode ParentRpCode Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetParentRpCode() string {
	return r._parentRpCode
}

// SetTagJson is TagJson Setter
// 更新RP时的 打标和去标 需求, 0 就是 去标, 1 就是打标,    key的含义:    non-direct-RP 表示非直连RP,    super-could-price-change-RP 表示rp的super标，打上这个tag，表明这个rateplan下单的时候支持变价，商家系统直接放开价格校验。   base-could-derived-RP 表示base rateplan标签，打上了这个tag，表明这是一个base的rateplan，基于该rateplan可以衍生出子rateplan,       ebk-tail-room-RP 表示 ebk尾房rate plan级别标, free-room 表示免房商品
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetTagJson(_tagJson string) error {
	r._tagJson = _tagJson
	r.Set("tag_json", _tagJson)
	return nil
}

// GetTagJson TagJson Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetTagJson() string {
	return r._tagJson
}

// SetAllotmentReleaseTime is AllotmentReleaseTime Setter
// 协议保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetAllotmentReleaseTime(_allotmentReleaseTime string) error {
	r._allotmentReleaseTime = _allotmentReleaseTime
	r.Set("allotment_release_time", _allotmentReleaseTime)
	return nil
}

// GetAllotmentReleaseTime AllotmentReleaseTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetAllotmentReleaseTime() string {
	return r._allotmentReleaseTime
}

// SetPackRoomFlag is PackRoomFlag Setter
// 是否包房RP 1包房RP,0 非包房rp
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetPackRoomFlag(_packRoomFlag string) error {
	r._packRoomFlag = _packRoomFlag
	r.Set("pack_room_flag", _packRoomFlag)
	return nil
}

// GetPackRoomFlag PackRoomFlag Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetPackRoomFlag() string {
	return r._packRoomFlag
}

// SetBottomPriceFlag is BottomPriceFlag Setter
// 是否底价加价，1是底价加价,0 非底价加价rp
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetBottomPriceFlag(_bottomPriceFlag string) error {
	r._bottomPriceFlag = _bottomPriceFlag
	r.Set("bottom_price_flag", _bottomPriceFlag)
	return nil
}

// GetBottomPriceFlag BottomPriceFlag Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetBottomPriceFlag() string {
	return r._bottomPriceFlag
}

// SetDisplayName is DisplayName Setter
// 价格计划名称name通过加工处理以后的值
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetDisplayName(_displayName string) error {
	r._displayName = _displayName
	r.Set("display_name", _displayName)
	return nil
}

// GetDisplayName DisplayName Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetDisplayName() string {
	return r._displayName
}

// SetCommonAllotReleaseTime is CommonAllotReleaseTime Setter
// 普通保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCommonAllotReleaseTime(_commonAllotReleaseTime string) error {
	r._commonAllotReleaseTime = _commonAllotReleaseTime
	r.Set("common_allot_release_time", _commonAllotReleaseTime)
	return nil
}

// GetCommonAllotReleaseTime CommonAllotReleaseTime Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCommonAllotReleaseTime() string {
	return r._commonAllotReleaseTime
}

// SetHotelCompanyMappingDOS is HotelCompanyMappingDOS Setter
// 酒店-企业-rp映射实体集合
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetHotelCompanyMappingDOS(_hotelCompanyMappingDOS string) error {
	r._hotelCompanyMappingDOS = _hotelCompanyMappingDOS
	r.Set("hotel_company_mapping_d_o_s", _hotelCompanyMappingDOS)
	return nil
}

// GetHotelCompanyMappingDOS HotelCompanyMappingDOS Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetHotelCompanyMappingDOS() string {
	return r._hotelCompanyMappingDOS
}

// SetResourceType is ResourceType Setter
// 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetResourceType(_resourceType string) error {
	r._resourceType = _resourceType
	r.Set("resource_type", _resourceType)
	return nil
}

// GetResourceType ResourceType Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetResourceType() string {
	return r._resourceType
}

// SetCanCheckoutEnd is CanCheckoutEnd Setter
// 最晚可选离店时间，小时房特有字段。格式为HH:mm
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCanCheckoutEnd(_canCheckoutEnd string) error {
	r._canCheckoutEnd = _canCheckoutEnd
	r.Set("can_checkout_end", _canCheckoutEnd)
	return nil
}

// GetCanCheckoutEnd CanCheckoutEnd Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCanCheckoutEnd() string {
	return r._canCheckoutEnd
}

// SetMemberDiscountCal is MemberDiscountCal Setter
// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)。该字段使用需要联系小二
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMemberDiscountCal(_memberDiscountCal string) error {
	r._memberDiscountCal = _memberDiscountCal
	r.Set("member_discount_cal", _memberDiscountCal)
	return nil
}

// GetMemberDiscountCal MemberDiscountCal Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMemberDiscountCal() string {
	return r._memberDiscountCal
}

// SetBenefits is Benefits Setter
// 卖点
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetBenefits(_benefits string) error {
	r._benefits = _benefits
	r.Set("benefits", _benefits)
	return nil
}

// GetBenefits Benefits Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetBenefits() string {
	return r._benefits
}

// SetActivityType is ActivityType Setter
// 活动类型。1通兑 4超级房券 8直连免房
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetActivityType(_activityType string) error {
	r._activityType = _activityType
	r.Set("activity_type", _activityType)
	return nil
}

// GetActivityType ActivityType Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetActivityType() string {
	return r._activityType
}

// SetGuestLimit is GuestLimit Setter
// 入住人限制
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetGuestLimit(_guestLimit string) error {
	r._guestLimit = _guestLimit
	r.Set("guest_limit", _guestLimit)
	return nil
}

// GetGuestLimit GuestLimit Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetGuestLimit() string {
	return r._guestLimit
}

// SetOnlineBookingBindingInfo is OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
	r._onlineBookingBindingInfo = _onlineBookingBindingInfo
	r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
	return nil
}

// GetOnlineBookingBindingInfo OnlineBookingBindingInfo Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetOnlineBookingBindingInfo() string {
	return r._onlineBookingBindingInfo
}

// SetRights is Rights Setter
// rp的权益信息。1. 额外积分 2. 优惠价格 3. 套餐权益 4.行政礼遇
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetRights(_rights string) error {
	r._rights = _rights
	r.Set("rights", _rights)
	return nil
}

// GetRights Rights Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetRights() string {
	return r._rights
}

// SetFreeRoomChargeDstRole is FreeRoomChargeDstRole Setter
// 商业化充值类型 seller充值到卖家 hotel充值到门店
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetFreeRoomChargeDstRole(_freeRoomChargeDstRole string) error {
	r._freeRoomChargeDstRole = _freeRoomChargeDstRole
	r.Set("free_room_charge_dst_role", _freeRoomChargeDstRole)
	return nil
}

// GetFreeRoomChargeDstRole FreeRoomChargeDstRole Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetFreeRoomChargeDstRole() string {
	return r._freeRoomChargeDstRole
}

// SetRpid is Rpid Setter
// 不推荐使用，使用ratePlanCode来标识要修改的RP
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetRpid(_rpid int64) error {
	r._rpid = _rpid
	r.Set("rpid", _rpid)
	return nil
}

// GetRpid Rpid Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetRpid() int64 {
	return r._rpid
}

// SetBreakfastCount is BreakfastCount Setter
// -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetBreakfastCount(_breakfastCount int64) error {
	r._breakfastCount = _breakfastCount
	r.Set("breakfast_count", _breakfastCount)
	return nil
}

// GetBreakfastCount BreakfastCount Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetBreakfastCount() int64 {
	return r._breakfastCount
}

// SetMinDays is MinDays Setter
// 最小入住天数（1-90）。默认1,小时房RP请设置为1
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMinDays(_minDays int64) error {
	r._minDays = _minDays
	r.Set("min_days", _minDays)
	return nil
}

// GetMinDays MinDays Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMinDays() int64 {
	return r._minDays
}

// SetMaxDays is MaxDays Setter
// 最大入住天数（1-90）。默认90,信用住最大入住天数不超过9天,小时房RP请设置为1,个别卖家支持180
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMaxDays(_maxDays int64) error {
	r._maxDays = _maxDays
	r.Set("max_days", _maxDays)
	return nil
}

// GetMaxDays MaxDays Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMaxDays() int64 {
	return r._maxDays
}

// SetMinAmount is MinAmount Setter
// 首日入住房间数（1-99）。默认1。【废弃】
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMinAmount(_minAmount int64) error {
	r._minAmount = _minAmount
	r.Set("min_amount", _minAmount)
	return nil
}

// GetMinAmount MinAmount Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMinAmount() int64 {
	return r._minAmount
}

// SetMinAdvHours is MinAdvHours Setter
// 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMinAdvHours(_minAdvHours int64) error {
	r._minAdvHours = _minAdvHours
	r.Set("min_adv_hours", _minAdvHours)
	return nil
}

// GetMinAdvHours MinAdvHours Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMinAdvHours() int64 {
	return r._minAdvHours
}

// SetMaxAdvHours is MaxAdvHours Setter
// 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMaxAdvHours(_maxAdvHours int64) error {
	r._maxAdvHours = _maxAdvHours
	r.Set("max_adv_hours", _maxAdvHours)
	return nil
}

// GetMaxAdvHours MaxAdvHours Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMaxAdvHours() int64 {
	return r._maxAdvHours
}

// SetStatus is Status Setter
// 1：开启（默认）2：关闭。如果没传值那么默认默认值为1；（添加RP时为必须）
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetGuaranteeType is GuaranteeType Setter
// 担保类型，只支持： 0 无担保 1 峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetGuaranteeType(_guaranteeType int64) error {
	r._guaranteeType = _guaranteeType
	r.Set("guarantee_type", _guaranteeType)
	return nil
}

// GetGuaranteeType GuaranteeType Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetGuaranteeType() int64 {
	return r._guaranteeType
}

// SetFirstStay is FirstStay Setter
// 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0-代表否。不填写代表不更新该字段。
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetFirstStay(_firstStay int64) error {
	r._firstStay = _firstStay
	r.Set("first_stay", _firstStay)
	return nil
}

// GetFirstStay FirstStay Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetFirstStay() int64 {
	return r._firstStay
}

// SetAgreement is Agreement Setter
// 价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetAgreement(_agreement int64) error {
	r._agreement = _agreement
	r.Set("agreement", _agreement)
	return nil
}

// GetAgreement Agreement Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetAgreement() int64 {
	return r._agreement
}

// SetPaymentType is PaymentType Setter
// 支付类型，只支持：1：预付5：现付6: 信用住7:在线预约8:在线预约信用住。其中5,6,7,8三种类型需要申请权限
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetPaymentType(_paymentType int64) error {
	r._paymentType = _paymentType
	r.Set("payment_type", _paymentType)
	return nil
}

// GetPaymentType PaymentType Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetPaymentType() int64 {
	return r._paymentType
}

// SetIsStudent is IsStudent Setter
// 学生价，1：是；0：否
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetIsStudent(_isStudent int64) error {
	r._isStudent = _isStudent
	r.Set("is_student", _isStudent)
	return nil
}

// GetIsStudent IsStudent Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetIsStudent() int64 {
	return r._isStudent
}

// SetHid is Hid Setter
// 酒店id
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetHid() int64 {
	return r._hid
}

// SetRid is Rid Setter
// 房型id
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetRid() int64 {
	return r._rid
}

// SetSuperRpFlag is SuperRpFlag Setter
// super rp标记，1是；0否
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetSuperRpFlag(_superRpFlag int64) error {
	r._superRpFlag = _superRpFlag
	r.Set("super_rp_flag", _superRpFlag)
	return nil
}

// GetSuperRpFlag SuperRpFlag Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetSuperRpFlag() int64 {
	return r._superRpFlag
}

// SetBaseRpFlag is BaseRpFlag Setter
// base rp标记，1是；0否
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetBaseRpFlag(_baseRpFlag int64) error {
	r._baseRpFlag = _baseRpFlag
	r.Set("base_rp_flag", _baseRpFlag)
	return nil
}

// GetBaseRpFlag BaseRpFlag Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetBaseRpFlag() int64 {
	return r._baseRpFlag
}

// SetGuaranteeMode is GuaranteeMode Setter
// -9999 清空担保, 2 VCC担保, 1 PCI担保，0 支付宝担保(信用住产品担保方式只能为支付宝担保)
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetGuaranteeMode(_guaranteeMode int64) error {
	r._guaranteeMode = _guaranteeMode
	r.Set("guarantee_mode", _guaranteeMode)
	return nil
}

// GetGuaranteeMode GuaranteeMode Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetGuaranteeMode() int64 {
	return r._guaranteeMode
}

// SetParentRpid is ParentRpid Setter
// 父rpid，使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetParentRpid(_parentRpid int64) error {
	r._parentRpid = _parentRpid
	r.Set("parent_rpid", _parentRpid)
	return nil
}

// GetParentRpid ParentRpid Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetParentRpid() int64 {
	return r._parentRpid
}

// SetSource is Source Setter
// 来源
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetSource() int64 {
	return r._source
}

// SetCompanyAssist is CompanyAssist Setter
// 是否企业托管RP 0-普通rp,1-企业托管rp
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetCompanyAssist(_companyAssist int64) error {
	r._companyAssist = _companyAssist
	r.Set("company_assist", _companyAssist)
	return nil
}

// GetCompanyAssist CompanyAssist Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetCompanyAssist() int64 {
	return r._companyAssist
}

// SetMemDiscFlag is MemDiscFlag Setter
// 会员价支持标识,1表示支持会员价规则
func (r *TaobaoXhotelRateplanUpdateAPIRequest) SetMemDiscFlag(_memDiscFlag int64) error {
	r._memDiscFlag = _memDiscFlag
	r.Set("mem_disc_flag", _memDiscFlag)
	return nil
}

// GetMemDiscFlag MemDiscFlag Getter
func (r TaobaoXhotelRateplanUpdateAPIRequest) GetMemDiscFlag() int64 {
	return r._memDiscFlag
}
