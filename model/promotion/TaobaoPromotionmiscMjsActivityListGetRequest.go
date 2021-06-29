package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动列表 API请求
taobao.promotionmisc.mjs.activity.list.get

查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
*/
type TaobaoPromotionmiscMjsActivityListGetRequest struct {
    model.Params
    // 活动类型： 1表示商品级别的活动；2表示店铺级别的活动。
    _activityType   int64
    // 页码。
    _pageNo   int64
    // 每页记录数，最大支持50 。
    _pageSize   int64
}

// 初始化TaobaoPromotionmiscMjsActivityListGetRequest对象
func NewTaobaoPromotionmiscMjsActivityListGetRequest() *TaobaoPromotionmiscMjsActivityListGetRequest{
    return &TaobaoPromotionmiscMjsActivityListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityType Setter
// 活动类型： 1表示商品级别的活动；2表示店铺级别的活动。
func (r *TaobaoPromotionmiscMjsActivityListGetRequest) SetActivityType(_activityType int64) error {
    r._activityType = _activityType
    r.Set("activity_type", _activityType)
    return nil
}

// ActivityType Getter
func (r TaobaoPromotionmiscMjsActivityListGetRequest) GetActivityType() int64 {
    return r._activityType
}
// PageNo Setter
// 页码。
func (r *TaobaoPromotionmiscMjsActivityListGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPromotionmiscMjsActivityListGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数，最大支持50 。
func (r *TaobaoPromotionmiscMjsActivityListGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionmiscMjsActivityListGetRequest) GetPageSize() int64 {
    return r._pageSize
}
