package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultiplerateUpdateAPIRequest 复杂价格推送接口（全量更新） API请求
// taobao.xhotel.multiplerate.update
//
// 酒店产品库复杂rate（多人价，连住价等）更新
// 同时完全涵盖taobao.xhotel.rate.update的功能
type TaobaoXhotelMultiplerateUpdateAPIRequest struct {
	model.Params
	// 卖家房型ID
	_outRid string
	// 价格开关 date：开关状态控制的那一天；rate_status：开关状态(0，关闭；1，打开); checkin_status：入住开关(0，关闭；1，打开)；checkout_status：离店开关 (0，关闭；1，打开)
	_rateSwitchCal string
	// 币种.CNY为人民币
	_currencyCode string
	// 卖家自己系统的房价code
	_ratePlanCode string
	// 价格和库存信息。 A:use_room_inventory:是否使用房型共享库存，可选值 true false ,false时：使用房价专有库存，此时要求价格和库存必填。 B:date 日期必须为 T---T+180 日内的日期（T为当天），且不能重复 C:price 价格 int类型 取值范围1-99999999 单位为分 D:quota 库存 int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) tax为税费，addBed为加床价，addPerson为加人价1,若连住大于1，price请推送总价
	_inventoryPrice string
	// 废弃
	_name string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockEndTime string
	// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockStartTime string
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
	// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款
	_cancelPolicy string
	// 在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。>=0,<=99 如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_breakfastCal string
	// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
	_cancelPolicyCal string
	//  是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&docType=1&articleId=121402&previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{"url":"https://xxxxx/","isMain":true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{"detailName":"免费寄存","detailValue":[""],"type":"single","priority":1}]
	_hotelXitemInfos string
	// 废弃，使用rate_plan_code
	_rpid int64
	// 价格状态。0为不可售；1为可售，默认可售
	_status int64
	// 连住天数(范围1~5)
	_lengthofstay int64
	// 入住人数(范围1~10)
	_occupancy int64
	// 废弃，使用out_rid
	_gid int64
	// 儿童人数
	_childnum int64
	// 婴儿人数
	_infantnum int64
	// -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）
	_breakfast int64
}

// NewTaobaoXhotelMultiplerateUpdateRequest 初始化TaobaoXhotelMultiplerateUpdateAPIRequest对象
func NewTaobaoXhotelMultiplerateUpdateRequest() *TaobaoXhotelMultiplerateUpdateAPIRequest {
	return &TaobaoXhotelMultiplerateUpdateAPIRequest{
		Params: model.NewParams(22),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) Reset() {
	r._outRid = ""
	r._rateSwitchCal = ""
	r._currencyCode = ""
	r._ratePlanCode = ""
	r._inventoryPrice = ""
	r._name = ""
	r._vendor = ""
	r._lockEndTime = ""
	r._lockStartTime = ""
	r._onlineBookingBindingInfo = ""
	r._cancelPolicy = ""
	r._breakfastCal = ""
	r._cancelPolicyCal = ""
	r._hotelXitemInfos = ""
	r._rpid = 0
	r._status = 0
	r._lengthofstay = 0
	r._occupancy = 0
	r._gid = 0
	r._childnum = 0
	r._infantnum = 0
	r._breakfast = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.multiplerate.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRid is OutRid Setter
// 卖家房型ID
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetRateSwitchCal is RateSwitchCal Setter
// 价格开关 date：开关状态控制的那一天；rate_status：开关状态(0，关闭；1，打开); checkin_status：入住开关(0，关闭；1，打开)；checkout_status：离店开关 (0，关闭；1，打开)
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetRateSwitchCal(_rateSwitchCal string) error {
	r._rateSwitchCal = _rateSwitchCal
	r.Set("rate_switch_cal", _rateSwitchCal)
	return nil
}

// GetRateSwitchCal RateSwitchCal Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetRateSwitchCal() string {
	return r._rateSwitchCal
}

// SetCurrencyCode is CurrencyCode Setter
// 币种.CNY为人民币
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetRatePlanCode is RatePlanCode Setter
// 卖家自己系统的房价code
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetRatePlanCode(_ratePlanCode string) error {
	r._ratePlanCode = _ratePlanCode
	r.Set("rate_plan_code", _ratePlanCode)
	return nil
}

// GetRatePlanCode RatePlanCode Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetRatePlanCode() string {
	return r._ratePlanCode
}

