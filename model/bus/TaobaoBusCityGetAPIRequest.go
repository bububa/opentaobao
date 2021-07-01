package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城市接口 API请求
taobao.bus.city.get

汽车票出发城市获取接口，获取所有出发城市
*/
type TaobaoBusCityGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoBusCityGetAPIRequest对象
func NewTaobaoBusCityGetRequest() *TaobaoBusCityGetAPIRequest{
    return &TaobaoBusCityGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusCityGetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.city.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusCityGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
