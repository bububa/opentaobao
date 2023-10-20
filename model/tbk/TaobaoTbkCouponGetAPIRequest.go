package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkcoupongetAPIRequest 淘宝客-公用-阿里妈妈推广券详情查询 API请求
// taobao.tbk.coupon.get
//
// 传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
type TaobaotbkcoupongetAPIRequest struct {
	model.Params
	// 带券ID与商品ID的加密串
	_me string
	// 商品ID
	_itemid string
	// 券ID
	_activityid string
}

// NewTaobaotbkcoupongetRequest 初始化TaobaotbkcoupongetAPIRequest对象
func NewTaobaotbkcoupongetRequest() *TaobaotbkcoupongetAPIRequest {
	return &TaobaotbkcoupongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkcoupongetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkcoupongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkcoupongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMe is Me Setter
// 带券ID与商品ID的加密串
func (r *TaobaotbkcoupongetAPIRequest) SetMe(_me string) error {
	r._me = _me
	r.Set("me", _me)
	return nil
}

// GetMe Me Getter
func (r TaobaotbkcoupongetAPIRequest) GetMe() string {
	return r._me
}

// SetItemid is Itemid Setter
// 商品ID
func (r *TaobaotbkcoupongetAPIRequest) SetItemid(_itemid string) error {
	r._itemid = _itemid
	r.Set("item_id", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaotbkcoupongetAPIRequest) GetItemid() string {
	return r._itemid
}

// SetActivityid is Activityid Setter
// 券ID
func (r *TaobaotbkcoupongetAPIRequest) SetActivityid(_activityid string) error {
	r._activityid = _activityid
	r.Set("activity_id", _activityid)
	return nil
}

// GetActivityid Activityid Getter
func (r TaobaotbkcoupongetAPIRequest) GetActivityid() string {
	return r._activityid
}
