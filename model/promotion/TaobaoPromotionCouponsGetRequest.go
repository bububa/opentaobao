package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家优惠券 APIRequest
taobao.promotion.coupons.get

查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量
*/
type TaobaoPromotionCouponsGetRequest struct {
    model.Params

    // 优惠券的id，唯一标识这个优惠券
    couponId   int64 

    // 优惠券的截止日期
    endTime   string 

    // 优惠券的面额，必须是3，5，10，20，50,100
    denominations   int64 

    // 查询的页号，结果集是分页返回的，每页20条
    pageNo   int64 

    // 每页条数
    pageSize   int64 

}

func NewTaobaoPromotionCouponsGetRequest() *TaobaoPromotionCouponsGetRequest{
    return &TaobaoPromotionCouponsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionCouponsGetRequest) GetApiMethodName() string {
    return "taobao.promotion.coupons.get"
}

func (r TaobaoPromotionCouponsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionCouponsGetRequest) SetCouponId(couponId int64) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

func (r TaobaoPromotionCouponsGetRequest) GetCouponId() int64 {
    return r.couponId
}

func (r *TaobaoPromotionCouponsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionCouponsGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoPromotionCouponsGetRequest) SetDenominations(denominations int64) error {
    r.denominations = denominations
    r.Set("denominations", denominations)
    return nil
}

func (r TaobaoPromotionCouponsGetRequest) GetDenominations() int64 {
    return r.denominations
}

func (r *TaobaoPromotionCouponsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoPromotionCouponsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoPromotionCouponsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPromotionCouponsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

