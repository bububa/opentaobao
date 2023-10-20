package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoyphordergetAPIRequest 一盘货商家单个查询采购单信息 API请求
// taobao.fenxiao.yph.order.get
//
// 一盘货商家单个查询采购单信息
type TaobaofenxiaoyphordergetAPIRequest struct {
	model.Params
	// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
	_userRoleType int64
	// 采购单id，单个查询此参数必填
	_purchaseOrderId int64
}

// NewTaobaofenxiaoyphordergetRequest 初始化TaobaofenxiaoyphordergetAPIRequest对象
func NewTaobaofenxiaoyphordergetRequest() *TaobaofenxiaoyphordergetAPIRequest {
	return &TaobaofenxiaoyphordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoyphordergetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.yph.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoyphordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoyphordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserRoleType is UserRoleType Setter
// 当前查询用户的角色 当不指定时，默认为供应商 供应商：2，分销商：1
func (r *TaobaofenxiaoyphordergetAPIRequest) SetUserRoleType(_userRoleType int64) error {
	r._userRoleType = _userRoleType
	r.Set("user_role_type", _userRoleType)
	return nil
}

// GetUserRoleType UserRoleType Getter
func (r TaobaofenxiaoyphordergetAPIRequest) GetUserRoleType() int64 {
	return r._userRoleType
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单id，单个查询此参数必填
func (r *TaobaofenxiaoyphordergetAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaofenxiaoyphordergetAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}
