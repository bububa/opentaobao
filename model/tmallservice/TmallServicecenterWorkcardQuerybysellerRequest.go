package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单查询接口（面向商家） API请求
tmall.servicecenter.workcard.querybyseller

查询工单
*/
type TmallServicecenterWorkcardQuerybysellerAPIRequest struct {
    model.Params
    // 门店/网点id
    _serviceStoreId   int64
    // 核销码
    _identifyCode   string
    // 工单id
    _id   int64
    // 工单创建开始时间
    _gmtCreateStart   string
    // 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
    _gmtCreateEnd   string
    // 淘宝交易订单号。主订单或子订单均可
    _bizOrderId   int64
    // 当前页数
    _currentPage   int64
    // 每页大小
    _pageSize   int64
}

// 初始化TmallServicecenterWorkcardQuerybysellerAPIRequest对象
func NewTmallServicecenterWorkcardQuerybysellerRequest() *TmallServicecenterWorkcardQuerybysellerAPIRequest{
    return &TmallServicecenterWorkcardQuerybysellerAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.querybyseller"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStoreId Setter
// 门店/网点id
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetServiceStoreId(_serviceStoreId int64) error {
    r._serviceStoreId = _serviceStoreId
    r.Set("service_store_id", _serviceStoreId)
    return nil
}

// ServiceStoreId Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetServiceStoreId() int64 {
    return r._serviceStoreId
}
// IdentifyCode Setter
// 核销码
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetIdentifyCode(_identifyCode string) error {
    r._identifyCode = _identifyCode
    r.Set("identify_code", _identifyCode)
    return nil
}

// IdentifyCode Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetIdentifyCode() string {
    return r._identifyCode
}
// Id Setter
// 工单id
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetId() int64 {
    return r._id
}
// GmtCreateStart Setter
// 工单创建开始时间
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
    r._gmtCreateStart = _gmtCreateStart
    r.Set("gmt_create_start", _gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetGmtCreateStart() string {
    return r._gmtCreateStart
}
// GmtCreateEnd Setter
// 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
    r._gmtCreateEnd = _gmtCreateEnd
    r.Set("gmt_create_end", _gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetGmtCreateEnd() string {
    return r._gmtCreateEnd
}
// BizOrderId Setter
// 淘宝交易订单号。主订单或子订单均可
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
// CurrentPage Setter
// 当前页数
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页大小
func (r *TmallServicecenterWorkcardQuerybysellerAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallServicecenterWorkcardQuerybysellerAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
