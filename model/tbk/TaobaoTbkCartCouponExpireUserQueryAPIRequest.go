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
	_ruleid string
	// 页码，从0开始
	_pagenum int64
	// 每页大小
	_pagesize int64
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

// SetRuleid is Ruleid Setter
// 规则ID，由接口提供方分配
func (r *TaobaotbkcartcouponexpireuserqueryAPIRequest) SetRuleid(_ruleid string) error {
	r._ruleid = _ruleid
	r.Set("rule_id", _ruleid)
	return nil
}

// GetRuleid Ruleid Getter
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetRuleid() string {
	return r._ruleid
}

// SetPagenum is Pagenum Setter
// 页码，从0开始
func (r *TaobaotbkcartcouponexpireuserqueryAPIRequest) SetPagenum(_pagenum int64) error {
	r._pagenum = _pagenum
	r.Set("page_num", _pagenum)
	return nil
}

// GetPagenum Pagenum Getter
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetPagenum() int64 {
	return r._pagenum
}

// SetPagesize is Pagesize Setter
// 每页大小
func (r *TaobaotbkcartcouponexpireuserqueryAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkcartcouponexpireuserqueryAPIRequest) GetPagesize() int64 {
	return r._pagesize
}
