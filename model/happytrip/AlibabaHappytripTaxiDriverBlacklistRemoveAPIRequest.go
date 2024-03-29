package happytrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) Reset() {
	r._orderId = ""
	r._uid = ""
	r._driverId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.blacklist.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商单号
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetUid is Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetUid() string {
	return r._uid
}

// SetDriverId is DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) SetDriverId(_driverId string) error {
	r._driverId = _driverId
	r.Set("driver_id", _driverId)
	return nil
}

// GetDriverId DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) GetDriverId() string {
	return r._driverId
}

var poolAlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiDriverBlacklistRemoveRequest()
	},
}

// GetAlibabaHappytripTaxiDriverBlacklistRemoveRequest 从 sync.Pool 获取 AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest
func GetAlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest() *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest {
	return poolAlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest.Get().(*AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest)
}

// ReleaseAlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest 将 AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest(v *AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest.Put(v)
}
