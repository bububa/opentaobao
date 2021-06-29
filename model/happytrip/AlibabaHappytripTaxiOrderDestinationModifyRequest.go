package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改目的地 APIRequest
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

func NewAlibabaHappytripTaxiOrderDestinationModifyRequest() *AlibabaHappytripTaxiOrderDestinationModifyRequest{
    return &AlibabaHappytripTaxiOrderDestinationModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.destination.modify"
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetTlng(tlng string) error {
    r.tlng = tlng
    r.Set("tlng", tlng)
    return nil
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetTlng() string {
    return r.tlng
}

func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetTlat(tlat string) error {
    r.tlat = tlat
    r.Set("tlat", tlat)
    return nil
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetTlat() string {
    return r.tlat
}

func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetEndName(endName string) error {
    r.endName = endName
    r.Set("end_name", endName)
    return nil
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetEndName() string {
    return r.endName
}

func (r *AlibabaHappytripTaxiOrderDestinationModifyRequest) SetEndAddress(endAddress string) error {
    r.endAddress = endAddress
    r.Set("end_address", endAddress)
    return nil
}

func (r AlibabaHappytripTaxiOrderDestinationModifyRequest) GetEndAddress() string {
    return r.endAddress
}

