package tvpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayOrderPartnerpayAPIRequest tv支付第三方支付订单 API请求
// taobao.tvpay.order.partnerpay
//
// tv支付第三方发起并支付订单（使用设备授权）
type TaobaoTvpayOrderPartnerpayAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 牌照方
	_license string
	// 订单信息
	_data string
	// 支付方式
	_payType string
}

// NewTaobaoTvpayOrderPartnerpayRequest 初始化TaobaoTvpayOrderPartnerpayAPIRequest对象
func NewTaobaoTvpayOrderPartnerpayRequest() *TaobaoTvpayOrderPartnerpayAPIRequest {
	return &TaobaoTvpayOrderPartnerpayAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTvpayOrderPartnerpayAPIRequest) Reset() {
	r._deviceId = ""
	r._from = ""
	r._license = ""
	r._data = ""
	r._payType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.order.partnerpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoTvpayOrderPartnerpayAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaoTvpayOrderPartnerpayAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetFrom() string {
	return r._from
}

// SetLicense is License Setter
// 牌照方
func (r *TaobaoTvpayOrderPartnerpayAPIRequest) SetLicense(_license string) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetLicense() string {
	return r._license
}

// SetData is Data Setter
// 订单信息
func (r *TaobaoTvpayOrderPartnerpayAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetData() string {
	return r._data
}

// SetPayType is PayType Setter
// 支付方式
func (r *TaobaoTvpayOrderPartnerpayAPIRequest) SetPayType(_payType string) error {
	r._payType = _payType
	r.Set("pay_type", _payType)
	return nil
}

// GetPayType PayType Getter
func (r TaobaoTvpayOrderPartnerpayAPIRequest) GetPayType() string {
	return r._payType
}

var poolTaobaoTvpayOrderPartnerpayAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTvpayOrderPartnerpayRequest()
	},
}

// GetTaobaoTvpayOrderPartnerpayRequest 从 sync.Pool 获取 TaobaoTvpayOrderPartnerpayAPIRequest
func GetTaobaoTvpayOrderPartnerpayAPIRequest() *TaobaoTvpayOrderPartnerpayAPIRequest {
	return poolTaobaoTvpayOrderPartnerpayAPIRequest.Get().(*TaobaoTvpayOrderPartnerpayAPIRequest)
}

// ReleaseTaobaoTvpayOrderPartnerpayAPIRequest 将 TaobaoTvpayOrderPartnerpayAPIRequest 放入 sync.Pool
func ReleaseTaobaoTvpayOrderPartnerpayAPIRequest(v *TaobaoTvpayOrderPartnerpayAPIRequest) {
	v.Reset()
	poolTaobaoTvpayOrderPartnerpayAPIRequest.Put(v)
}
