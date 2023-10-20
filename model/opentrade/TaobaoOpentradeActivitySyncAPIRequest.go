package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradeactivitysyncAPIRequest 尖货交易活动信息同步 API请求
// taobao.opentrade.activity.sync
//
// 尖货交易活动信息配置，创建或更新活动信息
// 在活动时间开始前，所有用户（包括标记可购买的用户），无法购买商品；
// 在活动时间内，标记可购买的用户可在小程序中跳转下单页，完成购买；
// 在活动结束后，对限购不再限制，平台开放购买，用户可在小程序内、商品详情、购物车下单购买；
type TaobaoopentradeactivitysyncAPIRequest struct {
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

// NewTaobaoopentradeactivitysyncRequest 初始化TaobaoopentradeactivitysyncAPIRequest对象
func NewTaobaoopentradeactivitysyncRequest() *TaobaoopentradeactivitysyncAPIRequest {
	return &TaobaoopentradeactivitysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradeactivitysyncAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.activity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradeactivitysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradeactivitysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIdList is ItemIdList Setter
// 活动关联的商品列表，使用逗号(,)分割
func (r *TaobaoopentradeactivitysyncAPIRequest) SetItemIdList(_itemIdList []int64) error {
	r._itemIdList = _itemIdList
	r.Set("item_id_list", _itemIdList)
	return nil
}

// GetItemIdList ItemIdList Getter
func (r TaobaoopentradeactivitysyncAPIRequest) GetItemIdList() []int64 {
	return r._itemIdList
}

// SetActivityId is ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoopentradeactivitysyncAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoopentradeactivitysyncAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetStartTime is StartTime Setter
// 活动开始时间
func (r *TaobaoopentradeactivitysyncAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoopentradeactivitysyncAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间（全流程结束时间，非排队结束时间）
func (r *TaobaoopentradeactivitysyncAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoopentradeactivitysyncAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *TaobaoopentradeactivitysyncAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r TaobaoopentradeactivitysyncAPIRequest) GetActivityName() string {
	return r._activityName
}
