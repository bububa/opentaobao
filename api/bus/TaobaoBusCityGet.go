package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
城市接口 
taobao.bus.city.get

汽车票出发城市获取接口，获取所有出发城市
*/
func TaobaoBusCityGet(clt *core.SDKClient, req *bus.TaobaoBusCityGetAPIRequest, session string) (*bus.TaobaoBusCityGetAPIResponse, error) {
    var resp bus.TaobaoBusCityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
