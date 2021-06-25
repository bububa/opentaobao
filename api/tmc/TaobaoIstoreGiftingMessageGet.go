package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
gifting消息获取 
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息
*/
func TaobaoIstoreGiftingMessageGet(clt *core.SDKClient, req *tmc.TaobaoIstoreGiftingMessageGetRequest, session string) (*tmc.TaobaoIstoreGiftingMessageGetResponse, error) {
    var resp tmc.TaobaoIstoreGiftingMessageGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
