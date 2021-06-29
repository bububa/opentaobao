package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院票务账单-支付订单 API请求
taobao.film.tfavatar.bill.ticket.payment.query

获取影院票务账单-支付订单
*/
type TaobaoFilmTfavatarBillTicketPaymentQueryRequest struct {
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

// 初始化TaobaoFilmTfavatarBillTicketPaymentQueryRequest对象
func NewTaobaoFilmTfavatarBillTicketPaymentQueryRequest() *TaobaoFilmTfavatarBillTicketPaymentQueryRequest{
    return &TaobaoFilmTfavatarBillTicketPaymentQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.ticket.payment.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetOpenAppKey(_openAppKey string) error {
    r._openAppKey = _openAppKey
    r.Set("open_app_key", _openAppKey)
    return nil
}

// OpenAppKey Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetOpenAppKey() string {
    return r._openAppKey
}
// CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetCinemaId(_cinemaId int64) error {
    r._cinemaId = _cinemaId
    r.Set("cinema_id", _cinemaId)
    return nil
}

// CinemaId Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetCinemaId() int64 {
    return r._cinemaId
}
// BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetBeginTime(_beginTime string) error {
    r._beginTime = _beginTime
    r.Set("begin_time", _beginTime)
    return nil
}

// BeginTime Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetBeginTime() string {
    return r._beginTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetEndTime() string {
    return r._endTime
}
// IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetIncludedOrderStatus(_includedOrderStatus []string) error {
    r._includedOrderStatus = _includedOrderStatus
    r.Set("included_order_status", _includedOrderStatus)
    return nil
}

// IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetIncludedOrderStatus() []string {
    return r._includedOrderStatus
}
// Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetOffset(_offset int64) error {
    r._offset = _offset
    r.Set("offset", _offset)
    return nil
}

// Offset Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetOffset() int64 {
    return r._offset
}
// PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetPageSize() int64 {
    return r._pageSize
}