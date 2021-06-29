package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝通知消息查询接口 API请求
taobao.wlb.notify.message.page.get

物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
*/
type TaobaoWlbNotifyMessagePageGetRequest struct {
    model.Params
    // 通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功
    msgCode   string
    // 分页查询页数
    pageNo   int64
    // 分页查询的每页页数
    pageSize   int64
    // 记录开始时间
    startDate   string
    // 记录截至时间
    endDate   string
    // 消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM
    status   string
}

// 初始化TaobaoWlbNotifyMessagePageGetRequest对象
func NewTaobaoWlbNotifyMessagePageGetRequest() *TaobaoWlbNotifyMessagePageGetRequest{
    return &TaobaoWlbNotifyMessagePageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbNotifyMessagePageGetRequest) GetApiMethodName() string {
    return "taobao.wlb.notify.message.page.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbNotifyMessagePageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MsgCode Setter
// 通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功
func (r *TaobaoWlbNotifyMessagePageGetRequest) SetMsgCode(msgCode string) error {
    r.msgCode = msgCode
    r.Set("msg_code", msgCode)
    return nil
}

// MsgCode Getter
func (r TaobaoWlbNotifyMessagePageGetRequest) GetMsgCode() string {
    return r.msgCode
}
// PageNo Setter
// 分页查询页数
func (r *TaobaoWlbNotifyMessagePageGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbNotifyMessagePageGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 分页查询的每页页数
func (r *TaobaoWlbNotifyMessagePageGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbNotifyMessagePageGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// StartDate Setter
// 记录开始时间
func (r *TaobaoWlbNotifyMessagePageGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoWlbNotifyMessagePageGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 记录截至时间
func (r *TaobaoWlbNotifyMessagePageGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoWlbNotifyMessagePageGetRequest) GetEndDate() string {
    return r.endDate
}
// Status Setter
// 消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM
func (r *TaobaoWlbNotifyMessagePageGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoWlbNotifyMessagePageGetRequest) GetStatus() string {
    return r.status
}
