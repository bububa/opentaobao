package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
配送拦截接口 
taobao.qimen.order.callback

配送拦截
*/
func TaobaoQimenOrderCallback(clt *core.SDKClient, req *qimen.TaobaoQimenOrderCallbackRequest, session string) (*qimen.TaobaoQimenOrderCallbackResponse, error) {
    var resp qimen.TaobaoQimenOrderCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
