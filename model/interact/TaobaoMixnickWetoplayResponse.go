package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘混淆nick转互动混淆nick API返回值 
taobao.mixnick.wetoplay

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickWetoplayAPIResponse struct {
    model.CommonResponse
    TaobaoMixnickWetoplayResponse
}

// 微淘混淆nick转互动混淆nick 成功返回结果
type TaobaoMixnickWetoplayResponse struct {
    XMLName xml.Name `xml:"mixnick_wetoplay_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 微淘转互动混淆nick
    PlayMixnickData   *MixNickResult `json:"play_mixnick_data,omitempty" xml:"play_mixnick_data,omitempty"`
}
