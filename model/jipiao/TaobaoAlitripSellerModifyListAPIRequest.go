package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerModifyListAPIRequest 【机票代理商订单】改签订单列表 API请求
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
type TaobaoAlitripSellerModifyListAPIRequest struct {
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

// NewTaobaoAlitripSellerModifyListRequest 初始化TaobaoAlitripSellerModifyListAPIRequest对象
func NewTaobaoAlitripSellerModifyListRequest() *TaobaoAlitripSellerModifyListAPIRequest {
	return &TaobaoAlitripSellerModifyListAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripSellerModifyListAPIRequest) Reset() {
	r._applyDateEnd = ""
	r._applyDateStart = ""
	r._flyDateEnd = ""
	r._flyDateStart = ""
	r._modifyDateEnd = ""
	r._modifyDateStart = ""
	r._applyId = 0
	r._currentPage = 0
	r._orderId = 0
	r._pageSize = 0
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerModifyListAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.modify.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerModifyListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerModifyListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyDateEnd is ApplyDateEnd Setter
// 改签发起时间的查询结束日期 和 更新时间必选其一
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetApplyDateEnd(_applyDateEnd string) error {
	r._applyDateEnd = _applyDateEnd
	r.Set("apply_date_end", _applyDateEnd)
	return nil
}

// GetApplyDateEnd ApplyDateEnd Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetApplyDateEnd() string {
	return r._applyDateEnd
}

// SetApplyDateStart is ApplyDateStart Setter
// 改签发起时间的查询开始日期 和 更新时间必选其一
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetApplyDateStart(_applyDateStart string) error {
	r._applyDateStart = _applyDateStart
	r.Set("apply_date_start", _applyDateStart)
	return nil
}

// GetApplyDateStart ApplyDateStart Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetApplyDateStart() string {
	return r._applyDateStart
}

// SetFlyDateEnd is FlyDateEnd Setter
// 乘客起飞时间的查询结束日期
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetFlyDateEnd(_flyDateEnd string) error {
	r._flyDateEnd = _flyDateEnd
	r.Set("fly_date_end", _flyDateEnd)
	return nil
}

// GetFlyDateEnd FlyDateEnd Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetFlyDateEnd() string {
	return r._flyDateEnd
}

// SetFlyDateStart is FlyDateStart Setter
// 乘客起飞时间的查询开始日期
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetFlyDateStart(_flyDateStart string) error {
	r._flyDateStart = _flyDateStart
	r.Set("fly_date_start", _flyDateStart)
	return nil
}

// GetFlyDateStart FlyDateStart Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetFlyDateStart() string {
	return r._flyDateStart
}

// SetModifyDateEnd is ModifyDateEnd Setter
// 记录修改结束时间  和 改签发起时间必选其一
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetModifyDateEnd(_modifyDateEnd string) error {
	r._modifyDateEnd = _modifyDateEnd
	r.Set("modify_date_end", _modifyDateEnd)
	return nil
}

// GetModifyDateEnd ModifyDateEnd Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetModifyDateEnd() string {
	return r._modifyDateEnd
}

// SetModifyDateStart is ModifyDateStart Setter
// 记录修改起始时间 和 改签发起时间必选其一
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetModifyDateStart(_modifyDateStart string) error {
	r._modifyDateStart = _modifyDateStart
	r.Set("modify_date_start", _modifyDateStart)
	return nil
}

// GetModifyDateStart ModifyDateStart Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetModifyDateStart() string {
	return r._modifyDateStart
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetApplyId() int64 {
	return r._applyId
}

// SetCurrentPage is CurrentPage Setter
// 页码
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetOrderId is OrderId Setter
// 淘宝订单号
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetStatus is Status Setter
// 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
func (r *TaobaoAlitripSellerModifyListAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoAlitripSellerModifyListAPIRequest) GetStatus() int64 {
	return r._status
}

var poolTaobaoAlitripSellerModifyListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripSellerModifyListRequest()
	},
}

// GetTaobaoAlitripSellerModifyListRequest 从 sync.Pool 获取 TaobaoAlitripSellerModifyListAPIRequest
func GetTaobaoAlitripSellerModifyListAPIRequest() *TaobaoAlitripSellerModifyListAPIRequest {
	return poolTaobaoAlitripSellerModifyListAPIRequest.Get().(*TaobaoAlitripSellerModifyListAPIRequest)
}

// ReleaseTaobaoAlitripSellerModifyListAPIRequest 将 TaobaoAlitripSellerModifyListAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripSellerModifyListAPIRequest(v *TaobaoAlitripSellerModifyListAPIRequest) {
	v.Reset()
	poolTaobaoAlitripSellerModifyListAPIRequest.Put(v)
}
