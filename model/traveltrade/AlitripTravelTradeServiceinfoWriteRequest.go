package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务信息写入接口 APIRequest
alitrip.travel.trade.serviceinfo.write

订单服务信息写入接口
*/
type AlitripTravelTradeServiceinfoWriteRequest struct {
    model.Params

    // 根据模版要求传递对应的订单服务信息，当涉及多值时，用英文分号隔开";"，目前api暂时不支持文件上传，只支持链接；链接必须外网能访问
    tipValue   string 

    // 对应的订单信息
    subTcOrderId   int64 

}

func NewAlitripTravelTradeServiceinfoWriteRequest() *AlitripTravelTradeServiceinfoWriteRequest{
    return &AlitripTravelTradeServiceinfoWriteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelTradeServiceinfoWriteRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.serviceinfo.write"
}

func (r AlitripTravelTradeServiceinfoWriteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelTradeServiceinfoWriteRequest) SetTipValue(tipValue string) error {
    r.tipValue = tipValue
    r.Set("tip_value", tipValue)
    return nil
}

func (r AlitripTravelTradeServiceinfoWriteRequest) GetTipValue() string {
    return r.tipValue
}

func (r *AlitripTravelTradeServiceinfoWriteRequest) SetSubTcOrderId(subTcOrderId int64) error {
    r.subTcOrderId = subTcOrderId
    r.Set("sub_tc_order_id", subTcOrderId)
    return nil
}

func (r AlitripTravelTradeServiceinfoWriteRequest) GetSubTcOrderId() int64 {
    return r.subTcOrderId
}

