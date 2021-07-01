package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateAddAPIRequest
新增专享房价 API请求
taobao.xhotel.rate.add

酒店产品库rate添加 */
type TaobaoXhotelRateAddAPIRequest struct {
	model.Params
	// gid酒店商品id
	_gid int64
	// 酒店RPID
	_rpid int64
	// 名称
	_name string
	// 额外服务-是否可以加床，1：不可以，2：可以
	_addBed int64
	// 额外服务-加床价格
	_addBedPrice int64
	// 币种（仅支持CNY）
	_currencyCode int64
	// 实价有房标签（RP支付类型为全额支付）
	_shijiaTag int64
	// 价格和库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
	_inventoryPrice string
	// “即时确认”标识，此类商品预订后直接发货。
	_jishiquerenTag int64
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
	_vendor string
	// 卖家房型ID, 这是卖家自己系统中的房型ID，注意：需按照规则组合
	_outRid string
	// 在添加新rate时，同时添加rate开关日历。可以只设定想设定的某些天，可以不连续。date：开关状态控制的是那一天rate_status：开关状态。0，关闭；1，打开
	_rateSwitchCal string
	// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockEndTime string
	// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockStartTime string
	// 币种信息,默认是CNY,  @see com.taobao.trip.hotel.model.enums.CurrencyEnum
	_currencyCodeName string
	// 操作人信息
	_operator string
	// 默认是2 ,
	_source int64
	// 1是开,0是关, 不填默认是开, rate状态
	_status int64
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
}

// New
