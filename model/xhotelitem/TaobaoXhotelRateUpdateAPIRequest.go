package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelrateupdateAPIRequest 价格推送接口（全量更新） API请求
// taobao.xhotel.rate.update
//
// 酒店产品库rate更新
type TaobaoxhotelrateupdateAPIRequest struct {
	model.Params
	// 废弃
	_name string
	// 每日价格和房价专有库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
	_inventoryPrice string
	// 系统商，一般不用填写，使用需要申请
	_vendor string
	// 商家价格计划编码
	_rateplanCode string
	// 商家房型ID
	_outRid string
	// 日历价格开关， date：开关状态控制的是那一天 rate_status：开关状态。0，关闭；1，打开
	_rateSwitchCal string
	// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockEndTime string
	// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockStartTime string
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
	//  是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&docType=1&articleId=121402&previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{"url":"https://xxxxx/","isMain":true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{"detailName":"免费寄存","detailValue":[""],"type":"single","priority":1}]
	_hotelXitemInfos string
	// 不推荐使用，请使用out_rid
	_gid int64
	// 不推荐使用，请使用rateplan_code
	_rpid int64
	// 不推荐使用
	_addBed int64
	// 不推荐使用
	_addBedPrice int64
	// 不推荐使用（仅支持CNY）
	_currencyCode int64
	// 不推荐使用
	_shijiaTag int64
	// 不推荐使用
	_jishiquerenTag int64
}

// NewTaobaoxhotelrateupdateRequest 初始化TaobaoxhotelrateupdateAPIRequest对象
func NewTaobaoxhotelrateupdateRequest() *TaobaoxhotelrateupdateAPIRequest {
	return &TaobaoxhotelrateupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelrateupdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelrateupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelrateupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 废弃
func (r *TaobaoxhotelrateupdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetName() string {
	return r._name
}

// SetInventoryPrice is InventoryPrice Setter
// 每日价格和房价专有库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
func (r *TaobaoxhotelrateupdateAPIRequest) SetInventoryPrice(_inventoryPrice string) error {
	r._inventoryPrice = _inventoryPrice
	r.Set("inventory_price", _inventoryPrice)
	return nil
}

// GetInventoryPrice InventoryPrice Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetInventoryPrice() string {
	return r._inventoryPrice
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用需要申请
func (r *TaobaoxhotelrateupdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格计划编码
func (r *TaobaoxhotelrateupdateAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetOutRid is OutRid Setter
// 商家房型ID
func (r *TaobaoxhotelrateupdateAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetRateSwitchCal is RateSwitchCal Setter
// 日历价格开关， date：开关状态控制的是那一天 rate_status：开关状态。0，关闭；1，打开
func (r *TaobaoxhotelrateupdateAPIRequest) SetRateSwitchCal(_rateSwitchCal string) error {
	r._rateSwitchCal = _rateSwitchCal
	r.Set("rate_switch_cal", _rateSwitchCal)
	return nil
}

// GetRateSwitchCal RateSwitchCal Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetRateSwitchCal() string {
	return r._rateSwitchCal
}

// SetLockEndTime is LockEndTime Setter
// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoxhotelrateupdateAPIRequest) SetLockEndTime(_lockEndTime string) error {
	r._lockEndTime = _lockEndTime
	r.Set("lock_end_time", _lockEndTime)
	return nil
}

// GetLockEndTime LockEndTime Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetLockEndTime() string {
	return r._lockEndTime
}

// SetLockStartTime is LockStartTime Setter
// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoxhotelrateupdateAPIRequest) SetLockStartTime(_lockStartTime string) error {
	r._lockStartTime = _lockStartTime
	r.Set("lock_start_time", _lockStartTime)
	return nil
}

// GetLockStartTime LockStartTime Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetLockStartTime() string {
	return r._lockStartTime
}

// SetOnlineBookingBindingInfo is OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoxhotelrateupdateAPIRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
	r._onlineBookingBindingInfo = _onlineBookingBindingInfo
	r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
	return nil
}

// GetOnlineBookingBindingInfo OnlineBookingBindingInfo Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetOnlineBookingBindingInfo() string {
	return r._onlineBookingBindingInfo
}

// SetHotelXitemInfos is HotelXitemInfos Setter
//
//	是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&amp;docType=1&amp;articleId=121402&amp;previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{&#34;url&#34;:&#34;https://xxxxx/&#34;,&#34;isMain&#34;:true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{&#34;detailName&#34;:&#34;免费寄存&#34;,&#34;detailValue&#34;:[&#34;&#34;],&#34;type&#34;:&#34;single&#34;,&#34;priority&#34;:1}]
func (r *TaobaoxhotelrateupdateAPIRequest) SetHotelXitemInfos(_hotelXitemInfos string) error {
	r._hotelXitemInfos = _hotelXitemInfos
	r.Set("hotel_xitem_infos", _hotelXitemInfos)
	return nil
}

// GetHotelXitemInfos HotelXitemInfos Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetHotelXitemInfos() string {
	return r._hotelXitemInfos
}

// SetGid is Gid Setter
// 不推荐使用，请使用out_rid
func (r *TaobaoxhotelrateupdateAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetGid() int64 {
	return r._gid
}

// SetRpid is Rpid Setter
// 不推荐使用，请使用rateplan_code
func (r *TaobaoxhotelrateupdateAPIRequest) SetRpid(_rpid int64) error {
	r._rpid = _rpid
	r.Set("rpid", _rpid)
	return nil
}

// GetRpid Rpid Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetRpid() int64 {
	return r._rpid
}

// SetAddBed is AddBed Setter
// 不推荐使用
func (r *TaobaoxhotelrateupdateAPIRequest) SetAddBed(_addBed int64) error {
	r._addBed = _addBed
	r.Set("add_bed", _addBed)
	return nil
}

// GetAddBed AddBed Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetAddBed() int64 {
	return r._addBed
}

// SetAddBedPrice is AddBedPrice Setter
// 不推荐使用
func (r *TaobaoxhotelrateupdateAPIRequest) SetAddBedPrice(_addBedPrice int64) error {
	r._addBedPrice = _addBedPrice
	r.Set("add_bed_price", _addBedPrice)
	return nil
}

// GetAddBedPrice AddBedPrice Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetAddBedPrice() int64 {
	return r._addBedPrice
}

// SetCurrencyCode is CurrencyCode Setter
// 不推荐使用（仅支持CNY）
func (r *TaobaoxhotelrateupdateAPIRequest) SetCurrencyCode(_currencyCode int64) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetCurrencyCode() int64 {
	return r._currencyCode
}

// SetShijiaTag is ShijiaTag Setter
// 不推荐使用
func (r *TaobaoxhotelrateupdateAPIRequest) SetShijiaTag(_shijiaTag int64) error {
	r._shijiaTag = _shijiaTag
	r.Set("shijia_tag", _shijiaTag)
	return nil
}

// GetShijiaTag ShijiaTag Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetShijiaTag() int64 {
	return r._shijiaTag
}

// SetJishiquerenTag is JishiquerenTag Setter
// 不推荐使用
func (r *TaobaoxhotelrateupdateAPIRequest) SetJishiquerenTag(_jishiquerenTag int64) error {
	r._jishiquerenTag = _jishiquerenTag
	r.Set("jishiqueren_tag", _jishiquerenTag)
	return nil
}

// GetJishiquerenTag JishiquerenTag Getter
func (r TaobaoxhotelrateupdateAPIRequest) GetJishiquerenTag() int64 {
	return r._jishiquerenTag
}
