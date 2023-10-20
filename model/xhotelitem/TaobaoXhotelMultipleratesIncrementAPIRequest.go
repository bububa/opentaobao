package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultipleratesIncrementAPIRequest 复杂房价推送接口（批量增量） API请求
// taobao.xhotel.multiplerates.increment
//
// 复杂房价批量增量更新，只会更新指定日期的信息
// 完全涵盖了taobao.xhotel.rates.increment接口的功能
type TaobaoXhotelMultipleratesIncrementAPIRequest struct {
	model.Params
	// 批量全量修改价格和库存信息,会以请求参数中的数据覆盖掉原来报价库存数据。A:useRoomInventory:是否使用room级别共享库存，可选值 true false  2、false时：使用rate级别私有库存，此时如果填写了库存，那么会写入库存表。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:basePrice 基本价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存） 支持状态库存， 60000(状态库存关) 61000(状态库存开);E:occupancy为入住人数，范围为1~10;F:lengthofStay为连住天数，范围为1~10；G:taxAndFee为总税费；H:addBedPrice为加床价；I:addPersonPrice为加人价；J:rateSwitch为开关房状态，1为开房，0为关房。K:支持outRoomId和ratePlanCode来更新报价库存。L:childnum为儿童人数。M:infantnum为婴儿人数。N:ckinSwitch为入住开关(0，关闭；1，打开) O：ckoutSwitch为离店开关 (0，关闭；1，打开) P:lockStartTime锁库存开始时间 Q:lockEndTime锁库存截止时间。childRule：儿童价规则。其中calculateType为计算类型1为固定金额，2为房费的百分比；ageRange为适用儿童的年龄范围，格式为：2~10，即适用2到10岁的儿童；childRange适用几位儿童，格式为1即适用1位儿童；feeType费用类型，如果calculateType=1那么feeType存一固定金额，单位为分，如果calculateType=2那么feeType存房费的百分比，格式为0.2。childAges：儿童年龄范围，格式为2~10，意味着所有儿童价格规则中的适用儿童年龄必须在这个范围之内。
	_rateQuotaMap string
}

// NewTaobaoXhotelMultipleratesIncrementRequest 初始化TaobaoXhotelMultipleratesIncrementAPIRequest对象
func NewTaobaoXhotelMultipleratesIncrementRequest() *TaobaoXhotelMultipleratesIncrementAPIRequest {
	return &TaobaoXhotelMultipleratesIncrementAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelMultipleratesIncrementAPIRequest) Reset() {
	r._rateQuotaMap = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.multiplerates.increment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRateQuotaMap is RateQuotaMap Setter
// 批量全量修改价格和库存信息,会以请求参数中的数据覆盖掉原来报价库存数据。A:useRoomInventory:是否使用room级别共享库存，可选值 true false  2、false时：使用rate级别私有库存，此时如果填写了库存，那么会写入库存表。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:basePrice 基本价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存） 支持状态库存， 60000(状态库存关) 61000(状态库存开);E:occupancy为入住人数，范围为1~10;F:lengthofStay为连住天数，范围为1~10；G:taxAndFee为总税费；H:addBedPrice为加床价；I:addPersonPrice为加人价；J:rateSwitch为开关房状态，1为开房，0为关房。K:支持outRoomId和ratePlanCode来更新报价库存。L:childnum为儿童人数。M:infantnum为婴儿人数。N:ckinSwitch为入住开关(0，关闭；1，打开) O：ckoutSwitch为离店开关 (0，关闭；1，打开) P:lockStartTime锁库存开始时间 Q:lockEndTime锁库存截止时间。childRule：儿童价规则。其中calculateType为计算类型1为固定金额，2为房费的百分比；ageRange为适用儿童的年龄范围，格式为：2~10，即适用2到10岁的儿童；childRange适用几位儿童，格式为1即适用1位儿童；feeType费用类型，如果calculateType=1那么feeType存一固定金额，单位为分，如果calculateType=2那么feeType存房费的百分比，格式为0.2。childAges：儿童年龄范围，格式为2~10，意味着所有儿童价格规则中的适用儿童年龄必须在这个范围之内。
func (r *TaobaoXhotelMultipleratesIncrementAPIRequest) SetRateQuotaMap(_rateQuotaMap string) error {
	r._rateQuotaMap = _rateQuotaMap
	r.Set("rate_quota_map", _rateQuotaMap)
	return nil
}

// GetRateQuotaMap RateQuotaMap Getter
func (r TaobaoXhotelMultipleratesIncrementAPIRequest) GetRateQuotaMap() string {
	return r._rateQuotaMap
}

var poolTaobaoXhotelMultipleratesIncrementAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelMultipleratesIncrementRequest()
	},
}

// GetTaobaoXhotelMultipleratesIncrementRequest 从 sync.Pool 获取 TaobaoXhotelMultipleratesIncrementAPIRequest
func GetTaobaoXhotelMultipleratesIncrementAPIRequest() *TaobaoXhotelMultipleratesIncrementAPIRequest {
	return poolTaobaoXhotelMultipleratesIncrementAPIRequest.Get().(*TaobaoXhotelMultipleratesIncrementAPIRequest)
}

// ReleaseTaobaoXhotelMultipleratesIncrementAPIRequest 将 TaobaoXhotelMultipleratesIncrementAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelMultipleratesIncrementAPIRequest(v *TaobaoXhotelMultipleratesIncrementAPIRequest) {
	v.Reset()
	poolTaobaoXhotelMultipleratesIncrementAPIRequest.Put(v)
}
