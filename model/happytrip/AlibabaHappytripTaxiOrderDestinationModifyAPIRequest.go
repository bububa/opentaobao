package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderDestinationModifyAPIRequest 修改目的地 API请求
// alibaba.happytrip.taxi.order.destination.modify
//
// 通知ISV修改订单信息
type AlibabaHappytripTaxiOrderDestinationModifyAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 目的地经度
	_tlng string
	// 目的地纬度
	_tlat string
	// 目的地名称(最多50个字)
	_endName string
	// 目的地详细地址(最多100个字)
	_endAddress string
	// 高德POI唯一标识
	_endPoiId string
}

// NewAlibabaHappytripTaxiOrderDestinationModifyRequest 初始化AlibabaHappytripTaxiOrderDestinationModifyAPIRequest对象
func NewAlibabaHappytripTaxiOrderDestinationModifyRequest() *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest {
	return &AlibabaHappytripTaxiOrderDestinationModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.destination.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetTlng is Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) SetTlng(_tlng string) error {
	r._tlng = _tlng
	r.Set("tlng", _tlng)
	return nil
}

// GetTlng Tlng Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetTlng() string {
	return r._tlng
}

// SetTlat is Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) SetTlat(_tlat string) error {
	r._tlat = _tlat
	r.Set("tlat", _tlat)
	return nil
}

// GetTlat Tlat Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetTlat() string {
	return r._tlat
}

// SetEndName is EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) SetEndName(_endName string) error {
	r._endName = _endName
	r.Set("end_name", _endName)
	return nil
}

// GetEndName EndName Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetEndName() string {
	return r._endName
}

// SetEndAddress is EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) SetEndAddress(_endAddress string) error {
	r._endAddress = _endAddress
	r.Set("end_address", _endAddress)
	return nil
}

// GetEndAddress EndAddress Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetEndAddress() string {
	return r._endAddress
}

// SetEndPoiId is EndPoiId Setter
// 高德POI唯一标识
func (r *AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) SetEndPoiId(_endPoiId string) error {
	r._endPoiId = _endPoiId
	r.Set("end_poi_id", _endPoiId)
	return nil
}

// GetEndPoiId EndPoiId Getter
func (r AlibabaHappytripTaxiOrderDestinationModifyAPIRequest) GetEndPoiId() string {
	return r._endPoiId
}
