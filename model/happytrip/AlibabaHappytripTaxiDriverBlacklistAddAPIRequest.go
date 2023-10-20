package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiDriverBlacklistAddAPIRequest 添加司机黑名单 API请求
// alibaba.happytrip.taxi.driver.blacklist.add
//
// 实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) Reset() {
	r._orderId = ""
	r._uid = ""
	r._driverId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.blacklist.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商单号
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetUid is Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetUid() string {
	return r._uid
}

// SetDriverId is DriverId Setter
// 司机唯一标识
func (r *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) SetDriverId(_driverId string) error {
	r._driverId = _driverId
	r.Set("driver_id", _driverId)
	return nil
}

// GetDriverId DriverId Getter
func (r AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) GetDriverId() string {
	return r._driverId
}

var poolAlibabaHappytripTaxiDriverBlacklistAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiDriverBlacklistAddRequest()
	},
}

// GetAlibabaHappytripTaxiDriverBlacklistAddRequest 从 sync.Pool 获取 AlibabaHappytripTaxiDriverBlacklistAddAPIRequest
func GetAlibabaHappytripTaxiDriverBlacklistAddAPIRequest() *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest {
	return poolAlibabaHappytripTaxiDriverBlacklistAddAPIRequest.Get().(*AlibabaHappytripTaxiDriverBlacklistAddAPIRequest)
}

// ReleaseAlibabaHappytripTaxiDriverBlacklistAddAPIRequest 将 AlibabaHappytripTaxiDriverBlacklistAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiDriverBlacklistAddAPIRequest(v *AlibabaHappytripTaxiDriverBlacklistAddAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiDriverBlacklistAddAPIRequest.Put(v)
}
