package eleenterprisecoupon

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发放优惠券 APIRequest
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

func NewAlibabaEleEnterpriseCouponSendRequest() *AlibabaEleEnterpriseCouponSendRequest{
    return &AlibabaEleEnterpriseCouponSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.coupon.send"
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseCouponSendRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetPhone() string {
    return r.phone
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetChannel() string {
    return r.channel
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetLatitude() string {
    return r.latitude
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetLongitude() string {
    return r.longitude
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetIp() string {
    return r.ip
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetUserAgent() string {
    return r.userAgent
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetBatchNo(batchNo string) error {
    r.batchNo = batchNo
    r.Set("batch_no", batchNo)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetBatchNo() string {
    return r.batchNo
}

func (r *AlibabaEleEnterpriseCouponSendRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaEleEnterpriseCouponSendRequest) GetDeviceId() string {
    return r.deviceId
}

