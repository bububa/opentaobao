package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest 修改通用单品优惠详情 API请求
// taobao.promotionmisc.common.item.detail.update
//
// 修改通用单品优惠详情。
// 1、该接口只修改活动下参与的商品的优惠信息，如需要增加、删除活动，请调用taobao.promotionmisc.common.item.activity.add和taobao.promotionmisc.common.item.activity.delete接口；
// 2、使用该接口时需要把未做修改的字段值也传入；
// 3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能修改
type TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 优惠详情ID
	_detailId int64
	// 商品ID
	_itemId int64
	// 优惠类型，只有两种可选值：0-减钱；1-打折
	_promotionType int64
	// 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
	_promotionValue int64
}

// NewTaobaoPromotionmiscCommonItemDetailUpdateRequest 初始化TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest对象
func NewTaobaoPromotionmiscCommonItemDetailUpdateRequest() *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest {
	return &TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) Reset() {
	r._activityId = 0
	r._detailId = 0
	r._itemId = 0
	r._promotionType = 0
	r._promotionValue = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetDetailId is DetailId Setter
// 优惠详情ID
func (r *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetDetailId() int64 {
	return r._detailId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPromotionType is PromotionType Setter
// 优惠类型，只有两种可选值：0-减钱；1-打折
func (r *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) SetPromotionType(_promotionType int64) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetPromotionType() int64 {
	return r._promotionType
}

// SetPromotionValue is PromotionValue Setter
// 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
func (r *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) SetPromotionValue(_promotionValue int64) error {
	r._promotionValue = _promotionValue
	r.Set("promotion_value", _promotionValue)
	return nil
}

// GetPromotionValue PromotionValue Getter
func (r TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) GetPromotionValue() int64 {
	return r._promotionValue
}

var poolTaobaoPromotionmiscCommonItemDetailUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscCommonItemDetailUpdateRequest()
	},
}

// GetTaobaoPromotionmiscCommonItemDetailUpdateRequest 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest
func GetTaobaoPromotionmiscCommonItemDetailUpdateAPIRequest() *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest {
	return poolTaobaoPromotionmiscCommonItemDetailUpdateAPIRequest.Get().(*TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailUpdateAPIRequest 将 TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailUpdateAPIRequest(v *TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailUpdateAPIRequest.Put(v)
}
