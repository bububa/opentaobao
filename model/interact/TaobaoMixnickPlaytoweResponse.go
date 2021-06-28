package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
互动mixNick转微淘 APIResponse
taobao.mixnick.playtowe

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickPlaytoweAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMixnickPlaytoweResponse `json:"mixnick_playtowe_response,omitempty"` 
    TaobaoMixnickPlaytoweResponse
}

/* model for simplify = false
type TaobaoMixnickPlaytoweResponse struct {

    // 微淘混淆nick
    
    WeMixnick   string `json:"we_mixnick,omitempty"`
    

}
*/

type TaobaoMixnickPlaytoweResponse struct {

    // 微淘混淆nick
    WeMixnick   string `json:"we_mixnick,omitempty"`

}