// SetInventoryPrice is InventoryPrice Setter
// 价格和库存信息。 A:use_room_inventory:是否使用房型共享库存，可选值 true false ,false时：使用房价专有库存，此时要求价格和库存必填。 B:date 日期必须为 T---T+180 日内的日期（T为当天），且不能重复 C:price 价格 int类型 取值范围1-99999999 单位为分 D:quota 库存 int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) tax为税费，addBed为加床价，addPerson为加人价1,若连住大于1，price请推送总价
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetInventoryPrice(_inventoryPrice string) error {
	r._inventoryPrice = _inventoryPrice
	r.Set("inventory_price", _inventoryPrice)
	return nil
}

// GetInventoryPrice InventoryPrice Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetInventoryPrice() string {
	return r._inventoryPrice
}

// SetName is Name Setter
// 废弃
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetName() string {
	return r._name
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetLockEndTime is LockEndTime Setter
// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetLockEndTime(_lockEndTime string) error {
	r._lockEndTime = _lockEndTime
	r.Set("lock_end_time", _lockEndTime)
	return nil
}

// GetLockEndTime LockEndTime Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetLockEndTime() string {
	return r._lockEndTime
}

// SetLockStartTime is LockStartTime Setter
// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetLockStartTime(_lockStartTime string) error {
	r._lockStartTime = _lockStartTime
	r.Set("lock_start_time", _lockStartTime)
	return nil
}

// GetLockStartTime LockStartTime Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetLockStartTime() string {
	return r._lockStartTime
}

// SetOnlineBookingBindingInfo is OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
	r._onlineBookingBindingInfo = _onlineBookingBindingInfo
	r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
	return nil
}

// GetOnlineBookingBindingInfo OnlineBookingBindingInfo Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetOnlineBookingBindingInfo() string {
	return r._onlineBookingBindingInfo
}

// SetCancelPolicy is CancelPolicy Setter
// 退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100&gt;Y&gt;=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100&gt;Y&gt;=0（不超过10条）。1.表示任意退{&#34;cancelPolicyType&#34;:1};2.表示不能退{&#34;cancelPolicyType&#34;:2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{&#34;cancelPolicyType&#34;:4,&#34;policyInfo&#34;:{&#34;48&#34;:10,&#34;24&#34;:20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{&#34;cancelPolicyType&#34;:5,&#34;policyInfo&#34;:{&#34;timeBefore&#34;:6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{&#34;cancelPolicyType&#34;:6,&#34;policyInfo&#34;:{&#34;14&#34;:1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetCancelPolicy(_cancelPolicy string) error {
	r._cancelPolicy = _cancelPolicy
	r.Set("cancel_policy", _cancelPolicy)
	return nil
}

// GetCancelPolicy CancelPolicy Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetCancelPolicy() string {
	return r._cancelPolicy
}

// SetBreakfastCal is BreakfastCal Setter
// 在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。&gt;=0,&lt;=99 如果date为空，那么会去读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetBreakfastCal(_breakfastCal string) error {
	r._breakfastCal = _breakfastCal
	r.Set("breakfast_cal", _breakfastCal)
	return nil
}

// GetBreakfastCal BreakfastCal Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetBreakfastCal() string {
	return r._breakfastCal
}

// SetCancelPolicyCal is CancelPolicyCal Setter
// 在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为&#34;yyyy-MM-dd&#34; cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为&#34;yyyy-MM-dd&#34;），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetCancelPolicyCal(_cancelPolicyCal string) error {
	r._cancelPolicyCal = _cancelPolicyCal
	r.Set("cancel_policy_cal", _cancelPolicyCal)
	return nil
}

