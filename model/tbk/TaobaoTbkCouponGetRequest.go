package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-阿里妈妈推广券详情查询 API请求
taobao.tbk.coupon.get

传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
*/
type TaobaoTbkCouponGetRequest struct {
    model.Params
    // 带券ID与商品ID的加密串
    me   string
    // 商品ID
    itemId   int64
    // 券ID
    activityId   string
}

// 初始化TaobaoTbkCouponGetRequest对象
func NewTaobaoTbkCouponGetRequest() *TaobaoTbkCouponGetRequest{
    return &TaobaoTbkCouponGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkCouponGetRequest) GetApiMethodName() string {
    return "taobao.tbk.coupon.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkCouponGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Me Setter
// 带券ID与商品ID的加密串
func (r *TaobaoTbkCouponGetRequest) SetMe(me string) error {
    r.me = me
    r.Set("me", me)
    return nil
}

// Me Getter
func (r TaobaoTbkCouponGetRequest) GetMe() string {
    return r.me
}
// ItemId Setter
// 商品ID
func (r *TaobaoTbkCouponGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoTbkCouponGetRequest) GetItemId() int64 {
    return r.itemId
}
// ActivityId Setter
// 券ID
func (r *TaobaoTbkCouponGetRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoTbkCouponGetRequest) GetActivityId() string {
    return r.activityId
}
