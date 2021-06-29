package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
购物车催付优惠券到期查询用户信息 API请求
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

// 初始化TaobaoTbkCartCouponExpireUserQueryRequest对象
func NewTaobaoTbkCartCouponExpireUserQueryRequest() *TaobaoTbkCartCouponExpireUserQueryRequest{
    return &TaobaoTbkCartCouponExpireUserQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetApiMethodName() string {
    return "taobao.tbk.cart.coupon.expire.user.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RuleId Setter
// 规则ID，由接口提供方分配
func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetRuleId(ruleId string) error {
    r.ruleId = ruleId
    r.Set("rule_id", ruleId)
    return nil
}

// RuleId Getter
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetRuleId() string {
    return r.ruleId
}
// PageSize Setter
// 每页大小
func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNum Setter
// 页码，从0开始
func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetPageNum() int64 {
    return r.pageNum
}
