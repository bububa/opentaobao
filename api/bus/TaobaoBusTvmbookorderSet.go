package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
线下自助机通知出票接口 
taobao.bus.tvmbookorder.set

出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
*/
func TaobaoBusTvmbookorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmbookorderSetRequest, session string) (*bus.TaobaoBusTvmbookorderSetResponse, error) {
    var resp bus.TaobaoBusTvmbookorderSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
