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
    _activityId   int64
    // 分页页码，页码从1开始
    _pageNo   int64
    // 分页大小，不能超过50
    _pageSize   int64
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
func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetActivityId() int64 {
    return r._activityId
}
// PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页大小，不能超过50
func (r *TaobaoPromotionmiscCommonItemDetailListGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetRequest) GetPageSize() int64 {
    return r._pageSize
}
