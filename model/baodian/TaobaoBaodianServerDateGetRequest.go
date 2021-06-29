package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务器时间获取 API请求
taobao.baodian.server.date.get

获取服务器时间
*/
type TaobaoBaodianServerDateGetRequest struct {
    model.Params
}

// 初始化TaobaoBaodianServerDateGetRequest对象
func NewTaobaoBaodianServerDateGetRequest() *TaobaoBaodianServerDateGetRequest{
    return &TaobaoBaodianServerDateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaodianServerDateGetRequest) GetApiMethodName() string {
    return "taobao.baodian.server.date.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaodianServerDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
