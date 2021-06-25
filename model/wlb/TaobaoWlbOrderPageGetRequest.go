package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询物流宝订单 APIRequest
taobao.wlb.order.page.get

分页查询物流宝订单
*/
type TaobaoWlbOrderPageGetRequest struct {
    model.Params

    // 每页多少条
    pageSize   int64 

    // 分页的第几页
    pageNo   int64 

    // TMS拒签：-100 接收方拒签：-200
    orderStatus   int64 

    // 物流订单编号
    orderCode   string 

    // 订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
    orderType   string 

    // 订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
    orderSubType   string 

    // 查询截止时间
    endTime   string 

    // 查询开始时间
    startTime   string 

}

func NewTaobaoWlbOrderPageGetRequest() *TaobaoWlbOrderPageGetRequest{
    return &TaobaoWlbOrderPageGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderPageGetRequest) GetApiMethodName() string {
    return "taobao.wlb.order.page.get"
}

func (r TaobaoWlbOrderPageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderPageGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoWlbOrderPageGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbOrderPageGetRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetOrderStatus() int64 {
    return r.orderStatus
}

func (r *TaobaoWlbOrderPageGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbOrderPageGetRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetOrderType() string {
    return r.orderType
}

func (r *TaobaoWlbOrderPageGetRequest) SetOrderSubType(orderSubType string) error {
    r.orderSubType = orderSubType
    r.Set("order_sub_type", orderSubType)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetOrderSubType() string {
    return r.orderSubType
}

func (r *TaobaoWlbOrderPageGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoWlbOrderPageGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoWlbOrderPageGetRequest) GetStartTime() string {
    return r.startTime
}

