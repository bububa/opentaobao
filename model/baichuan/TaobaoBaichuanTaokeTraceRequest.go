package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川淘客打点 API请求
taobao.baichuan.taoke.trace

百川淘客打点
*/
type TaobaoBaichuanTaokeTraceRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanTaokeTraceRequest对象
func NewTaobaoBaichuanTaokeTraceRequest() *TaobaoBaichuanTaokeTraceRequest{
    return &TaobaoBaichuanTaokeTraceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanTaokeTraceRequest) GetApiMethodName() string {
    return "taobao.baichuan.taoke.trace"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanTaokeTraceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanTaokeTraceRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanTaokeTraceRequest) GetName() string {
    return r._name
}
