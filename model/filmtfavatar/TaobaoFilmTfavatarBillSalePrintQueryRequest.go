package filmtfavatar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院卖品账单-核销账单 API请求
taobao.film.tfavatar.bill.sale.print.query

获取影院卖品账单-核销账单
返回值data属于加密字段, 并非大字段.
*/
type TaobaoFilmTfavatarBillSalePrintQueryRequest struct {
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

// 初始化TaobaoFilmTfavatarBillSalePrintQueryRequest对象
func NewTaobaoFilmTfavatarBillSalePrintQueryRequest() *TaobaoFilmTfavatarBillSalePrintQueryRequest{
    return &TaobaoFilmTfavatarBillSalePrintQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetApiMethodName() string {
    return "taobao.film.tfavatar.bill.sale.print.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAppKey Setter
// 自运营开放平台APPKEY
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

// OpenAppKey Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}
// CinemaId Setter
// 影院ID
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetCinemaId(cinemaId int64) error {
    r.cinemaId = cinemaId
    r.Set("cinema_id", cinemaId)
    return nil
}

// CinemaId Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetCinemaId() int64 {
    return r.cinemaId
}
// BeginTime Setter
// 开始时间
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetBeginTime(beginTime string) error {
    r.beginTime = beginTime
    r.Set("begin_time", beginTime)
    return nil
}

// BeginTime Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetBeginTime() string {
    return r.beginTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetEndTime() string {
    return r.endTime
}
// IncludedOrderStatus Setter
// 包含的订单状态, 默认不填
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetIncludedOrderStatus(includedOrderStatus []string) error {
    r.includedOrderStatus = includedOrderStatus
    r.Set("included_order_status", includedOrderStatus)
    return nil
}

// IncludedOrderStatus Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetIncludedOrderStatus() []string {
    return r.includedOrderStatus
}
// Offset Setter
// offset 下标, 从0开始
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

// Offset Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetOffset() int64 {
    return r.offset
}
// PageSize Setter
// 页大小
func (r *TaobaoFilmTfavatarBillSalePrintQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFilmTfavatarBillSalePrintQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
