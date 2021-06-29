package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服排班信息查询接口 API请求
taobao.weike.eservice.schedule.get

客服排班信息查询接口
*/
type TaobaoWeikeEserviceScheduleGetRequest struct {
    model.Params
    // 订单ID，orderId、sellerNick、spNick三者不能同时为Null
    orderId   int64
    // 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
    sellerNick   string
    // 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
    spNick   string
    // 起始日期，起始日期和结束日期跨度不能超过31天
    startDate   string
    // 结束日期，起始日期和结束日期跨度不能超过31天
    endDate   string
}

// 初始化TaobaoWeikeEserviceScheduleGetRequest对象
func NewTaobaoWeikeEserviceScheduleGetRequest() *TaobaoWeikeEserviceScheduleGetRequest{
    return &TaobaoWeikeEserviceScheduleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceScheduleGetRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.schedule.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceScheduleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoWeikeEserviceScheduleGetRequest) GetOrderId() int64 {
    return r.orderId
}
// SellerNick Setter
// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoWeikeEserviceScheduleGetRequest) GetSellerNick() string {
    return r.sellerNick
}
// SpNick Setter
// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetRequest) SetSpNick(spNick string) error {
    r.spNick = spNick
    r.Set("sp_nick", spNick)
    return nil
}

// SpNick Getter
func (r TaobaoWeikeEserviceScheduleGetRequest) GetSpNick() string {
    return r.spNick
}
// StartDate Setter
// 起始日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoWeikeEserviceScheduleGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoWeikeEserviceScheduleGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoWeikeEserviceScheduleGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoWeikeEserviceScheduleGetRequest) GetEndDate() string {
    return r.endDate
}
