package c2m

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/c2m"
)

/* 
isv机器人回调接口 
taobao.sebp.isv.wxrobot.callback

机器人入群回调，进行校验、功能开通等操作
*/
func TaobaoSebpIsvWxrobotCallback(clt *core.SDKClient, req *c2m.TaobaoSebpIsvWxrobotCallbackAPIRequest, session string) (*c2m.TaobaoSebpIsvWxrobotCallbackAPIResponse, error) {
    var resp c2m.TaobaoSebpIsvWxrobotCallbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
