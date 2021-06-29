package mc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云码权益查询 API请求
aliyun.unimkt.task.charge.launch

云码线上流量投放链路，用于判断用户是否有匹配的投放计划
*/
type AliyunUnimktTaskChargeLaunchRequest struct {
    model.Params
    // 服务商附加url参数
    extra   string
    // urlID
    urlId   string
    // 支付宝openID
    alipayOpenId   string
    // 渠道ID
    channelId   string
    // 淘宝ID
    userId   string
}

// 初始化AliyunUnimktTaskChargeLaunchRequest对象
func NewAliyunUnimktTaskChargeLaunchRequest() *AliyunUnimktTaskChargeLaunchRequest{
    return &AliyunUnimktTaskChargeLaunchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunUnimktTaskChargeLaunchRequest) GetApiMethodName() string {
    return "aliyun.unimkt.task.charge.launch"
}

// IRequest interface 方法, 获取API参数
func (r AliyunUnimktTaskChargeLaunchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extra Setter
// 服务商附加url参数
func (r *AliyunUnimktTaskChargeLaunchRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r AliyunUnimktTaskChargeLaunchRequest) GetExtra() string {
    return r.extra
}
// UrlId Setter
// urlID
func (r *AliyunUnimktTaskChargeLaunchRequest) SetUrlId(urlId string) error {
    r.urlId = urlId
    r.Set("url_id", urlId)
    return nil
}

// UrlId Getter
func (r AliyunUnimktTaskChargeLaunchRequest) GetUrlId() string {
    return r.urlId
}
// AlipayOpenId Setter
// 支付宝openID
func (r *AliyunUnimktTaskChargeLaunchRequest) SetAlipayOpenId(alipayOpenId string) error {
    r.alipayOpenId = alipayOpenId
    r.Set("alipay_open_id", alipayOpenId)
    return nil
}

// AlipayOpenId Getter
func (r AliyunUnimktTaskChargeLaunchRequest) GetAlipayOpenId() string {
    return r.alipayOpenId
}
// ChannelId Setter
// 渠道ID
func (r *AliyunUnimktTaskChargeLaunchRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r AliyunUnimktTaskChargeLaunchRequest) GetChannelId() string {
    return r.channelId
}
// UserId Setter
// 淘宝ID
func (r *AliyunUnimktTaskChargeLaunchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AliyunUnimktTaskChargeLaunchRequest) GetUserId() string {
    return r.userId
}
