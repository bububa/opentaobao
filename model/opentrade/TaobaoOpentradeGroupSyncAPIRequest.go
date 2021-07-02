package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupSyncAPIRequest 组团购场景创建或更新组团活动 API请求
// taobao.opentrade.group.sync
//
// 组团购场景中创建团购活动
type TaobaoOpentradeGroupSyncAPIRequest struct {
	model.Params
	// 组团活动开始时间
	_startTime string
	// 组团活动结束时间
	_endTime string
	// 成团有效期，单位为妙
	_expiration int64
	// 成团的目标人数
	_goal int64
	// 组团类型，0：拼团；1：团购
	_groupType int64
	// 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
	_allowType string
	// 允许开团的淘宝账号列表
	_allowWhiteList []string
	// 组团类型为团购，可限制团长针对一个商品的开团数量上限
	_openLimit int64
	// 未成团处理办法，close：系统关单；continue：订单继续下行
	_failProcess string
	// 组团购买的折扣价，单位为分
	_discountPrice int64
	// 商品ID
	_itemId int64
	// 组团活动id
	_groupActivityId int64
}

// NewTaobaoOpentradeGroupSyncRequest 初始化TaobaoOpentradeGroupSyncAPIRequest对象
func NewTaobaoOpentradeGroupSyncRequest() *TaobaoOpentradeGroupSyncAPIRequest {
	return &TaobaoOpentradeGroupSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupSyncAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartTime Setter
// 组团活动开始时间
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 组团活动结束时间
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is Expiration Setter
// 成团有效期，单位为妙
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetExpiration(_expiration int64) error {
	r._expiration = _expiration
	r.Set("expiration", _expiration)
	return nil
}

// Get Expiration Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetExpiration() int64 {
	return r._expiration
}

// Set is Goal Setter
// 成团的目标人数
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetGoal(_goal int64) error {
	r._goal = _goal
	r.Set("goal", _goal)
	return nil
}

// Get Goal Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetGoal() int64 {
	return r._goal
}

// Set is GroupType Setter
// 组团类型，0：拼团；1：团购
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetGroupType(_groupType int64) error {
	r._groupType = _groupType
	r.Set("group_type", _groupType)
	return nil
}

// Get GroupType Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetGroupType() int64 {
	return r._groupType
}

// Set is AllowType Setter
// 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetAllowType(_allowType string) error {
	r._allowType = _allowType
	r.Set("allow_type", _allowType)
	return nil
}

// Get AllowType Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetAllowType() string {
	return r._allowType
}

// Set is AllowWhiteList Setter
// 允许开团的淘宝账号列表
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetAllowWhiteList(_allowWhiteList []string) error {
	r._allowWhiteList = _allowWhiteList
	r.Set("allow_white_list", _allowWhiteList)
	return nil
}

// Get AllowWhiteList Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetAllowWhiteList() []string {
	return r._allowWhiteList
}

// Set is OpenLimit Setter
// 组团类型为团购，可限制团长针对一个商品的开团数量上限
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetOpenLimit(_openLimit int64) error {
	r._openLimit = _openLimit
	r.Set("open_limit", _openLimit)
	return nil
}

// Get OpenLimit Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetOpenLimit() int64 {
	return r._openLimit
}

// Set is FailProcess Setter
// 未成团处理办法，close：系统关单；continue：订单继续下行
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetFailProcess(_failProcess string) error {
	r._failProcess = _failProcess
	r.Set("fail_process", _failProcess)
	return nil
}

// Get FailProcess Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetFailProcess() string {
	return r._failProcess
}

// Set is DiscountPrice Setter
// 组团购买的折扣价，单位为分
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetDiscountPrice(_discountPrice int64) error {
	r._discountPrice = _discountPrice
	r.Set("discount_price", _discountPrice)
	return nil
}

// Get DiscountPrice Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetDiscountPrice() int64 {
	return r._discountPrice
}

// Set is ItemId Setter
// 商品ID
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupSyncAPIRequest) SetGroupActivityId(_groupActivityId int64) error {
	r._groupActivityId = _groupActivityId
	r.Set("group_activity_id", _groupActivityId)
	return nil
}

// Get GroupActivityId Getter
func (r TaobaoOpentradeGroupSyncAPIRequest) GetGroupActivityId() int64 {
	return r._groupActivityId
}
