package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动mixNick转微淘 APIRequest
taobao.mixnick.playtowe

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickPlaytoweRequest struct {
    model.Params

    // 用户的混淆nick
    mixMix   string 

}

func NewTaobaoMixnickPlaytoweRequest() *TaobaoMixnickPlaytoweRequest{
    return &TaobaoMixnickPlaytoweRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMixnickPlaytoweRequest) GetApiMethodName() string {
    return "taobao.mixnick.playtowe"
}

func (r TaobaoMixnickPlaytoweRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMixnickPlaytoweRequest) SetMixMix(mixMix string) error {
    r.mixMix = mixMix
    r.Set("mix_mix", mixMix)
    return nil
}

func (r TaobaoMixnickPlaytoweRequest) GetMixMix() string {
    return r.mixMix
}

