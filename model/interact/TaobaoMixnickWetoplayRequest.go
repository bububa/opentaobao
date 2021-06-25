package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘混淆nick转互动混淆nick APIRequest
taobao.mixnick.wetoplay

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickWetoplayRequest struct {
    model.Params

    // 排查问题id
    traceId   string 

    // 微淘混淆nick
    weMixnick   string 

}

func NewTaobaoMixnickWetoplayRequest() *TaobaoMixnickWetoplayRequest{
    return &TaobaoMixnickWetoplayRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMixnickWetoplayRequest) GetApiMethodName() string {
    return "taobao.mixnick.wetoplay"
}

func (r TaobaoMixnickWetoplayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMixnickWetoplayRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

func (r TaobaoMixnickWetoplayRequest) GetTraceId() string {
    return r.traceId
}

func (r *TaobaoMixnickWetoplayRequest) SetWeMixnick(weMixnick string) error {
    r.weMixnick = weMixnick
    r.Set("we_mixnick", weMixnick)
    return nil
}

func (r TaobaoMixnickWetoplayRequest) GetWeMixnick() string {
    return r.weMixnick
}

