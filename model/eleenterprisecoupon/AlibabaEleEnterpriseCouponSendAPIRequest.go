package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterprisecouponsendAPIRequest 发放优惠券 API请求
// alibaba.ele.enterprise.coupon.send
//
// 发放优惠券
type AlibabaeleenterprisecouponsendAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 固定值，申请获取
	_channel string
	// 纬度
	_latitude string
	// 经度
	_longitude string
	// 客户端IP地址
	_ip string
	// 客户端User-Agent信息
	_userAgent string
	// 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
	_batchNo string
	// 设备ID
	_deviceId string
}

// NewAlibabaeleenterprisecouponsendRequest 初始化AlibabaeleenterprisecouponsendAPIRequest对象
func NewAlibabaeleenterprisecouponsendRequest() *AlibabaeleenterprisecouponsendAPIRequest {
	return &AlibabaeleenterprisecouponsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterprisecouponsendAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterprisecouponsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterprisecouponsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetPhone() string {
	return r._phone
}

// SetChannel is Channel Setter
// 固定值，申请获取
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetChannel() string {
	return r._channel
}

// SetLatitude is Latitude Setter
// 纬度
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 经度
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetIp is Ip Setter
// 客户端IP地址
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetIp() string {
	return r._ip
}

// SetUserAgent is UserAgent Setter
// 客户端User-Agent信息
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetUserAgent(_userAgent string) error {
	r._userAgent = _userAgent
	r.Set("user_agent", _userAgent)
	return nil
}

// GetUserAgent UserAgent Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetUserAgent() string {
	return r._userAgent
}

// SetBatchNo is BatchNo Setter
// 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetBatchNo(_batchNo string) error {
	r._batchNo = _batchNo
	r.Set("batch_no", _batchNo)
	return nil
}

// GetBatchNo BatchNo Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetBatchNo() string {
	return r._batchNo
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaeleenterprisecouponsendAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaeleenterprisecouponsendAPIRequest) GetDeviceId() string {
	return r._deviceId
}
