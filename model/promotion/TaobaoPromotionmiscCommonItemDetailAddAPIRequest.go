package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailAddAPIRequest 创建通用单品优惠详情 API请求
// taobao.promotionmisc.common.item.detail.add
//
// 创建通用单品优惠详情。
// 1、使用此接口在指定的优惠活动下创建参与的商品的优惠信息，如还未创建活动，需要先使用接口taobao.promotionmisc.common.item.activity.add创建优惠活动；
// 2、同一卖家同一活动下的优惠详情数量限制为150个，超过限制需先调用taobao.promotionmisc.common.item.detail.delete接口删除无用的详情后才可再创建新的优惠详情；
// 3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能创建
type TaobaoPromotionmiscCommonItemDetailAddAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 商品ID
	_itemId int64
	// 优惠类型，只有两种可选值：0-减钱；1-打折
	_promotionType int64
	// 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
	_promotionValue int64
}

// NewTaobaoPromotionmiscCommonItemDetailAddRequest 初始化TaobaoPromotionmiscCommonItemDetailAddAPIRequest对象
func NewTaobaoPromotionmiscCommonItemDetailAddRequest() *TaobaoPromotionmiscCommonItemDetailAddAPIRequest {
	return &TaobaoPromotionmiscCommonItemDetailAddAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscCommonItemDetailAddAPIRequest) Reset() {
	r._activityId = 0
	r._itemId = 0
	r._promotionType = 0
	r._promotionValue = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailAddAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoPromotionmiscCommonItemDetailAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPromotionType is PromotionType Setter
// 优惠类型，只有两种可选值：0-减钱；1-打折
func (r *TaobaoPromotionmiscCommonItemDetailAddAPIRequest) SetPromotionType(_promotionType int64) error {
	r._promotionType = _promotionType
	r.Set("promotion_type", _promotionType)
	return nil
}

// GetPromotionType PromotionType Getter
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetPromotionType() int64 {
	return r._promotionType
}

// SetPromotionValue is PromotionValue Setter
// 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
func (r *TaobaoPromotionmiscCommonItemDetailAddAPIRequest) SetPromotionValue(_promotionValue int64) error {
	r._promotionValue = _promotionValue
	r.Set("promotion_value", _promotionValue)
	return nil
}

// GetPromotionValue PromotionValue Getter
func (r TaobaoPromotionmiscCommonItemDetailAddAPIRequest) GetPromotionValue() int64 {
	return r._promotionValue
}

var poolTaobaoPromotionmiscCommonItemDetailAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscCommonItemDetailAddRequest()
	},
}

// GetTaobaoPromotionmiscCommonItemDetailAddRequest 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailAddAPIRequest
func GetTaobaoPromotionmiscCommonItemDetailAddAPIRequest() *TaobaoPromotionmiscCommonItemDetailAddAPIRequest {
	return poolTaobaoPromotionmiscCommonItemDetailAddAPIRequest.Get().(*TaobaoPromotionmiscCommonItemDetailAddAPIRequest)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailAddAPIRequest 将 TaobaoPromotionmiscCommonItemDetailAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailAddAPIRequest(v *TaobaoPromotionmiscCommonItemDetailAddAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailAddAPIRequest.Put(v)
}
