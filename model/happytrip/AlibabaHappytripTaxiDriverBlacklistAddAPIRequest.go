package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiDriverBlacklistAddAPIRequest
添加司机黑名单 API请求
alibaba.happytrip.taxi.driver.blacklist.add

实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行 */
type AlibabaHappytripTaxiDriverBlacklistAddAPIRequest struct {
	model.Params
	// 供应商单号
	_orderId string
	// 用户唯一标识
	_uid string
	// 司机唯一标识
	_driverId string
}

// NewAlibabaHappytripTaxiDriverBlacklistAddRequest 初始化AlibabaHappytripTaxiDriverBlacklistAddAPIRequest对象
func NewAlibabaHappytripTaxiDriverBlacklistAddRequest() *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest {
	return &AlibabaHappytripTaxiDriverBlacklistAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.blacklist.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 供应商单号
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// Get Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetUid() string {
	return r._uid
}

// Set is DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) SetDriverId(_driverId string) error {
	r._driverId = _driverId
	r.Set("driver_id", _driverId)
	return nil
}

// Get DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetDriverId() string {
	return r._driverId
}
