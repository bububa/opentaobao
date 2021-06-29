package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菲住联盟分账成功订单同步 APIRequest
alitrip.hotel.alliance.settle.order.syn

用于菲住联盟分账成功订单同步
*/
type AlitripHotelAllianceSettleOrderSynRequest struct {
    model.Params

    // 订单信息
    orderInfo   *AllianceSettleOrderInfo 

}

func NewAlitripHotelAllianceSettleOrderSynRequest() *AlitripHotelAllianceSettleOrderSynRequest{
    return &AlitripHotelAllianceSettleOrderSynRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelAllianceSettleOrderSynRequest) GetApiMethodName() string {
    return "alitrip.hotel.alliance.settle.order.syn"
}

func (r AlitripHotelAllianceSettleOrderSynRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelAllianceSettleOrderSynRequest) SetOrderInfo(orderInfo *AllianceSettleOrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

func (r AlitripHotelAllianceSettleOrderSynRequest) GetOrderInfo() *AllianceSettleOrderInfo {
    return r.orderInfo
}

