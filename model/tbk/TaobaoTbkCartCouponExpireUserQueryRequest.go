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
    _ruleId   string
    // 每页大小
    _pageSize   int64
    // 页码，从0开始
    _pageNum   int64
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
func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetRuleId(_ruleId string) error {
    r._ruleId = _ruleId
    r.Set("rule_id", _ruleId)
    return nil
}

// RuleId Getter
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetRuleId() string {
    return r._ruleId
}
// PageSize Setter
// 每页大小
func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNum Setter
// 页码，从0开始
func (r *TaobaoTbkCartCouponExpireUserQueryRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoTbkCartCouponExpireUserQueryRequest) GetPageNum() int64 {
    return r._pageNum
}
