package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动mixNick转微淘 API请求
taobao.mixnick.playtowe

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickPlaytoweRequest struct {
    model.Params
    // 用户的混淆nick
    _mixMix   string
}

// 初始化TaobaoMixnickPlaytoweRequest对象
func NewTaobaoMixnickPlaytoweRequest() *TaobaoMixnickPlaytoweRequest{
    return &TaobaoMixnickPlaytoweRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMixnickPlaytoweRequest) GetApiMethodName() string {
    return "taobao.mixnick.playtowe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMixnickPlaytoweRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MixMix Setter
// 用户的混淆nick
func (r *TaobaoMixnickPlaytoweRequest) SetMixMix(_mixMix string) error {
    r._mixMix = _mixMix
    r.Set("mix_mix", _mixMix)
    return nil
}

// MixMix Getter
func (r TaobaoMixnickPlaytoweRequest) GetMixMix() string {
    return r._mixMix
}