// GetCancelPolicyCal CancelPolicyCal Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetCancelPolicyCal() string {
	return r._cancelPolicyCal
}

// SetHotelXitemInfos is HotelXitemInfos Setter
//
//	是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&amp;docType=1&amp;articleId=121402&amp;previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{&#34;url&#34;:&#34;https://xxxxx/&#34;,&#34;isMain&#34;:true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{&#34;detailName&#34;:&#34;免费寄存&#34;,&#34;detailValue&#34;:[&#34;&#34;],&#34;type&#34;:&#34;single&#34;,&#34;priority&#34;:1}]
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetHotelXitemInfos(_hotelXitemInfos string) error {
	r._hotelXitemInfos = _hotelXitemInfos
	r.Set("hotel_xitem_infos", _hotelXitemInfos)
	return nil
}

// GetHotelXitemInfos HotelXitemInfos Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetHotelXitemInfos() string {
	return r._hotelXitemInfos
}

// SetRpid is Rpid Setter
// 废弃，使用rate_plan_code
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetRpid(_rpid int64) error {
	r._rpid = _rpid
	r.Set("rpid", _rpid)
	return nil
}

// GetRpid Rpid Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetRpid() int64 {
	return r._rpid
}

// SetStatus is Status Setter
// 价格状态。0为不可售；1为可售，默认可售
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetLengthofstay is Lengthofstay Setter
// 连住天数(范围1~5)
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetLengthofstay(_lengthofstay int64) error {
	r._lengthofstay = _lengthofstay
	r.Set("lengthofstay", _lengthofstay)
	return nil
}

// GetLengthofstay Lengthofstay Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetLengthofstay() int64 {
	return r._lengthofstay
}

// SetOccupancy is Occupancy Setter
// 入住人数(范围1~10)
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetOccupancy(_occupancy int64) error {
	r._occupancy = _occupancy
	r.Set("occupancy", _occupancy)
	return nil
}

// GetOccupancy Occupancy Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetOccupancy() int64 {
	return r._occupancy
}

// SetGid is Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetGid() int64 {
	return r._gid
}

// SetChildnum is Childnum Setter
// 儿童人数
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetChildnum(_childnum int64) error {
	r._childnum = _childnum
	r.Set("childnum", _childnum)
	return nil
}

// GetChildnum Childnum Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetChildnum() int64 {
	return r._childnum
}

// SetInfantnum is Infantnum Setter
// 婴儿人数
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetInfantnum(_infantnum int64) error {
	r._infantnum = _infantnum
	r.Set("infantnum", _infantnum)
	return nil
}

// GetInfantnum Infantnum Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetInfantnum() int64 {
	return r._infantnum
}

// SetBreakfast is Breakfast Setter
// -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）
func (r *TaobaoXhotelMultiplerateUpdateAPIRequest) SetBreakfast(_breakfast int64) error {
	r._breakfast = _breakfast
	r.Set("breakfast", _breakfast)
	return nil
}

// GetBreakfast Breakfast Getter
func (r TaobaoXhotelMultiplerateUpdateAPIRequest) GetBreakfast() int64 {
	return r._breakfast
}

var poolTaobaoXhotelMultiplerateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelMultiplerateUpdateRequest()
	},
}

// GetTaobaoXhotelMultiplerateUpdateRequest 从 sync.Pool 获取 TaobaoXhotelMultiplerateUpdateAPIRequest
func GetTaobaoXhotelMultiplerateUpdateAPIRequest() *TaobaoXhotelMultiplerateUpdateAPIRequest {
	return poolTaobaoXhotelMultiplerateUpdateAPIRequest.Get().(*TaobaoXhotelMultiplerateUpdateAPIRequest)
}

// ReleaseTaobaoXhotelMultiplerateUpdateAPIRequest 将 TaobaoXhotelMultiplerateUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelMultiplerateUpdateAPIRequest(v *TaobaoXhotelMultiplerateUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelMultiplerateUpdateAPIRequest.Put(v)
}
