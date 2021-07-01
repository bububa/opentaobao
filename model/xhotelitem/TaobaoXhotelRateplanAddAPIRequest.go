package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateplanAddAPIRequest
酒店产品库rateplan添加 API请求
taobao.xhotel.rateplan.add

酒店产品库rateplan */
type TaobaoXhotelRateplanAddAPIRequest struct {
	model.Params
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 在淘宝搜索页面展示的房价名称。请注意名称里不要维护早餐信息，如果想设置早餐信息，请设置breakfast_count字段即可
	_name string
	// 支付类型，只支持：1：预付5：现付6: 信用住7:预付在线预约8:信用住在线预约。其中5,6,7,8四种类型需要申请权限
	_paymentType int64
	// -1：状态早餐,有具体几人价有关系，几人价是几份早餐;0：不含早1：含单早2：含双早N：含N早（-1-99可选）
	_breakfastCount int64
	// 不推荐使用
	_feeServicePercent int64
	// 最小入住天数（1-90）。默认1,小时房RP请设置为1
	_minDays int64
	// 最大入住天数（1-90）。默认90 信用住不超过9天,小时房RP请设置为1
	_maxDays int64
	// 首日入住房间数（1-99）。默认1。【废弃】
	_minAmount int64
	// 最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。
	_minAdvHours int64
	// 最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。
	_maxAdvHours int64
	// 产品每日开始销售时间，start_time一定为当天时间
	_startTime string
	// 产品每日结束销售时间,当end_time<start_time时，表示end_time为第二天，此时附加限制end_time<=06:00:00并且start_time>=12:00:00,表明可售时间从当天12点到次日的凌晨6点（扩展此信息主要为了描述尾房的rp）注意start_time一定是当天的时间。尾房18：00起可售
	_endTime string
	// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间,多间夜扣款。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
	_cancelPolicy string
	// 1：开启（默认）2：关闭。如果没传值那么默认默认值为1
	_status int64
	// RP的英文名称
	_englishName string
	// 担保类型，只支持： 0  无担保  1  峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
	_guaranteeType int64
	// 分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）
	_guaranteeStartTime string
	// 双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）
	_memberLevel string
	// 销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪 O:钉钉商旅 A:集团内部商旅。如果只投放飞猪，改字段不用填写或者只填H；如果有多个用","分开。如果需要投放其他渠道，请联系飞猪运营或者技术支持。
	_channel string
	// 不推送则默认2人，如有低于2人的RP限制请推送该字段。
	_occupancy int64
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 需申请会员权限。是否是新用户首住优惠rp。1-代表是。0或者不填写代表否
	_firstStay int64
	// 废弃。价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价
	_agreement int64
	// 在添加rateplan时，同时新增早餐日历。date：说明这条记录的早餐政策breakfast_count：这一天早餐的数量。>=-1,<=99。如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_breakfastCal string
	// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814
	_cancelPolicyCal string
	// 在新增rateplan的同时，新增担保日历。date：担保日历的某一天。guarantee:担保政策。其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm"。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_guaranteeCal string
	// 不推荐使用，使用改规则
	_cancelBeforeHour string
	// 不推荐使用，使用改规则
	_cancelBeforeDay int64
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
	// 是否学生价，0：否；1：是。
	_isStudent int64
	// 酒店id
	_hid int64
	// 房型id
	_rid int64
	// 外部房型id
	_outRid string
	// 外部酒店id
	_outHid string
	// super rp标记，1是；0否
	_superRpFlag int64
	// base rp标记，1是；0否
	_baseRpFlag int64
	// 2 VCC担保 1 PCI担保 0 支付宝担保(信用住产品担保方式只能为支付宝担保)
	_guaranteeMode int64
	// 父rpid,使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部
	_parentRpCode string
	// 父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部
	_parentRpid int64
	// 操作rateplan的来源
	_operator string
	// 新增RP时的 打标和去标 需求,
	_tagJson string
	// 来源
	_source int64
	// 保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
	_allotmentReleaseTime string
	// 普通保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认
	_commonAllotReleaseTime string
	// 商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。
	_resourceType string
	// 是否底价加价，1是底价加价,0 非底价加价rp
	_bottomPriceFlag int64
	// 最晚可选离店时间，小时房特有字段。格式为HH:mm
	_canCheckoutEnd string
	// 会员价支持标识,1表示支持会员价规则
	_memDiscFlag int64
	// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)
	_memberDiscountCal string
	// RP入住人限制信息。JSON格式
	_guestLimit string
	// RP参与的活动，3为尾房,4超级房券
	_activityType string
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
}

// New
