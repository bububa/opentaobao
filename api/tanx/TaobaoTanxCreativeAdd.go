package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
创意预审新增接口 
taobao.tanx.creative.add

创意预审新增接口
*/
func TaobaoTanxCreativeAdd(clt *core.SDKClient, req *tanx.TaobaoTanxCreativeAddRequest, session string) (*tanx.TaobaoTanxCreativeAddAPIResponse, error) {
    var resp tanx.TaobaoTanxCreativeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
