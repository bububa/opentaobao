package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
自助机失效二维码 
taobao.bus.disableqrcode.set

使创建的二维码失效
*/
func TaobaoBusDisableqrcodeSet(clt *core.SDKClient, req *bus.TaobaoBusDisableqrcodeSetRequest, session string) (*bus.TaobaoBusDisableqrcodeSetAPIResponse, error) {
    var resp bus.TaobaoBusDisableqrcodeSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
