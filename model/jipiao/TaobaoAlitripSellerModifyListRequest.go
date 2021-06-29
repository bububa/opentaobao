package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】改签订单列表 API请求
taobao.alitrip.seller.modify.list

提供机票代理商查询改签订单列表
*/
type TaobaoAlitripSellerModifyListRequest struct {
    model.Params
    // 申请单ID
    _applyId   int64
    // 淘宝订单号
    _orderId   int64
    // 改签发起时间的查询结束日期 和 更新时间必选其一
    _applyDateEnd   string
    // 改签发起时间的查询开始日期 和 更新时间必选其一
    _applyDateStart   string
    // 页码
    _currentPage   int64
    // 乘客起飞时间的查询结束日期
    _flyDateEnd   string
    // 乘客起飞时间的查询开始日期
    _flyDateStart   string
    // 每页记录数
    _pageSize   int64
    // 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
    _status   int64
    // 记录修改结束时间  和 改签发起时间必选其一
    _modifyDateEnd   string
    // 记录修改起始时间 和 改签发起时间必选其一
    _modifyDateStart   string
}

// 初始化TaobaoAlitripSellerModifyListRequest对象
func NewTaobaoAlitripSellerModifyListRequest() *TaobaoAlitripSellerModifyListRequest{
    return &TaobaoAlitripSellerModifyListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerModifyListRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.modify.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerModifyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerModifyListRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerModifyListRequest) GetApplyId() int64 {
    return r._applyId
}
// OrderId Setter
// 淘宝订单号
func (r *TaobaoAlitripSellerModifyListRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripSellerModifyListRequest) GetOrderId() int64 {
    return r._orderId
}
// ApplyDateEnd Setter
// 改签发起时间的查询结束日期 和 更新时间必选其一
func (r *TaobaoAlitripSellerModifyListRequest) SetApplyDateEnd(_applyDateEnd string) error {
    r._applyDateEnd = _applyDateEnd
    r.Set("apply_date_end", _applyDateEnd)
    return nil
}

// ApplyDateEnd Getter
func (r TaobaoAlitripSellerModifyListRequest) GetApplyDateEnd() string {
    return r._applyDateEnd
}
// ApplyDateStart Setter
// 改签发起时间的查询开始日期 和 更新时间必选其一
func (r *TaobaoAlitripSellerModifyListRequest) SetApplyDateStart(_applyDateStart string) error {
    r._applyDateStart = _applyDateStart
    r.Set("apply_date_start", _applyDateStart)
    return nil
}

// ApplyDateStart Getter
func (r TaobaoAlitripSellerModifyListRequest) GetApplyDateStart() string {
    return r._applyDateStart
}
// CurrentPage Setter
// 页码
func (r *TaobaoAlitripSellerModifyListRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoAlitripSellerModifyListRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// FlyDateEnd Setter
// 乘客起飞时间的查询结束日期
func (r *TaobaoAlitripSellerModifyListRequest) SetFlyDateEnd(_flyDateEnd string) error {
    r._flyDateEnd = _flyDateEnd
    r.Set("fly_date_end", _flyDateEnd)
    return nil
}

// FlyDateEnd Getter
func (r TaobaoAlitripSellerModifyListRequest) GetFlyDateEnd() string {
    return r._flyDateEnd
}
// FlyDateStart Setter
// 乘客起飞时间的查询开始日期
func (r *TaobaoAlitripSellerModifyListRequest) SetFlyDateStart(_flyDateStart string) error {
    r._flyDateStart = _flyDateStart
    r.Set("fly_date_start", _flyDateStart)
    return nil
}

// FlyDateStart Getter
func (r TaobaoAlitripSellerModifyListRequest) GetFlyDateStart() string {
    return r._flyDateStart
}
// PageSize Setter
// 每页记录数
func (r *TaobaoAlitripSellerModifyListRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlitripSellerModifyListRequest) GetPageSize() int64 {
    return r._pageSize
}
// Status Setter
// 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
func (r *TaobaoAlitripSellerModifyListRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoAlitripSellerModifyListRequest) GetStatus() int64 {
    return r._status
}
// ModifyDateEnd Setter
// 记录修改结束时间  和 改签发起时间必选其一
func (r *TaobaoAlitripSellerModifyListRequest) SetModifyDateEnd(_modifyDateEnd string) error {
    r._modifyDateEnd = _modifyDateEnd
    r.Set("modify_date_end", _modifyDateEnd)
    return nil
}

// ModifyDateEnd Getter
func (r TaobaoAlitripSellerModifyListRequest) GetModifyDateEnd() string {
    return r._modifyDateEnd
}
// ModifyDateStart Setter
// 记录修改起始时间 和 改签发起时间必选其一
func (r *TaobaoAlitripSellerModifyListRequest) SetModifyDateStart(_modifyDateStart string) error {
    r._modifyDateStart = _modifyDateStart
    r.Set("modify_date_start", _modifyDateStart)
    return nil
}

// ModifyDateStart Getter
func (r TaobaoAlitripSellerModifyListRequest) GetModifyDateStart() string {
    return r._modifyDateStart
}
