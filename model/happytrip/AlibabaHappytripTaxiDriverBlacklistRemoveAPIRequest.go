package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxidriverblacklistremoveAPIRequest 移除司机黑名单 API请求
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
type AlibabahappytriptaxidriverblacklistremoveAPIRequest struct {
	model.Params
	// 供应商单号
	_orderId string
	// 用户唯一标识
	_uid string
	// 司机唯一标识
	_driverId string
}

// NewAlibabahappytriptaxidriverblacklistremoveRequest 初始化AlibabahappytriptaxidriverblacklistremoveAPIRequest对象
func NewAlibabahappytriptaxidriverblacklistremoveRequest() *AlibabahappytriptaxidriverblacklistremoveAPIRequest {
	return &AlibabahappytriptaxidriverblacklistremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxidriverblacklistremoveAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.blacklist.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxidriverblacklistremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxidriverblacklistremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商单号
func (r *AlibabahappytriptaxidriverblacklistremoveAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxidriverblacklistremoveAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetUid is Uid Setter
// 用户唯一标识
func (r *AlibabahappytriptaxidriverblacklistremoveAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabahappytriptaxidriverblacklistremoveAPIRequest) GetUid() string {
	return r._uid
}

// SetDriverId is DriverId Setter
// 司机唯一标识
func (r *AlibabahappytriptaxidriverblacklistremoveAPIRequest) SetDriverId(_driverId string) error {
	r._driverId = _driverId
	r.Set("driver_id", _driverId)
	return nil
}

// GetDriverId DriverId Getter
func (r AlibabahappytriptaxidriverblacklistremoveAPIRequest) GetDriverId() string {
	return r._driverId
}
