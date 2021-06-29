package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院卖品账单--支付账单 APIRequest
taobao.film.tfavatar.bill.sale.payment.query

获取影院卖品账单--支付账单
*/
type TaobaoFilmTfavatarBillSalePaymentQueryRequest struct {
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

func NewTaobaoFilmTfavatarBillSalePaymentQueryRequest() *TaobaoFilmTfavatarBillSalePaymentQueryRequest{
    return &TaobaoFilmTfavatarBillSalePaymentQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.sale.payment.query"
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}

func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetCinemaId(cinemaId int64) error {
    r.cinemaId = cinemaId
    r.Set("cinema_id", cinemaId)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetCinemaId() int64 {
    return r.cinemaId
}

func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetIncludedOrderStatus(includedOrderStatus []string) error {
    r.includedOrderStatus = includedOrderStatus
    r.Set("included_order_status", includedOrderStatus)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetIncludedOrderStatus() []string {
    return r.includedOrderStatus
}

func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetOffset() int64 {
    return r.offset
}

func (r *TaobaoFilmTfavatarBillSalePaymentQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFilmTfavatarBillSalePaymentQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

