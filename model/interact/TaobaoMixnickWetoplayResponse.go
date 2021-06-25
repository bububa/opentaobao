package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
微淘混淆nick转互动混淆nick APIResponse
taobao.mixnick.wetoplay

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickWetoplayAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMixnickWetoplayResponse `json:"taobao_mixnick_wetoplay_response,omitempty"`
}

type TaobaoMixnickWetoplayResponse struct {

    // 微淘转互动混淆nick
    PlayMixnickData   *MixNickResult `json:"play_mixnick_data,omitempty"`

}
