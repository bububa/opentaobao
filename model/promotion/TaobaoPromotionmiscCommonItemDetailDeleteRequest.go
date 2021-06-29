package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除通用单品优惠详情 API请求
taobao.promotionmisc.common.item.detail.delete

删除通用单品优惠详情。
*/
type TaobaoPromotionmiscCommonItemDetailDeleteRequest struct {
    model.Params
    // 优惠活动ID
    activityId   int64
    // 优惠详情ID
    detailId   int64
}

// 初始化TaobaoPromotionmiscCommonItemDetailDeleteRequest对象
func NewTaobaoPromotionmiscCommonItemDetailDeleteRequest() *TaobaoPromotionmiscCommonItemDetailDeleteRequest{
    return &TaobaoPromotionmiscCommonItemDetailDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.detail.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetActivityId() int64 {
    return r.activityId
}
// DetailId Setter
// 优惠详情ID
func (r *TaobaoPromotionmiscCommonItemDetailDeleteRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

// DetailId Getter
func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetDetailId() int64 {
    return r.detailId
}
