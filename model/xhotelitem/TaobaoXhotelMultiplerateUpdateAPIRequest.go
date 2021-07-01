package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelMultiplerateUpdateAPIRequest
复杂价格推送接口（全量更新） API请求
taobao.xhotel.multiplerate.update

酒店产品库复杂rate（多人价，连住价等）更新
同时完全涵盖taobao.xhotel.rate.update的功能 */
type TaobaoXhotelMultiplerateUpdateAPIRequest struct {
	model.Params
	// 废弃，使用out_rid
	_gid int64
	// 卖家房型ID
	_outRid string
	// 废弃，使用rate_plan_code
	_rpid int64
	// 卖家自己系统的房价code
	_ratePlanCode string
	// 废弃
	_name string
	// 入住人数(范围1~10)
	_occupancy int64
	// 连住天数(范围1~5)
	_lengthofstay int64
	// 价格和库存信息。 A:use_room_inventory:是否使用房型共享库存，可选值 true false ,false时：使用房价专有库存，此时要求价格和库存必填。 B:date 日期必须为 T---T+180 日内的日期（T为当天），且不能重复 C:price 价格 int类型 取值范围1-99999999 单位为分 D:quota 库存 int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) tax为税费，addBed为加床价，addPerson为加人价1,若连住大于1，price请推送总价
	_inventoryPrice string
	// 价格开关 date：开关状态控制的那一天；rate_status：开关状态(0，关闭；1，打开); checkin_status：入住开关(0，关闭；1，打开)；checkout_status：离店开关 (0，关闭；1，打开)
	_rateSwitchCal string
	// 币种.CNY为人民币
	_currencyCode string
	// 价格状态。0为不可售；1为可售，默认可售
	_status int64
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 儿童人数
	_childnum int64
	// 婴儿人数
	_infantnum int64
	// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockEndTime string
	// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
	_lockStartTime string
	// 在线预约关联关系推送，priceRuleNumber：加价规则序号
	_onlineBookingBindingInfo string
}

// New
