package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服排班信息查询接口 APIRequest
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

func NewTaobaoWeikeEserviceScheduleGetRequest() *TaobaoWeikeEserviceScheduleGetRequest{
    return &TaobaoWeikeEserviceScheduleGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.schedule.get"
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeikeEserviceScheduleGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoWeikeEserviceScheduleGetRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetSellerNick() string {
    return r.sellerNick
}

func (r *TaobaoWeikeEserviceScheduleGetRequest) SetSpNick(spNick string) error {
    r.spNick = spNick
    r.Set("sp_nick", spNick)
    return nil
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetSpNick() string {
    return r.spNick
}

func (r *TaobaoWeikeEserviceScheduleGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoWeikeEserviceScheduleGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoWeikeEserviceScheduleGetRequest) GetEndDate() string {
    return r.endDate
}

