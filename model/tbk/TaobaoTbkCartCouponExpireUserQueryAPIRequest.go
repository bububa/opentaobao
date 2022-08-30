package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkCartCouponExpireUserQueryAPIRequest 购物车催付优惠券到期查询用户信息 API请求
// taobao.tbk.cart.coupon.expire.user.query
//
// 购物车催付根据对应规则查询用户信息。
type TaobaoTbkCartCouponExpireUserQueryAPIRequest struct {
	model.Params
	// 规则ID，由接口提供方分配
	_ruleId string
	// 页码，从0开始
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoTbkCartCouponExpireUserQueryRequest 初始化TaobaoTbkCartCouponExpireUserQueryAPIRequest对象
func NewTaobaoTbkCartCouponExpireUserQueryRequest() *TaobaoTbkCartCouponExpireUserQueryAPIRequest {
	return &TaobaoTbkCartCouponExpireUserQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkCartCouponExpireUserQueryAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.cart.coupon.expire.user.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkCartCouponExpireUserQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRuleId is RuleId Setter
// 规则ID，由接口提供方分配
func (r *TaobaoTbkCartCouponExpireUserQueryAPIRequest) SetRuleId(_ruleId string) error {
	r._ruleId = _ruleId
	r.Set("rule_id", _ruleId)
	return nil
}

// GetRuleId RuleId Getter
func (r TaobaoTbkCartCouponExpireUserQueryAPIRequest) GetRuleId() string {
	return r._ruleId
}

// SetPageNum is PageNum Setter
// 页码，从0开始
func (r *TaobaoTbkCartCouponExpireUserQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoTbkCartCouponExpireUserQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoTbkCartCouponExpireUserQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkCartCouponExpireUserQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
