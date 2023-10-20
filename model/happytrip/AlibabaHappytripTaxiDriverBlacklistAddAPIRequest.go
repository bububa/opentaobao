package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxidriverblacklistaddAPIRequest 添加司机黑名单 API请求
// alibaba.happytrip.taxi.driver.blacklist.add
//
// 实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行
type AlibabahappytriptaxidriverblacklistaddAPIRequest struct {
	model.Params
	// 供应商单号
	_orderId string
	// 用户唯一标识
	_uid string
	// 司机唯一标识
	_driverId string
}

// NewAlibabahappytriptaxidriverblacklistaddRequest 初始化AlibabahappytriptaxidriverblacklistaddAPIRequest对象
func NewAlibabahappytriptaxidriverblacklistaddRequest() *AlibabahappytriptaxidriverblacklistaddAPIRequest {
	return &AlibabahappytriptaxidriverblacklistaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxidriverblacklistaddAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.blacklist.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxidriverblacklistaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxidriverblacklistaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商单号
func (r *AlibabahappytriptaxidriverblacklistaddAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxidriverblacklistaddAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetUid is Uid Setter
// 用户唯一标识
func (r *AlibabahappytriptaxidriverblacklistaddAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabahappytriptaxidriverblacklistaddAPIRequest) GetUid() string {
	return r._uid
}

// SetDriverId is DriverId Setter
// 司机唯一标识
func (r *AlibabahappytriptaxidriverblacklistaddAPIRequest) SetDriverId(_driverId string) error {
	r._driverId = _driverId
	r.Set("driver_id", _driverId)
	return nil
}

// GetDriverId DriverId Getter
func (r AlibabahappytriptaxidriverblacklistaddAPIRequest) GetDriverId() string {
	return r._driverId
}
