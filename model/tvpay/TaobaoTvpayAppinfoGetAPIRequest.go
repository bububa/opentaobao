package tvpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAppinfoGetAPIRequest tv支付获取应用信息 API请求
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
type TaobaoTvpayAppinfoGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 客户端版本号
	_clientVersion string
}

// NewTaobaoTvpayAppinfoGetRequest 初始化TaobaoTvpayAppinfoGetAPIRequest对象
func NewTaobaoTvpayAppinfoGetRequest() *TaobaoTvpayAppinfoGetAPIRequest {
	return &TaobaoTvpayAppinfoGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTvpayAppinfoGetAPIRequest) Reset() {
	r._deviceId = ""
	r._from = ""
	r._clientVersion = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAppinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.appinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAppinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTvpayAppinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetFrom() string {
	return r._from
}

// SetClientVersion is ClientVersion Setter
// 客户端版本号
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetClientVersion(_clientVersion string) error {
	r._clientVersion = _clientVersion
	r.Set("client_version", _clientVersion)
	return nil
}

// GetClientVersion ClientVersion Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetClientVersion() string {
	return r._clientVersion
}

var poolTaobaoTvpayAppinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTvpayAppinfoGetRequest()
	},
}

// GetTaobaoTvpayAppinfoGetRequest 从 sync.Pool 获取 TaobaoTvpayAppinfoGetAPIRequest
func GetTaobaoTvpayAppinfoGetAPIRequest() *TaobaoTvpayAppinfoGetAPIRequest {
	return poolTaobaoTvpayAppinfoGetAPIRequest.Get().(*TaobaoTvpayAppinfoGetAPIRequest)
}

// ReleaseTaobaoTvpayAppinfoGetAPIRequest 将 TaobaoTvpayAppinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTvpayAppinfoGetAPIRequest(v *TaobaoTvpayAppinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoTvpayAppinfoGetAPIRequest.Put(v)
}
