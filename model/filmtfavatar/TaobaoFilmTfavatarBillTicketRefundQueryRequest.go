package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院票务账单-退款账单 API请求
taobao.film.tfavatar.bill.ticket.refund.query

获取影院票务账单-支付订单
data字段为加密字段, 不可分拆.
*/
type TaobaoFilmTfavatarBillTicketRefundQueryRequest struct {
    model.Params
    // 自运营开放平台APPKEY
    _openAppKey   string
    // 影院ID
    _cinemaId   int64
    // 开始时间
    _beginTime   string
    // 结束时间
    _endTime   string
    // 包含的订单状态, 默认不填
    _includedOrderStatus   []string
    // offset 下标, 从0开始
    _offset   int64
    // 页大小
    _pageSize   int64
}

// 初始化TaobaoFilmTfavatarBillTicketRefundQueryRequest对象
func NewTaobaoFilmTfavatarBillTicketRefundQueryRequest() *TaobaoFilmTfavatarBillTicketRefundQueryRequest{
    return &TaobaoFilmTfavatarBillTicketRefundQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.ticket.refund.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetOpenAppKey(_openAppKey string) error {
    r._openAppKey = _openAppKey
    r.Set("open_app_key", _openAppKey)
    return nil
}

// OpenAppKey Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetOpenAppKey() string {
    return r._openAppKey
}
// CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetCinemaId(_cinemaId int64) error {
    r._cinemaId = _cinemaId
    r.Set("cinema_id", _cinemaId)
    return nil
}

// CinemaId Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetCinemaId() int64 {
    return r._cinemaId
}
// BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetBeginTime(_beginTime string) error {
    r._beginTime = _beginTime
    r.Set("begin_time", _beginTime)
    return nil
}

// BeginTime Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetBeginTime() string {
    return r._beginTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetEndTime() string {
    return r._endTime
}
// IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetIncludedOrderStatus(_includedOrderStatus []string) error {
    r._includedOrderStatus = _includedOrderStatus
    r.Set("included_order_status", _includedOrderStatus)
    return nil
}

// IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetIncludedOrderStatus() []string {
    return r._includedOrderStatus
}
// Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetOffset(_offset int64) error {
    r._offset = _offset
    r.Set("offset", _offset)
    return nil
}

// Offset Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetOffset() int64 {
    return r._offset
}
// PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
