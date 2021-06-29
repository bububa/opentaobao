package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院卖品账单--退款账单 APIRequest
taobao.film.tfavatar.bill.sale.refund.query

获取影院卖品账单--退款账单
*/
type TaobaoFilmTfavatarBillSaleRefundQueryRequest struct {
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

func NewTaobaoFilmTfavatarBillSaleRefundQueryRequest() *TaobaoFilmTfavatarBillSaleRefundQueryRequest{
    return &TaobaoFilmTfavatarBillSaleRefundQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.sale.refund.query"
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}

func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetCinemaId(cinemaId int64) error {
    r.cinemaId = cinemaId
    r.Set("cinema_id", cinemaId)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetCinemaId() int64 {
    return r.cinemaId
}

func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetBeginTime() string {
    return r.beginTime
}

func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetIncludedOrderStatus(includedOrderStatus []string) error {
    r.includedOrderStatus = includedOrderStatus
    r.Set("included_order_status", includedOrderStatus)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetIncludedOrderStatus() []string {
    return r.includedOrderStatus
}

func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetOffset() int64 {
    return r.offset
}

func (r *TaobaoFilmTfavatarBillSaleRefundQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFilmTfavatarBillSaleRefundQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

