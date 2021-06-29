package mc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云码设备圈选情况查询 API请求
tmall.mc.device.circle.check

云码设备圈选情况查询
*/
type TmallMcDeviceCircleCheckRequest struct {
    model.Params
    // 外部设备编码
    outerCode   string
    // 渠道编码
    channelId   string
}

// 初始化TmallMcDeviceCircleCheckRequest对象
func NewTmallMcDeviceCircleCheckRequest() *TmallMcDeviceCircleCheckRequest{
    return &TmallMcDeviceCircleCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMcDeviceCircleCheckRequest) GetApiMethodName() string {
    return "tmall.mc.device.circle.check"
}

// IRequest interface 方法, 获取API参数
func (r TmallMcDeviceCircleCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterCode Setter
// 外部设备编码
func (r *TmallMcDeviceCircleCheckRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

// OuterCode Getter
func (r TmallMcDeviceCircleCheckRequest) GetOuterCode() string {
    return r.outerCode
}
// ChannelId Setter
// 渠道编码
func (r *TmallMcDeviceCircleCheckRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r TmallMcDeviceCircleCheckRequest) GetChannelId() string {
    return r.channelId
}
