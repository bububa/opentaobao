package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加活动参与的商品 API请求
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
*/
type TaobaoPromotionmiscActivityRangeAddRequest struct {
    model.Params
    // 活动id。
    _activityId   int64
    // 商品id,多个id用逗号隔开，一次不超过50个。
    _ids   string
}

// 初始化TaobaoPromotionmiscActivityRangeAddRequest对象
func NewTaobaoPromotionmiscActivityRangeAddRequest() *TaobaoPromotionmiscActivityRangeAddRequest{
    return &TaobaoPromotionmiscActivityRangeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeAddRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeAddRequest) GetActivityId() int64 {
    return r._activityId
}
// Ids Setter
// 商品id,多个id用逗号隔开，一次不超过50个。
func (r *TaobaoPromotionmiscActivityRangeAddRequest) SetIds(_ids string) error {
    r._ids = _ids
    r.Set("ids", _ids)
    return nil
}

// Ids Getter
func (r TaobaoPromotionmiscActivityRangeAddRequest) GetIds() string {
    return r._ids
}
