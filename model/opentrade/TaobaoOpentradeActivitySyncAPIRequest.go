package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeActivitySyncAPIRequest 尖货交易活动信息同步 API请求
// taobao.opentrade.activity.sync
//
// 尖货交易活动信息配置，创建或更新活动信息
// 在活动时间开始前，所有用户（包括标记可购买的用户），无法购买商品；
// 在活动时间内，标记可购买的用户可在小程序中跳转下单页，完成购买；
// 在活动结束后，对限购不再限制，平台开放购买，用户可在小程序内、商品详情、购物车下单购买；
type TaobaoOpentradeActivitySyncAPIRequest struct {
	model.Params
	// 活动关联的商品列表，使用逗号(,)分割
	_itemIdList []int64
	// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
	_activityId string
	// 活动开始时间
	_startTime string
	// 活动结束时间（全流程结束时间，非排队结束时间）
	_endTime string
	// 活动名称
	_activityName string
}

// NewTaobaoOpentradeActivitySyncRequest 初始化TaobaoOpentradeActivitySyncAPIRequest对象
func NewTaobaoOpentradeActivitySyncRequest() *TaobaoOpentradeActivitySyncAPIRequest {
	return &TaobaoOpentradeActivitySyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeActivitySyncAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.activity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeActivitySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeActivitySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIdList is ItemIdList Setter
// 活动关联的商品列表，使用逗号(,)分割
func (r *TaobaoOpentradeActivitySyncAPIRequest) SetItemIdList(_itemIdList []int64) error {
	r._itemIdList = _itemIdList
	r.Set("item_id_list", _itemIdList)
	return nil
}

// GetItemIdList ItemIdList Getter
func (r TaobaoOpentradeActivitySyncAPIRequest) GetItemIdList() []int64 {
	return r._itemIdList
}

// SetActivityId is ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoOpentradeActivitySyncAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoOpentradeActivitySyncAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetStartTime is StartTime Setter
// 活动开始时间
func (r *TaobaoOpentradeActivitySyncAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoOpentradeActivitySyncAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间（全流程结束时间，非排队结束时间）
func (r *TaobaoOpentradeActivitySyncAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoOpentradeActivitySyncAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *TaobaoOpentradeActivitySyncAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r TaobaoOpentradeActivitySyncAPIRequest) GetActivityName() string {
	return r._activityName
}
