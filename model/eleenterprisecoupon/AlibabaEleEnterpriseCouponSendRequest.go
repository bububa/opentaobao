package eleenterprisecoupon

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发放优惠券 API请求
alibaba.ele.enterprise.coupon.send

发放优惠券
*/
type AlibabaEleEnterpriseCouponSendRequest struct {
    model.Params
    // 手机号
    _phone   string
    // 固定值，申请获取
    _channel   string
    // 纬度
    _latitude   string
    // 经度
    _longitude   string
    // 客户端IP地址
    _ip   string
    // 客户端User-Agent信息
    _userAgent   string
    // 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
    _batchNo   string
    // 设备ID
    _deviceId   string
}

// 初始化AlibabaEleEnterpriseCouponSendRequest对象
func NewAlibabaEleEnterpriseCouponSendRequest() *AlibabaEleEnterpriseCouponSendRequest{
    return &AlibabaEleEnterpriseCouponSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCouponSendRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.coupon.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCouponSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 手机号
func (r *AlibabaEleEnterpriseCouponSendRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetPhone() string {
    return r._phone
}
// Channel Setter
// 固定值，申请获取
func (r *AlibabaEleEnterpriseCouponSendRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetChannel() string {
    return r._channel
}
// Latitude Setter
// 纬度
func (r *AlibabaEleEnterpriseCouponSendRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetLatitude() string {
    return r._latitude
}
// Longitude Setter
// 经度
func (r *AlibabaEleEnterpriseCouponSendRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetLongitude() string {
    return r._longitude
}
// Ip Setter
// 客户端IP地址
func (r *AlibabaEleEnterpriseCouponSendRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetIp() string {
    return r._ip
}
// UserAgent Setter
// 客户端User-Agent信息
func (r *AlibabaEleEnterpriseCouponSendRequest) SetUserAgent(_userAgent string) error {
    r._userAgent = _userAgent
    r.Set("user_agent", _userAgent)
    return nil
}

// UserAgent Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetUserAgent() string {
    return r._userAgent
}
// BatchNo Setter
// 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
func (r *AlibabaEleEnterpriseCouponSendRequest) SetBatchNo(_batchNo string) error {
    r._batchNo = _batchNo
    r.Set("batch_no", _batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetBatchNo() string {
    return r._batchNo
}
// DeviceId Setter
// 设备ID
func (r *AlibabaEleEnterpriseCouponSendRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetDeviceId() string {
    return r._deviceId
}
