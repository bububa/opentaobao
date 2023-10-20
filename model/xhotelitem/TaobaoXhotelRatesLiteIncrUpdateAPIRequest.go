package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRatesLiteIncrUpdateAPIRequest 酒店价格库存轻量级增量接口 API请求
// taobao.xhotel.rates.lite.incr.update
//
// 多个rate的库存房价开关的增量更新接口
type TaobaoXhotelRatesLiteIncrUpdateAPIRequest struct {
	model.Params
	// json格式: [{"rp_code":"031612_181220000074_F1_1","out_rid":"031612_HD","vendor":"taobao","useRoomInv":0,"invAndSwitch":[{"date":"2020-01-09","price":1200,"quota":23,"rateStatus":1},{"date":"2020-01-10","price":1200,"quota":23,"rateStatus":0}]},{"rp_code":"1234","out_rid":"123456","vendor":"taobao","useRoomInv":0,"invAndSwitch":[{"date":"yyyy-MM-dd","price":1200,"quota":23,"rateStatus":1},{"date":"yyyy-MM-dd","price":1200,"quota":23,"rateStatus":1}]}] rp_code:房价code,out_rid:房型code,必传字段 useRoomInv: 表示是否使用房型库存(1表示是,0表示否),可选字段 invAndSwitch: 表示要更新的日历化价格库存以及开关,增量更新,更新哪天就传那天的日历化单元, 可选字段 invAndSwitch.date: 表示日历化单元里的日期,格式:yyyy-MM-dd, 在 invAndSwitch 里的每个日历化单元里是必填字段 invAndSwitch.price: 表示要更新的价格,单位 分, 可选字段 invAndSwitch.quota: 表示要更新的普通库存,可选字段  invAndSwitch.rateStatus: 表示要更新的价格开关,1表示开,0表示关,可选字段。请注意，该接口为轻量级批量增量更新接口，只能更新未来4天内价格库存开关信息，传入的参数日期超过4天了会报错。该接口核心是解决近几天价格库存实时性。
	_rateListInvInfo string
}

// NewTaobaoXhotelRatesLiteIncrUpdateRequest 初始化TaobaoXhotelRatesLiteIncrUpdateAPIRequest对象
func NewTaobaoXhotelRatesLiteIncrUpdateRequest() *TaobaoXhotelRatesLiteIncrUpdateAPIRequest {
	return &TaobaoXhotelRatesLiteIncrUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRatesLiteIncrUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rates.lite.incr.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRatesLiteIncrUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRatesLiteIncrUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRateListInvInfo is RateListInvInfo Setter
// json格式: [{&#34;rp_code&#34;:&#34;031612_181220000074_F1_1&#34;,&#34;out_rid&#34;:&#34;031612_HD&#34;,&#34;vendor&#34;:&#34;taobao&#34;,&#34;useRoomInv&#34;:0,&#34;invAndSwitch&#34;:[{&#34;date&#34;:&#34;2020-01-09&#34;,&#34;price&#34;:1200,&#34;quota&#34;:23,&#34;rateStatus&#34;:1},{&#34;date&#34;:&#34;2020-01-10&#34;,&#34;price&#34;:1200,&#34;quota&#34;:23,&#34;rateStatus&#34;:0}]},{&#34;rp_code&#34;:&#34;1234&#34;,&#34;out_rid&#34;:&#34;123456&#34;,&#34;vendor&#34;:&#34;taobao&#34;,&#34;useRoomInv&#34;:0,&#34;invAndSwitch&#34;:[{&#34;date&#34;:&#34;yyyy-MM-dd&#34;,&#34;price&#34;:1200,&#34;quota&#34;:23,&#34;rateStatus&#34;:1},{&#34;date&#34;:&#34;yyyy-MM-dd&#34;,&#34;price&#34;:1200,&#34;quota&#34;:23,&#34;rateStatus&#34;:1}]}] rp_code:房价code,out_rid:房型code,必传字段 useRoomInv: 表示是否使用房型库存(1表示是,0表示否),可选字段 invAndSwitch: 表示要更新的日历化价格库存以及开关,增量更新,更新哪天就传那天的日历化单元, 可选字段 invAndSwitch.date: 表示日历化单元里的日期,格式:yyyy-MM-dd, 在 invAndSwitch 里的每个日历化单元里是必填字段 invAndSwitch.price: 表示要更新的价格,单位 分, 可选字段 invAndSwitch.quota: 表示要更新的普通库存,可选字段  invAndSwitch.rateStatus: 表示要更新的价格开关,1表示开,0表示关,可选字段。请注意，该接口为轻量级批量增量更新接口，只能更新未来4天内价格库存开关信息，传入的参数日期超过4天了会报错。该接口核心是解决近几天价格库存实时性。
func (r *TaobaoXhotelRatesLiteIncrUpdateAPIRequest) SetRateListInvInfo(_rateListInvInfo string) error {
	r._rateListInvInfo = _rateListInvInfo
	r.Set("rate_list_inv_info", _rateListInvInfo)
	return nil
}

// GetRateListInvInfo RateListInvInfo Getter
func (r TaobaoXhotelRatesLiteIncrUpdateAPIRequest) GetRateListInvInfo() string {
	return r._rateListInvInfo
}
