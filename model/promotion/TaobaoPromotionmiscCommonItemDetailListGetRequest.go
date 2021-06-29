package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠详情列表 API请求
taobao.promotionmisc.common.item.detail.list.get

查询通用单品优惠详情列表。
*/
type TaobaoPromotionmiscCommonItemDetailListGetRequest struct {
    model.Params
    // 优惠活动ID
    activityId   int64
    // 分页页码，页码从1开始
    pageNo   int64
    // 分页大小，不能超过50
    pageSize   int64
}

// 初始化TaobaoPromotionmiscCommonItemDetailListGetRequest对象
func NewTaobaoPromotionmiscCommonItemDetailListGetRequest() *TaobaoPromotionmiscCommonItemDetailListGetRequest{
    return &TaobaoPromotionmiscCommonItemDetailListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.detail.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetActivityId() int64 {
    return r.activityId
}
// PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 分页大小，不能超过50
func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetPageSize() int64 {
    return r.pageSize
}
