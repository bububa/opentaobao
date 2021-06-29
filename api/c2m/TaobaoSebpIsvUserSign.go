package c2m

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/c2m"
)

/* 
淘小铺三方签约同步 
taobao.sebp.isv.user.sign

同步淘小铺三方服务签约信息
*/
func TaobaoSebpIsvUserSign(clt *core.SDKClient, req *c2m.TaobaoSebpIsvUserSignRequest, session string) (*c2m.TaobaoSebpIsvUserSignAPIResponse, error) {
    var resp c2m.TaobaoSebpIsvUserSignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
