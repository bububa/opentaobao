package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiorderdestinationmodifyAPIRequest 修改目的地 API请求
// alibaba.happytrip.taxi.order.destination.modify
//
// 通知ISV修改订单信息
type AlibabahappytriptaxiorderdestinationmodifyAPIRequest struct {
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

// NewAlibabahappytriptaxiorderdestinationmodifyRequest 初始化AlibabahappytriptaxiorderdestinationmodifyAPIRequest对象
func NewAlibabahappytriptaxiorderdestinationmodifyRequest() *AlibabahappytriptaxiorderdestinationmodifyAPIRequest {
	return &AlibabahappytriptaxiorderdestinationmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.destination.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytriptaxiorderdestinationmodifyAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetTlng is Tlng Setter
// 目的地经度
func (r *AlibabahappytriptaxiorderdestinationmodifyAPIRequest) SetTlng(_tlng string) error {
	r._tlng = _tlng
	r.Set("tlng", _tlng)
	return nil
}

// GetTlng Tlng Getter
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetTlng() string {
	return r._tlng
}

// SetTlat is Tlat Setter
// 目的地纬度
func (r *AlibabahappytriptaxiorderdestinationmodifyAPIRequest) SetTlat(_tlat string) error {
	r._tlat = _tlat
	r.Set("tlat", _tlat)
	return nil
}

// GetTlat Tlat Getter
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetTlat() string {
	return r._tlat
}

// SetEndName is EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabahappytriptaxiorderdestinationmodifyAPIRequest) SetEndName(_endName string) error {
	r._endName = _endName
	r.Set("end_name", _endName)
	return nil
}

// GetEndName EndName Getter
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetEndName() string {
	return r._endName
}

// SetEndAddress is EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabahappytriptaxiorderdestinationmodifyAPIRequest) SetEndAddress(_endAddress string) error {
	r._endAddress = _endAddress
	r.Set("end_address", _endAddress)
	return nil
}

// GetEndAddress EndAddress Getter
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetEndAddress() string {
	return r._endAddress
}

// SetEndPoiId is EndPoiId Setter
// 高德POI唯一标识
func (r *AlibabahappytriptaxiorderdestinationmodifyAPIRequest) SetEndPoiId(_endPoiId string) error {
	r._endPoiId = _endPoiId
	r.Set("end_poi_id", _endPoiId)
	return nil
}

// GetEndPoiId EndPoiId Getter
func (r AlibabahappytriptaxiorderdestinationmodifyAPIRequest) GetEndPoiId() string {
	return r._endPoiId
}
