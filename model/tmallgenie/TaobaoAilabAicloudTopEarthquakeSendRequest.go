package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地震局发送地震消息 APIRequest
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

func NewTaobaoAilabAicloudTopEarthquakeSendRequest() *TaobaoAilabAicloudTopEarthquakeSendRequest{
    return &TaobaoAilabAicloudTopEarthquakeSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.earthquake.send"
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetSignature() string {
    return r.signature
}

func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetNonceStr(nonceStr string) error {
    r.nonceStr = nonceStr
    r.Set("nonce_str", nonceStr)
    return nil
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetNonceStr() string {
    return r.nonceStr
}

func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetTimestampStr(timestampStr string) error {
    r.timestampStr = timestampStr
    r.Set("timestamp_str", timestampStr)
    return nil
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetTimestampStr() string {
    return r.timestampStr
}

func (r *TaobaoAilabAicloudTopEarthquakeSendRequest) SetEarthquakeInfo(earthquakeInfo string) error {
    r.earthquakeInfo = earthquakeInfo
    r.Set("earthquake_info", earthquakeInfo)
    return nil
}

func (r TaobaoAilabAicloudTopEarthquakeSendRequest) GetEarthquakeInfo() string {
    return r.earthquakeInfo
}

