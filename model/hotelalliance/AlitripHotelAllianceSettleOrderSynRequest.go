package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菲住联盟分账成功订单同步 API请求
alitrip.hotel.alliance.settle.order.syn

用于菲住联盟分账成功订单同步
*/
type AlitripHotelAllianceSettleOrderSynRequest struct {
    model.Params
    // 订单信息
    orderInfo   *AllianceSettleOrderInfo
}

// 初始化AlitripHotelAllianceSettleOrderSynRequest对象
func NewAlitripHotelAllianceSettleOrderSynRequest() *AlitripHotelAllianceSettleOrderSynRequest{
    return &AlitripHotelAllianceSettleOrderSynRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelAllianceSettleOrderSynRequest) GetApiMethodName() string {
    return "alitrip.hotel.alliance.settle.order.syn"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelAllianceSettleOrderSynRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderInfo Setter
// 订单信息
func (r *AlitripHotelAllianceSettleOrderSynRequest) SetOrderInfo(orderInfo *AllianceSettleOrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlitripHotelAllianceSettleOrderSynRequest) GetOrderInfo() *AllianceSettleOrderInfo {
    return r.orderInfo
}
