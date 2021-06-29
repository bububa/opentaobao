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
    _ext   string
    // 签名
    _signature   string
    // 随机值
    _nonceStr   string
    // 时间戳
    _timestampStr   string
    // 地震信息
    _earthquakeInfo   string
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
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetExt() string {
    return r._ext
}
// Signature Setter
// 签名
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetSignature() string {
    return r._signature
}
// NonceStr Setter
// 随机值
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetNonceStr(_nonceStr string) error {
    r._nonceStr = _nonceStr
    r.Set("nonce_str", _nonceStr)
    return nil
}

// NonceStr Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetNonceStr() string {
    return r._nonceStr
}
// TimestampStr Setter
// 时间戳
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetTimestampStr(_timestampStr string) error {
    r._timestampStr = _timestampStr
    r.Set("timestamp_str", _timestampStr)
    return nil
}

// TimestampStr Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetTimestampStr() string {
    return r._timestampStr
}
// EarthquakeInfo Setter
// 地震信息
func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetEarthquakeInfo(_earthquakeInfo string) error {
    r._earthquakeInfo = _earthquakeInfo
    r.Set("earthquake_info", _earthquakeInfo)
    return nil
}

// EarthquakeInfo Getter
func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetEarthquakeInfo() string {
    return r._earthquakeInfo
}
