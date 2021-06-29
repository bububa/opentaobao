package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地震局发送地震消息 API请求
taobao.ailab.aicloud.top.earthquake.send

地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
*/
type TaobaoAilabAicloudTopEarthquakeSendRequest struct {
    model.Params
    // 扩展占位字段
    ext   string
    // 签名
    signature   string
    // 随机值
    nonceStr   string
    // 时间戳
    timestampStr   string
    // 地震信息
    earthquakeInfo   string
}

// 初始化TaobaoAilabAicloudTopEarthquakeSendRequest对象
func NewTaobaoAilabAicloudTopEarthquakeSendRequest() *TaobaoAilabAicloudTopEarthquakeSendRequest{
    return &TaobaoAilabAicloudTopEarthquakeSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.earthquake.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展占位字段
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetExt() string {
    return r.ext
}
// Signature Setter
// 签名
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

// Signature Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetSignature() string {
    return r.signature
}
// NonceStr Setter
// 随机值
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetNonceStr(nonceStr string) error {
    r.nonceStr = nonceStr
    r.Set("nonce_str", nonceStr)
    return nil
}

// NonceStr Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetNonceStr() string {
    return r.nonceStr
}
// TimestampStr Setter
// 时间戳
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetTimestampStr(timestampStr string) error {
    r.timestampStr = timestampStr
    r.Set("timestamp_str", timestampStr)
    return nil
}

// TimestampStr Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetTimestampStr() string {
    return r.timestampStr
}
// EarthquakeInfo Setter
// 地震信息
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetEarthquakeInfo(earthquakeInfo string) error {
    r.earthquakeInfo = earthquakeInfo
    r.Set("earthquake_info", earthquakeInfo)
    return nil
}

// EarthquakeInfo Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetEarthquakeInfo() string {
    return r.earthquakeInfo
}
