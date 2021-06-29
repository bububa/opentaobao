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
    _orderId   string
    // 目的地经度
    _tlng   string
    // 目的地纬度
    _tlat   string
    // 目的地名称(最多50个字)
    _endName   string
    // 目的地详细地址(最多100个字)
    _endAddress   string
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
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetOrderId() string {
    return r._orderId
}
// Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetTlng(_tlng string) error {
    r._tlng = _tlng
    r.Set("tlng", _tlng)
    return nil
}

// Tlng Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetTlng() string {
    return r._tlng
}
// Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetTlat(_tlat string) error {
    r._tlat = _tlat
    r.Set("tlat", _tlat)
    return nil
}

// Tlat Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetTlat() string {
    return r._tlat
}
// EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetEndName(_endName string) error {
    r._endName = _endName
    r.Set("end_name", _endName)
    return nil
}

// EndName Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetEndName() string {
    return r._endName
}
// EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetEndAddress(_endAddress string) error {
    r._endAddress = _endAddress
    r.Set("end_address", _endAddress)
    return nil
}

// EndAddress Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetEndAddress() string {
    return r._endAddress
}
