package opentrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeQueueUsersMarkAPIRequest 尖货交易可购买用户标记 API请求
// taobao.opentrade.queue.users.mark
//
// 尖货交易用户标记信息回传，根据openId标记用户可购买商品
type TaobaoOpentradeQueueUsersMarkAPIRequest struct {
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

// NewTaobaoOpentradeQueueUsersMarkRequest 初始化TaobaoOpentradeQueueUsersMarkAPIRequest对象
func NewTaobaoOpentradeQueueUsersMarkRequest() *TaobaoOpentradeQueueUsersMarkAPIRequest {
	return &TaobaoOpentradeQueueUsersMarkAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) Reset() {
	r._openUserIds = r._openUserIds[:0]
	r._status = ""
	r._activityId = ""
	r._skuId = 0
	r._itemId = 0
	r._hit = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.queue.users.mark"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUserIds is OpenUserIds Setter
// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) SetOpenUserIds(_openUserIds []string) error {
	r._openUserIds = _openUserIds
	r.Set("open_user_ids", _openUserIds)
	return nil
}

// GetOpenUserIds OpenUserIds Getter
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetOpenUserIds() []string {
	return r._openUserIds
}

// SetStatus is Status Setter
// 用户状态，可任意传入，后续查询返回
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetStatus() string {
	return r._status
}

// SetActivityId is ActivityId Setter
// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetSkuId is SkuId Setter
// 排队商品SKU ID，不存在传0
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetItemId is ItemId Setter
// 排队商品ID
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetHit is Hit Setter
// 是否目标用户，传入true后，用户可购买商品
func (r *TaobaoOpentradeQueueUsersMarkAPIRequest) SetHit(_hit bool) error {
	r._hit = _hit
	r.Set("hit", _hit)
	return nil
}

// GetHit Hit Getter
func (r TaobaoOpentradeQueueUsersMarkAPIRequest) GetHit() bool {
	return r._hit
}

var poolTaobaoOpentradeQueueUsersMarkAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeQueueUsersMarkRequest()
	},
}

// GetTaobaoOpentradeQueueUsersMarkRequest 从 sync.Pool 获取 TaobaoOpentradeQueueUsersMarkAPIRequest
func GetTaobaoOpentradeQueueUsersMarkAPIRequest() *TaobaoOpentradeQueueUsersMarkAPIRequest {
	return poolTaobaoOpentradeQueueUsersMarkAPIRequest.Get().(*TaobaoOpentradeQueueUsersMarkAPIRequest)
}

// ReleaseTaobaoOpentradeQueueUsersMarkAPIRequest 将 TaobaoOpentradeQueueUsersMarkAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeQueueUsersMarkAPIRequest(v *TaobaoOpentradeQueueUsersMarkAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeQueueUsersMarkAPIRequest.Put(v)
}
