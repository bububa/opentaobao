package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠详情列表 APIRequest
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

func NewTaobaoPromotionmiscCommonItemDetailListGetRequest() *TaobaoPromotionmiscCommonItemDetailListGetRequest{
    return &TaobaoPromotionmiscCommonItemDetailListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.detail.list.get"
}

func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetPageSize() int64 {
    return r.pageSize
}

