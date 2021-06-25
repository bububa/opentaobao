package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建通用单品优惠详情 APIRequest
taobao.promotionmisc.common.item.detail.add

创建通用单品优惠详情。
1、使用此接口在指定的优惠活动下创建参与的商品的优惠信息，如还未创建活动，需要先使用接口taobao.promotionmisc.common.item.activity.add创建优惠活动；
2、同一卖家同一活动下的优惠详情数量限制为150个，超过限制需先调用taobao.promotionmisc.common.item.detail.delete接口删除无用的详情后才可再创建新的优惠详情；
3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能创建
*/
type TaobaoPromotionmiscCommonItemDetailAddRequest struct {
    model.Params

    // 优惠活动ID
    activityId   int64 

    // 商品ID
    itemId   int64 

    // 优惠类型，只有两种可选值：0-减钱；1-打折
    promotionType   int64 

    // 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
    promotionValue   int64 

}

func NewTaobaoPromotionmiscCommonItemDetailAddRequest() *TaobaoPromotionmiscCommonItemDetailAddRequest{
    return &TaobaoPromotionmiscCommonItemDetailAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemDetailAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.detail.add"
}

func (r TaobaoPromotionmiscCommonItemDetailAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemDetailAddRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailAddRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscCommonItemDetailAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailAddRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoPromotionmiscCommonItemDetailAddRequest) SetPromotionType(promotionType int64) error {
    r.promotionType = promotionType
    r.Set("promotion_type", promotionType)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailAddRequest) GetPromotionType() int64 {
    return r.promotionType
}

func (r *TaobaoPromotionmiscCommonItemDetailAddRequest) SetPromotionValue(promotionValue int64) error {
    r.promotionValue = promotionValue
    r.Set("promotion_value", promotionValue)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailAddRequest) GetPromotionValue() int64 {
    return r.promotionValue
}

