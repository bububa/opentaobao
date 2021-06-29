package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单查询 API请求
alibaba.servicecenter.spserviceorder.query

服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
*/
type AlibabaServicecenterSpserviceorderQueryRequest struct {
    model.Params
    // 状态码，可传多个
    _statusCodes   string
    // 查询第几页，默认1
    _currentPage   int64
    // 每页大小，默认50，最大50
    _pageSize   int64
    // 服务单修改时间(时间段15分钟以内)(包含时分秒)
    _gmtModifiedEnd   string
    // 服务单修改时间(包含时分秒)
    _gmtModifiedStart   string
    // 实物主订单id(消费者在淘宝订单里看到的订单号)
    _masterParentBizOrderId   int64
    // 服务单id
    _spServiceOrderId   int64
}

// 初始化AlibabaServicecenterSpserviceorderQueryRequest对象
func NewAlibabaServicecenterSpserviceorderQueryRequest() *AlibabaServicecenterSpserviceorderQueryRequest{
    return &AlibabaServicecenterSpserviceorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.spserviceorder.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StatusCodes Setter
// 状态码，可传多个
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetStatusCodes(_statusCodes string) error {
    r._statusCodes = _statusCodes
    r.Set("status_codes", _statusCodes)
    return nil
}

// StatusCodes Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetStatusCodes() string {
    return r._statusCodes
}
// CurrentPage Setter
// 查询第几页，默认1
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页大小，默认50，最大50
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// GmtModifiedEnd Setter
// 服务单修改时间(时间段15分钟以内)(包含时分秒)
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetGmtModifiedEnd(_gmtModifiedEnd string) error {
    r._gmtModifiedEnd = _gmtModifiedEnd
    r.Set("gmt_modified_end", _gmtModifiedEnd)
    return nil
}

// GmtModifiedEnd Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetGmtModifiedEnd() string {
    return r._gmtModifiedEnd
}
// GmtModifiedStart Setter
// 服务单修改时间(包含时分秒)
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetGmtModifiedStart(_gmtModifiedStart string) error {
    r._gmtModifiedStart = _gmtModifiedStart
    r.Set("gmt_modified_start", _gmtModifiedStart)
    return nil
}

// GmtModifiedStart Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetGmtModifiedStart() string {
    return r._gmtModifiedStart
}
// MasterParentBizOrderId Setter
// 实物主订单id(消费者在淘宝订单里看到的订单号)
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetMasterParentBizOrderId(_masterParentBizOrderId int64) error {
    r._masterParentBizOrderId = _masterParentBizOrderId
    r.Set("master_parent_biz_order_id", _masterParentBizOrderId)
    return nil
}

// MasterParentBizOrderId Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetMasterParentBizOrderId() int64 {
    return r._masterParentBizOrderId
}
// SpServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
    r._spServiceOrderId = _spServiceOrderId
    r.Set("sp_service_order_id", _spServiceOrderId)
    return nil
}

// SpServiceOrderId Getter
func (r AlibabaServicecenterSpserviceorderQueryRequest) GetSpServiceOrderId() int64 {
    return r._spServiceOrderId
}
