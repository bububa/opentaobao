package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询物流宝订单 API请求
taobao.wlb.order.page.get

分页查询物流宝订单
*/
type TaobaoWlbOrderPageGetRequest struct {
    model.Params
    // 每页多少条
    _pageSize   int64
    // 分页的第几页
    _pageNo   int64
    // TMS拒签：-100 接收方拒签：-200
    _orderStatus   int64
    // 物流订单编号
    _orderCode   string
    // 订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
    _orderType   string
    // 订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
    _orderSubType   string
    // 查询截止时间
    _endTime   string
    // 查询开始时间
    _startTime   string
}

// 初始化TaobaoWlbOrderPageGetRequest对象
func NewTaobaoWlbOrderPageGetRequest() *TaobaoWlbOrderPageGetRequest{
    return &TaobaoWlbOrderPageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderPageGetRequest) GetApiMethodName() string {
    return "taobao.wlb.order.page.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderPageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 每页多少条
func (r *TaobaoWlbOrderPageGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbOrderPageGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 分页的第几页
func (r *TaobaoWlbOrderPageGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbOrderPageGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// OrderStatus Setter
// TMS拒签：-100 接收方拒签：-200
func (r *TaobaoWlbOrderPageGetRequest) SetOrderStatus(_orderStatus int64) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r TaobaoWlbOrderPageGetRequest) GetOrderStatus() int64 {
    return r._orderStatus
}
// OrderCode Setter
// 物流订单编号
func (r *TaobaoWlbOrderPageGetRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbOrderPageGetRequest) GetOrderCode() string {
    return r._orderCode
}
// OrderType Setter
// 订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
func (r *TaobaoWlbOrderPageGetRequest) SetOrderType(_orderType string) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbOrderPageGetRequest) GetOrderType() string {
    return r._orderType
}
// OrderSubType Setter
// 订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
func (r *TaobaoWlbOrderPageGetRequest) SetOrderSubType(_orderSubType string) error {
    r._orderSubType = _orderSubType
    r.Set("order_sub_type", _orderSubType)
    return nil
}

// OrderSubType Getter
func (r TaobaoWlbOrderPageGetRequest) GetOrderSubType() string {
    return r._orderSubType
}
// EndTime Setter
// 查询截止时间
func (r *TaobaoWlbOrderPageGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWlbOrderPageGetRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 查询开始时间
func (r *TaobaoWlbOrderPageGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWlbOrderPageGetRequest) GetStartTime() string {
    return r._startTime
}
