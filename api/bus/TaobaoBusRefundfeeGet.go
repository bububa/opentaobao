package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
查询退票费用明细 
taobao.bus.refundfee.get

查询退票的费用信息
*/
func TaobaoBusRefundfeeGet(clt *core.SDKClient, req *bus.TaobaoBusRefundfeeGetAPIRequest, session string) (*bus.TaobaoBusRefundfeeGetAPIResponse, error) {
    var resp bus.TaobaoBusRefundfeeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
