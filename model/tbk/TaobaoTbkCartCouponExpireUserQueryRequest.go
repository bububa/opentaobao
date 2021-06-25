package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
购物车催付优惠券到期查询用户信息 APIRequest
taobao.tbk.cart.coupon.expire.user.query

购物车催付根据对应规则查询用户信息。
*/
type TaobaoTbkCartCouponExpireUserQueryRequest struct {
    model.Params

    // 规则ID，由接口提供方分配
    ruleId   string 

    // 每页大小
    pageSize   int64 

    // 页码，从0开始
    pageNum   int64 

}

func NewTaobaoTbkCartCouponExpireUserQueryRequest() *TaobaoTbkCartCouponExpireUserQueryRequest{
    return &TaobaoTbkCartCouponExpireUserQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetApiMethodName() string {
    return "taobao.tbk.cart.coupon.expire.user.query"
}

func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetRuleId(ruleId string) error {
    r.ruleId = ruleId
    r.Set("rule_id", ruleId)
    return nil
}

func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetRuleId() string {
    return r.ruleId
}

func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetPageNum() int64 {
    return r.pageNum
}

