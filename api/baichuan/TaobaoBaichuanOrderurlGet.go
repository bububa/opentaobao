package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川订单详情 
taobao.baichuan.orderurl.get

百川订单详情
*/
func TaobaoBaichuanOrderurlGet(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOrderurlGetRequest, session string) (*baichuan.TaobaoBaichuanOrderurlGetResponse, error) {
    var resp baichuan.TaobaoBaichuanOrderurlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
