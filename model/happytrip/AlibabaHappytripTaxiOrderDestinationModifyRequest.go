package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改目的地 API请求
alibaba.happytrip.taxi.order.destination.modify

通知ISV修改订单信息
*/
type AlibabaHappytripTaxiOrderDestinationModifyRequest struct {
    model.Params
    // 订单id
    orderId   string
    // 目的地经度
    tlng   string
    // 目的地纬度
    tlat   string
    // 目的地名称(最多50个字)
    endName   string
    // 目的地详细地址(最多100个字)
    endAddress   string
}

// 初始化AlibabaHappytripTaxiOrderDestinationModifyRequest对象
func NewAlibabaHappytripTaxiOrderDestinationModifyRequest() *AlibabaHappytripTaxiOrderDestinationModifyRequest{
    return &AlibabaHappytripTaxiOrderDestinationModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.destination.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetOrderId() string {
    return r.orderId
}
// Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetTlng(tlng string) error {
    r.tlng = tlng
    r.Set("tlng", tlng)
    return nil
}

// Tlng Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetTlng() string {
    return r.tlng
}
// Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetTlat(tlat string) error {
    r.tlat = tlat
    r.Set("tlat", tlat)
    return nil
}

// Tlat Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetTlat() string {
    return r.tlat
}
// EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetEndName(endName string) error {
    r.endName = endName
    r.Set("end_name", endName)
    return nil
}

// EndName Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetEndName() string {
    return r.endName
}
// EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetEndAddress(endAddress string) error {
    r.endAddress = endAddress
    r.Set("end_address", endAddress)
    return nil
}

// EndAddress Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetEndAddress() string {
    return r.endAddress
}
