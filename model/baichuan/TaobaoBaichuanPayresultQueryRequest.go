package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川支付完成回调 API请求
taobao.baichuan.payresult.query

百川支付完成回调
*/
type TaobaoBaichuanPayresultQueryRequest struct {
    model.Params
    // name
    _name   string
}

// 初始化TaobaoBaichuanPayresultQueryRequest对象
func NewTaobaoBaichuanPayresultQueryRequest() *TaobaoBaichuanPayresultQueryRequest{
    return &TaobaoBaichuanPayresultQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanPayresultQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.payresult.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanPayresultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanPayresultQueryRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanPayresultQueryRequest) GetName() string {
    return r._name
}
