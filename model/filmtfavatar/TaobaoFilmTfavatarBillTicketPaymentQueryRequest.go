package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院票务账单-支付订单 APIRequest
taobao.film.tfavatar.bill.ticket.payment.query

获取影院票务账单-支付订单
*/
type TaobaoFilmTfavatarBillTicketPaymentQueryRequest struct {
    model.Params

    // 自运营开放平台APPKEY
    openAppKey   string 

    // 影院ID
    cinemaId   int64 

    // 开始时间
    beginTime   string 

    // 结束时间
    endTime   string 

    // 包含的订单状态, 默认不填
    includedOrderStatus   []string 

    // offset 下标, 从0开始
    offset   int64 

    // 页大小
    pageSize   int64 

}

func NewTaobaoFilmTfavatarBillTicketPaymentQueryRequest() *TaobaoFilmTfavatarBillTicketPaymentQueryRequest{
    return &TaobaoFilmTfavatarBillTicketPaymentQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.ticket.payment.query"
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}

func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetCinemaId(cinemaId int64) error {
    r.cinemaId = cinemaId
    r.Set("cinema_id", cinemaId)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetCinemaId() int64 {
    return r.cinemaId
}

func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetIncludedOrderStatus(includedOrderStatus []string) error {
    r.includedOrderStatus = includedOrderStatus
    r.Set("included_order_status", includedOrderStatus)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetIncludedOrderStatus() []string {
    return r.includedOrderStatus
}

func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetOffset() int64 {
    return r.offset
}

func (r *TaobaoFilmTfavatarBillTicketPaymentQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketPaymentQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

