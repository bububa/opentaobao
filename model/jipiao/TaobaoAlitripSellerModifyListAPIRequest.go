package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellermodifylistAPIRequest 【机票代理商订单】改签订单列表 API请求
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
type TaobaoalitripsellermodifylistAPIRequest struct {
	model.Params
	// 改签发起时间的查询结束日期 和 更新时间必选其一
	_applyDateEnd string
	// 改签发起时间的查询开始日期 和 更新时间必选其一
	_applyDateStart string
	// 乘客起飞时间的查询结束日期
	_flyDateEnd string
	// 乘客起飞时间的查询开始日期
	_flyDateStart string
	// 记录修改结束时间  和 改签发起时间必选其一
	_modifyDateEnd string
	// 记录修改起始时间 和 改签发起时间必选其一
	_modifyDateStart string
	// 申请单ID
	_applyId int64
	// 页码
	_currentPage int64
	// 淘宝订单号
	_orderId int64
	// 每页记录数
	_pageSize int64
	// 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
	_status int64
}

// NewTaobaoalitripsellermodifylistRequest 初始化TaobaoalitripsellermodifylistAPIRequest对象
func NewTaobaoalitripsellermodifylistRequest() *TaobaoalitripsellermodifylistAPIRequest {
	return &TaobaoalitripsellermodifylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellermodifylistAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.modify.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellermodifylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellermodifylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyDateEnd is ApplyDateEnd Setter
// 改签发起时间的查询结束日期 和 更新时间必选其一
func (r *TaobaoalitripsellermodifylistAPIRequest) SetApplyDateEnd(_applyDateEnd string) error {
	r._applyDateEnd = _applyDateEnd
	r.Set("apply_date_end", _applyDateEnd)
	return nil
}

// GetApplyDateEnd ApplyDateEnd Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetApplyDateEnd() string {
	return r._applyDateEnd
}

// SetApplyDateStart is ApplyDateStart Setter
// 改签发起时间的查询开始日期 和 更新时间必选其一
func (r *TaobaoalitripsellermodifylistAPIRequest) SetApplyDateStart(_applyDateStart string) error {
	r._applyDateStart = _applyDateStart
	r.Set("apply_date_start", _applyDateStart)
	return nil
}

// GetApplyDateStart ApplyDateStart Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetApplyDateStart() string {
	return r._applyDateStart
}

// SetFlyDateEnd is FlyDateEnd Setter
// 乘客起飞时间的查询结束日期
func (r *TaobaoalitripsellermodifylistAPIRequest) SetFlyDateEnd(_flyDateEnd string) error {
	r._flyDateEnd = _flyDateEnd
	r.Set("fly_date_end", _flyDateEnd)
	return nil
}

// GetFlyDateEnd FlyDateEnd Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetFlyDateEnd() string {
	return r._flyDateEnd
}

// SetFlyDateStart is FlyDateStart Setter
// 乘客起飞时间的查询开始日期
func (r *TaobaoalitripsellermodifylistAPIRequest) SetFlyDateStart(_flyDateStart string) error {
	r._flyDateStart = _flyDateStart
	r.Set("fly_date_start", _flyDateStart)
	return nil
}

// GetFlyDateStart FlyDateStart Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetFlyDateStart() string {
	return r._flyDateStart
}

// SetModifyDateEnd is ModifyDateEnd Setter
// 记录修改结束时间  和 改签发起时间必选其一
func (r *TaobaoalitripsellermodifylistAPIRequest) SetModifyDateEnd(_modifyDateEnd string) error {
	r._modifyDateEnd = _modifyDateEnd
	r.Set("modify_date_end", _modifyDateEnd)
	return nil
}

// GetModifyDateEnd ModifyDateEnd Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetModifyDateEnd() string {
	return r._modifyDateEnd
}

// SetModifyDateStart is ModifyDateStart Setter
// 记录修改起始时间 和 改签发起时间必选其一
func (r *TaobaoalitripsellermodifylistAPIRequest) SetModifyDateStart(_modifyDateStart string) error {
	r._modifyDateStart = _modifyDateStart
	r.Set("modify_date_start", _modifyDateStart)
	return nil
}

// GetModifyDateStart ModifyDateStart Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetModifyDateStart() string {
	return r._modifyDateStart
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoalitripsellermodifylistAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetCurrentPage is CurrentPage Setter
// 页码
func (r *TaobaoalitripsellermodifylistAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetOrderId is OrderId Setter
// 淘宝订单号
func (r *TaobaoalitripsellermodifylistAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *TaobaoalitripsellermodifylistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStatus is Status Setter
// 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
func (r *TaobaoalitripsellermodifylistAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoalitripsellermodifylistAPIRequest) GetStatus() int64 {
	return r._status
}
