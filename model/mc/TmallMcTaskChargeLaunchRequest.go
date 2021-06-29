package mc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云码充电宝投放链路 API请求
tmall.mc.task.charge.launch

云码充电宝投放链路，用于判断用户是否有匹配的投放计划
*/
type TmallMcTaskChargeLaunchRequest struct {
    model.Params
    // 外部设备编码
    outerCode   string
    // 渠道ID
    channelId   string
    // 支付宝openID
    alipayOpenId   string
    // urlID
    urlId   string
    // 服务商附加url参数
    extra   string
}

// 初始化TmallMcTaskChargeLaunchRequest对象
func NewTmallMcTaskChargeLaunchRequest() *TmallMcTaskChargeLaunchRequest{
    return &TmallMcTaskChargeLaunchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMcTaskChargeLaunchRequest) GetApiMethodName() string {
    return "tmall.mc.task.charge.launch"
}

// IRequest interface 方法, 获取API参数
func (r TmallMcTaskChargeLaunchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterCode Setter
// 外部设备编码
func (r *TmallMcTaskChargeLaunchRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

// OuterCode Getter
func (r TmallMcTaskChargeLaunchRequest) GetOuterCode() string {
    return r.outerCode
}
// ChannelId Setter
// 渠道ID
func (r *TmallMcTaskChargeLaunchRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r TmallMcTaskChargeLaunchRequest) GetChannelId() string {
    return r.channelId
}
// AlipayOpenId Setter
// 支付宝openID
func (r *TmallMcTaskChargeLaunchRequest) SetAlipayOpenId(alipayOpenId string) error {
    r.alipayOpenId = alipayOpenId
    r.Set("alipay_open_id", alipayOpenId)
    return nil
}

// AlipayOpenId Getter
func (r TmallMcTaskChargeLaunchRequest) GetAlipayOpenId() string {
    return r.alipayOpenId
}
// UrlId Setter
// urlID
func (r *TmallMcTaskChargeLaunchRequest) SetUrlId(urlId string) error {
    r.urlId = urlId
    r.Set("url_id", urlId)
    return nil
}

// UrlId Getter
func (r TmallMcTaskChargeLaunchRequest) GetUrlId() string {
    return r.urlId
}
// Extra Setter
// 服务商附加url参数
func (r *TmallMcTaskChargeLaunchRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r TmallMcTaskChargeLaunchRequest) GetExtra() string {
    return r.extra
}
