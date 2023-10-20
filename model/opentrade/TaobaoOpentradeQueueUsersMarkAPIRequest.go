package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradequeueusersmarkAPIRequest 尖货交易可购买用户标记 API请求
// taobao.opentrade.queue.users.mark
//
// 尖货交易用户标记信息回传，根据openId标记用户可购买商品
type TaobaoopentradequeueusersmarkAPIRequest struct {
	model.Params
	// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
	_openUserIds []string
	// 用户状态，可任意传入，后续查询返回
	_status string
	// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
	_activityId string
	// 排队商品SKU ID，不存在传0
	_skuId int64
	// 排队商品ID
	_itemId int64
	// 是否目标用户，传入true后，用户可购买商品
	_hit bool
}

// NewTaobaoopentradequeueusersmarkRequest 初始化TaobaoopentradequeueusersmarkAPIRequest对象
func NewTaobaoopentradequeueusersmarkRequest() *TaobaoopentradequeueusersmarkAPIRequest {
	return &TaobaoopentradequeueusersmarkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradequeueusersmarkAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.queue.users.mark"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradequeueusersmarkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradequeueusersmarkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUserIds is OpenUserIds Setter
// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
func (r *TaobaoopentradequeueusersmarkAPIRequest) SetOpenUserIds(_openUserIds []string) error {
	r._openUserIds = _openUserIds
	r.Set("open_user_ids", _openUserIds)
	return nil
}

// GetOpenUserIds OpenUserIds Getter
func (r TaobaoopentradequeueusersmarkAPIRequest) GetOpenUserIds() []string {
	return r._openUserIds
}

// SetStatus is Status Setter
// 用户状态，可任意传入，后续查询返回
func (r *TaobaoopentradequeueusersmarkAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoopentradequeueusersmarkAPIRequest) GetStatus() string {
	return r._status
}

// SetActivityId is ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoopentradequeueusersmarkAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoopentradequeueusersmarkAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetSkuId is SkuId Setter
// 排队商品SKU ID，不存在传0
func (r *TaobaoopentradequeueusersmarkAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoopentradequeueusersmarkAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetItemId is ItemId Setter
// 排队商品ID
func (r *TaobaoopentradequeueusersmarkAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoopentradequeueusersmarkAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetHit is Hit Setter
// 是否目标用户，传入true后，用户可购买商品
func (r *TaobaoopentradequeueusersmarkAPIRequest) SetHit(_hit bool) error {
	r._hit = _hit
	r.Set("hit", _hit)
	return nil
}

// GetHit Hit Getter
func (r TaobaoopentradequeueusersmarkAPIRequest) GetHit() bool {
	return r._hit
}
