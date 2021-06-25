package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改通用单品优惠详情 APIRequest
taobao.promotionmisc.common.item.detail.update

修改通用单品优惠详情。
1、该接口只修改活动下参与的商品的优惠信息，如需要增加、删除活动，请调用taobao.promotionmisc.common.item.activity.add和taobao.promotionmisc.common.item.activity.delete接口；
2、使用该接口时需要把未做修改的字段值也传入；
3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能修改
*/
type TaobaoPromotionmiscCommonItemDetailUpdateRequest struct {
    model.Params

    // 优惠活动ID
    activityId   int64 

    // 优惠详情ID
    detailId   int64 

    // 商品ID
    itemId   int64 

    // 优惠类型，只有两种可选值：0-减钱；1-打折
    promotionType   int64 

    // 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
    promotionValue   int64 

}

func NewTaobaoPromotionmiscCommonItemDetailUpdateRequest() *TaobaoPromotionmiscCommonItemDetailUpdateRequest{
    return &TaobaoPromotionmiscCommonItemDetailUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.detail.update"
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemDetailUpdateRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscCommonItemDetailUpdateRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetDetailId() int64 {
    return r.detailId
}

func (r *TaobaoPromotionmiscCommonItemDetailUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoPromotionmiscCommonItemDetailUpdateRequest) SetPromotionType(promotionType int64) error {
    r.promotionType = promotionType
    r.Set("promotion_type", promotionType)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetPromotionType() int64 {
    return r.promotionType
}

func (r *TaobaoPromotionmiscCommonItemDetailUpdateRequest) SetPromotionValue(promotionValue int64) error {
    r.promotionValue = promotionValue
    r.Set("promotion_value", promotionValue)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailUpdateRequest) GetPromotionValue() int64 {
    return r.promotionValue
}

