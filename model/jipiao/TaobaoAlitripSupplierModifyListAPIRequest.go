package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】改签通知单列表 API请求
taobao.alitrip.supplier.modify.list

提供供应商查询改签通知单列表
*/
type TaobaoAlitripSupplierModifyListAPIRequest struct {
    model.Params
    // 页码
    _currentPage   int64
    // 乘客出发时间查询结束日期
    _depEnd   string
    // 乘客出发时间查询开始日期
    _depStart   string
    // 申请单创建时间查询结束日期
    _gmtCreateEnd   string
    // 申请单创建时间查询开始日期
    _gmtCreateStart   string
    // 每页记录数
    _pageSize   int64
    // 1：改签申请列表，2：改签已支付列表，3：改签成功列表
    _status   int64
}

// 初始化TaobaoAlitripSupplierModifyListAPIRequest对象
func NewTaobaoAlitripSupplierModifyListRequest() *TaobaoAlitripSupplierModifyListAPIRequest{
    return &TaobaoAlitripSupplierModifyListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.supplier.modify.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// 页码
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// DepEnd Setter
// 乘客出发时间查询结束日期
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetDepEnd(_depEnd string) error {
    r._depEnd = _depEnd
    r.Set("dep_end", _depEnd)
    return nil
}

// DepEnd Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetDepEnd() string {
    return r._depEnd
}
// DepStart Setter
// 乘客出发时间查询开始日期
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetDepStart(_depStart string) error {
    r._depStart = _depStart
    r.Set("dep_start", _depStart)
    return nil
}

// DepStart Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetDepStart() string {
    return r._depStart
}
// GmtCreateEnd Setter
// 申请单创建时间查询结束日期
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
    r._gmtCreateEnd = _gmtCreateEnd
    r.Set("gmt_create_end", _gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetGmtCreateEnd() string {
    return r._gmtCreateEnd
}
// GmtCreateStart Setter
// 申请单创建时间查询开始日期
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
    r._gmtCreateStart = _gmtCreateStart
    r.Set("gmt_create_start", _gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetGmtCreateStart() string {
    return r._gmtCreateStart
}
// PageSize Setter
// 每页记录数
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Status Setter
// 1：改签申请列表，2：改签已支付列表，3：改签成功列表
func (r *TaobaoAlitripSupplierModifyListAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoAlitripSupplierModifyListAPIRequest) GetStatus() int64 {
    return r._status
}
