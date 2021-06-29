package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院票务账单-退款账单 APIRequest
taobao.film.tfavatar.bill.ticket.refund.query

获取影院票务账单-支付订单
data字段为加密字段, 不可分拆.
*/
type TaobaoFilmTfavatarBillTicketRefundQueryRequest struct {
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

func NewTaobaoFilmTfavatarBillTicketRefundQueryRequest() *TaobaoFilmTfavatarBillTicketRefundQueryRequest{
    return &TaobaoFilmTfavatarBillTicketRefundQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.ticket.refund.query"
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}

func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetCinemaId(cinemaId int64) error {
    r.cinemaId = cinemaId
    r.Set("cinema_id", cinemaId)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetCinemaId() int64 {
    return r.cinemaId
}

func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetIncludedOrderStatus(includedOrderStatus []string) error {
    r.includedOrderStatus = includedOrderStatus
    r.Set("included_order_status", includedOrderStatus)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetIncludedOrderStatus() []string {
    return r.includedOrderStatus
}

func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetOffset() int64 {
    return r.offset
}

func (r *TaobaoFilmTfavatarBillTicketRefundQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFilmTfavatarBillTicketRefundQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

