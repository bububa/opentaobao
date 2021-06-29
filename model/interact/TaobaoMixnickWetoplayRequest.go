package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘混淆nick转互动混淆nick API请求
taobao.mixnick.wetoplay

微淘应用的混淆nick转为互动类型混淆nick
*/
type TaobaoMixnickWetoplayRequest struct {
    model.Params
    // 排查问题id
    _traceId   string
    // 微淘混淆nick
    _weMixnick   string
}

// 初始化TaobaoMixnickWetoplayRequest对象
func NewTaobaoMixnickWetoplayRequest() *TaobaoMixnickWetoplayRequest{
    return &TaobaoMixnickWetoplayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMixnickWetoplayRequest) GetApiMethodName() string {
    return "taobao.mixnick.wetoplay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMixnickWetoplayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceId Setter
// 排查问题id
func (r *TaobaoMixnickWetoplayRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoMixnickWetoplayRequest) GetTraceId() string {
    return r._traceId
}
// WeMixnick Setter
// 微淘混淆nick
func (r *TaobaoMixnickWetoplayRequest) SetWeMixnick(_weMixnick string) error {
    r._weMixnick = _weMixnick
    r.Set("we_mixnick", _weMixnick)
    return nil
}

// WeMixnick Getter
func (r TaobaoMixnickWetoplayRequest) GetWeMixnick() string {
    return r._weMixnick
}
