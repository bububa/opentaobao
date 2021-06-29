package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账单信息 APIRequest
taobao.xhotel.order.statement.get

阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
*/
type TaobaoXhotelOrderStatementGetRequest struct {
    model.Params

    // 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
    orderTids   string 

    // 查询条数，最大支持500条
    pageSize   int64 

    // 数据查询开始下标
    start   int64 

    // 0：check_in, 1：check_out,2：分账时间
    dateType   int64 

    // 查询结束时间
    to   string 

    // 查询开始时间
    from   string 

    // 淘宝订单号
    tid   int64 

    // 外部酒店编码
    hotelCode   string 

    // 系统商vendor
    vendor   string 

}

func NewTaobaoXhotelOrderStatementGetRequest() *TaobaoXhotelOrderStatementGetRequest{
    return &TaobaoXhotelOrderStatementGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderStatementGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.statement.get"
}

func (r TaobaoXhotelOrderStatementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderStatementGetRequest) SetOrderTids(orderTids string) error {
    r.orderTids = orderTids
    r.Set("order_tids", orderTids)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetOrderTids() string {
    return r.orderTids
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetStart() int64 {
    return r.start
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetDateType(dateType int64) error {
    r.dateType = dateType
    r.Set("date_type", dateType)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetDateType() int64 {
    return r.dateType
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetTo(to string) error {
    r.to = to
    r.Set("to", to)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetTo() string {
    return r.to
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelOrderStatementGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelOrderStatementGetRequest) GetVendor() string {
    return r.vendor
}

