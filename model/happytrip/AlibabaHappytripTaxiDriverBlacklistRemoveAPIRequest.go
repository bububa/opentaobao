package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest 移除司机黑名单 API请求
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
type AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest struct {
	model.Params
	// 供应商单号
	_orderId string
	// 用户唯一标识
	_uid string
	// 司机唯一标识
	_driverId string
}

// NewAlibabaHappytripTaxiDriverBlacklistRemoveRequest 初始化AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest对象
func NewAlibabaHappytripTaxiDriverBlacklistRemoveRequest() *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest {
	return &AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.blacklist.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 供应商单号
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// Get Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetUid() string {
	return r._uid
}

// Set is DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) SetDriverId(_driverId string) error {
	r._driverId = _driverId
	r.Set("driver_id", _driverId)
	return nil
}

// Get DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetDriverId() string {
	return r._driverId
}
