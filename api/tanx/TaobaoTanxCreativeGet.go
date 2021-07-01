package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
获取单个审核创意状态 
taobao.tanx.creative.get

获取单个审核创意状态
*/
func TaobaoTanxCreativeGet(clt *core.SDKClient, req *tanx.TaobaoTanxCreativeGetAPIRequest, session string) (*tanx.TaobaoTanxCreativeGetAPIResponse, error) {
    var resp tanx.TaobaoTanxCreativeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
