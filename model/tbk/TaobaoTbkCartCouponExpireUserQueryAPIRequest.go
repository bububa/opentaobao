package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkcartcouponexpireuserqueryAPIRequest 购物车催付优惠券到期查询用户信息 API请求
// taobao.tbk.cart.coupon.expire.user.query
//
// 购物车催付根据对应规则查询用户信息。
type TaobaotbkcartcouponexpireuserqueryAPIRequest struct {
	model.Params
	// 规则ID，由接口提供方分配
	_ruleId string
	// 页码，从0开始
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// NewTaobaotbkcartcouponexpireuserqueryRequest 初始化TaobaotbkcartcouponexpireuserqueryAPIRequest对象
func NewTaobaotbkcartcouponexpireuserqueryRequest() *TaobaotbkcartcouponexpireuserqueryAPIRequest {
	return &TaobaotbkcartcouponexpireuserqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.cart.coupon.expire.user.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRuleId is RuleId Setter
// 规则ID，由接口提供方分配
func (r *TaobaotbkcartcouponexpireuserqueryAPIRequest) SetRuleId(_ruleId string) error {
	r._ruleId = _ruleId
	r.Set("rule_id", _ruleId)
	return nil
}

// GetRuleId RuleId Getter
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetRuleId() string {
	return r._ruleId
}

// SetPageNum is PageNum Setter
// 页码，从0开始
func (r *TaobaotbkcartcouponexpireuserqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaotbkcartcouponexpireuserqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
