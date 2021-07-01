package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayAccessDataGetAPIRequest
tv支付 API请求
taobao.tvpay.access.data.get

在匿名用户支付后尝试为其登陆绑定的淘宝账号 */
type TaobaoTvpayAccessDataGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 订单id
	_outOrderNo string
	// 账号客户端版本
	_accountClientVersion string
}

// NewTaobaoTvpayAccessDataGetRequest 初始化TaobaoTvpayAccessDataGetAPIRequest对象
func NewTaobaoTvpayAccessDataGetRequest() *TaobaoTvpayAccessDataGetAPIRequest {
	return &TaobaoTvpayAccessDataGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAccessDataGetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.access.data.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAccessDataGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceId Setter
// 设备id
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// Set is From Setter
// 来源
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// Get From Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetFrom() string {
	return r._from
}

// Set is OutOrderNo Setter
// 订单id
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetOutOrderNo(_outOrderNo string) error {
	r._outOrderNo = _outOrderNo
	r.Set("out_order_no", _outOrderNo)
	return nil
}

// Get OutOrderNo Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetOutOrderNo() string {
	return r._outOrderNo
}

// Set is AccountClientVersion Setter
// 账号客户端版本
func (r *TaobaoTvpayAccessDataGetAPIRequest) SetAccountClientVersion(_accountClientVersion string) error {
	r._accountClientVersion = _accountClientVersion
	r.Set("account_client_version", _accountClientVersion)
	return nil
}

// Get AccountClientVersion Getter
func (r TaobaoTvpayAccessDataGetAPIRequest) GetAccountClientVersion() string {
	return r._accountClientVersion
}
