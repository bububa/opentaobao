package tvpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAccessDataGetAPIRequest tv支付 API请求
// taobao.tvpay.access.data.get
//
// 在匿名用户支付后尝试为其登陆绑定的淘宝账号
type TaobaoTvpayAccessDataGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 账号客户端版本
	_accountClientVersion string
	// 订单id
	_outOrderNo string
}

// NewTaobaoTvpayAccessDataGetRequest 初始化TaobaoTvpayAccessDataGetAPIRequest对象
func NewTaobaoTvpayAccessDataGetRequest() *TaobaoTvpayAccessDataGetAPIRequest {
	return &TaobaoTvpayAccessDataGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTvpayAccessDataGetAPIRequest) Reset() {
	r._deviceId = ""
	r._from = ""
	r._accountClientVersion = ""
	r._outOrderNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAccessDataGetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.access.data.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAccessDataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTvpayAccessDataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetFrom() string {
	return r._from
}

// SetAccountClientVersion is AccountClientVersion Setter
// 账号客户端版本
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetAccountClientVersion(_accountClientVersion string) error {
	r._accountClientVersion = _accountClientVersion
	r.Set("account_client_version", _accountClientVersion)
	return nil
}

// GetAccountClientVersion AccountClientVersion Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetAccountClientVersion() string {
	return r._accountClientVersion
}

// SetOutOrderNo is OutOrderNo Setter
// 订单id
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetOutOrderNo(_outOrderNo string) error {
	r._outOrderNo = _outOrderNo
	r.Set("out_order_no", _outOrderNo)
	return nil
}

// GetOutOrderNo OutOrderNo Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetOutOrderNo() string {
	return r._outOrderNo
}

var poolTaobaoTvpayAccessDataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTvpayAccessDataGetRequest()
	},
}

// GetTaobaoTvpayAccessDataGetRequest 从 sync.Pool 获取 TaobaoTvpayAccessDataGetAPIRequest
func GetTaobaoTvpayAccessDataGetAPIRequest() *TaobaoTvpayAccessDataGetAPIRequest {
	return poolTaobaoTvpayAccessDataGetAPIRequest.Get().(*TaobaoTvpayAccessDataGetAPIRequest)
}

// ReleaseTaobaoTvpayAccessDataGetAPIRequest 将 TaobaoTvpayAccessDataGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTvpayAccessDataGetAPIRequest(v *TaobaoTvpayAccessDataGetAPIRequest) {
	v.Reset()
	poolTaobaoTvpayAccessDataGetAPIRequest.Put(v)
}
