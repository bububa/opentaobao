package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelratesincrementAPIRequest 价格推送接口（批量增量） API请求
// taobao.xhotel.rates.increment
//
// Rate库存&amp;价格增量更新接口，用户仅需要更新Rate中发生变化的库存日历&amp;价格日历即可
type TaobaoxhotelratesincrementAPIRequest struct {
	model.Params
	// 批量修改价格和房价专有库存信息，json格式,可同时修改多套房型+价格计划的价格：A:use_room_inventory:是否使用房型共享库存，可选值 true false 1、true时：使用房型共享库存 2、false时：使用房价专有库存，此时要求房价专有库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），不能重复。 C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 房价专有库存 int 类型 取值范围  0-999（数量库存） 60000(状态库存关) 61000(状态库存开) E:status 价格开关，0为关，1为开。lock_start_time为锁库存开始时间，lock_end_time为锁库存结束时间，如果当前时间在这个时间范围内，那么不允许修改库存。示例值：（1）[{"out_rid":"ABCDE_123","rateplan_code":"ABCDE_WHL01","vendor":"","lock_start_time":"","lock_end_time":"","data":{"use_room_inventory":false,"inventory_price":[{"date":"2013-11-18","quota":1,"price":1000,"status":1},{"date":"2013-11-19","quota":1,"price":1000,"status":0}]}},{"out_rid":"ABCDE_234","rateplan_code":"ABCDE_WHL01","vendor":"","data":{"use_room_inventory":false,"inventory_price":[{"date":"2013-11-18","quota":1,"price":1000,"status":1},{"date":"2013-11-19","quota":1,"price":1000,"status":0}]}}]
	_rateInventoryPriceMap string
}

// NewTaobaoxhotelratesincrementRequest 初始化TaobaoxhotelratesincrementAPIRequest对象
func NewTaobaoxhotelratesincrementRequest() *TaobaoxhotelratesincrementAPIRequest {
	return &TaobaoxhotelratesincrementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelratesincrementAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rates.increment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelratesincrementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelratesincrementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRateInventoryPriceMap is RateInventoryPriceMap Setter
// 批量修改价格和房价专有库存信息，json格式,可同时修改多套房型+价格计划的价格：A:use_room_inventory:是否使用房型共享库存，可选值 true false 1、true时：使用房型共享库存 2、false时：使用房价专有库存，此时要求房价专有库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），不能重复。 C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 房价专有库存 int 类型 取值范围  0-999（数量库存） 60000(状态库存关) 61000(状态库存开) E:status 价格开关，0为关，1为开。lock_start_time为锁库存开始时间，lock_end_time为锁库存结束时间，如果当前时间在这个时间范围内，那么不允许修改库存。示例值：（1）[{&#34;out_rid&#34;:&#34;ABCDE_123&#34;,&#34;rateplan_code&#34;:&#34;ABCDE_WHL01&#34;,&#34;vendor&#34;:&#34;&#34;,&#34;lock_start_time&#34;:&#34;&#34;,&#34;lock_end_time&#34;:&#34;&#34;,&#34;data&#34;:{&#34;use_room_inventory&#34;:false,&#34;inventory_price&#34;:[{&#34;date&#34;:&#34;2013-11-18&#34;,&#34;quota&#34;:1,&#34;price&#34;:1000,&#34;status&#34;:1},{&#34;date&#34;:&#34;2013-11-19&#34;,&#34;quota&#34;:1,&#34;price&#34;:1000,&#34;status&#34;:0}]}},{&#34;out_rid&#34;:&#34;ABCDE_234&#34;,&#34;rateplan_code&#34;:&#34;ABCDE_WHL01&#34;,&#34;vendor&#34;:&#34;&#34;,&#34;data&#34;:{&#34;use_room_inventory&#34;:false,&#34;inventory_price&#34;:[{&#34;date&#34;:&#34;2013-11-18&#34;,&#34;quota&#34;:1,&#34;price&#34;:1000,&#34;status&#34;:1},{&#34;date&#34;:&#34;2013-11-19&#34;,&#34;quota&#34;:1,&#34;price&#34;:1000,&#34;status&#34;:0}]}}]
func (r *TaobaoxhotelratesincrementAPIRequest) SetRateInventoryPriceMap(_rateInventoryPriceMap string) error {
	r._rateInventoryPriceMap = _rateInventoryPriceMap
	r.Set("rate_inventory_price_map", _rateInventoryPriceMap)
	return nil
}

// GetRateInventoryPriceMap RateInventoryPriceMap Getter
func (r TaobaoxhotelratesincrementAPIRequest) GetRateInventoryPriceMap() string {
	return r._rateInventoryPriceMap
}
