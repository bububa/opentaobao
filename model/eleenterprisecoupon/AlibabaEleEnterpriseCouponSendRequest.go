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
    phone   string
    // 固定值，申请获取
    channel   string
    // 纬度
    latitude   string
    // 经度
    longitude   string
    // 客户端IP地址
    ip   string
    // 客户端User-Agent信息
    userAgent   string
    // 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
    batchNo   string
    // 设备ID
    deviceId   string
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
func (r *AlibabaEleEnterpriseCouponSendRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetPhone() string {
    return r.phone
}
// Channel Setter
// 固定值，申请获取
func (r *AlibabaEleEnterpriseCouponSendRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetChannel() string {
    return r.channel
}
// Latitude Setter
// 纬度
func (r *AlibabaEleEnterpriseCouponSendRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetLatitude() string {
    return r.latitude
}
// Longitude Setter
// 经度
func (r *AlibabaEleEnterpriseCouponSendRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetLongitude() string {
    return r.longitude
}
// Ip Setter
// 客户端IP地址
func (r *AlibabaEleEnterpriseCouponSendRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetIp() string {
    return r.ip
}
// UserAgent Setter
// 客户端User-Agent信息
func (r *AlibabaEleEnterpriseCouponSendRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

// UserAgent Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetUserAgent() string {
    return r.userAgent
}
// BatchNo Setter
// 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
func (r *AlibabaEleEnterpriseCouponSendRequest) SetBatchNo(batchNo string) error {
    r.batchNo = batchNo
    r.Set("batch_no", batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetBatchNo() string {
    return r.batchNo
}
// DeviceId Setter
// 设备ID
func (r *AlibabaEleEnterpriseCouponSendRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaEleEnterpriseCouponSendRequest) GetDeviceId() string {
    return r.deviceId
}
