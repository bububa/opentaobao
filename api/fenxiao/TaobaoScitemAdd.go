package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
发布后端商品 
taobao.scitem.add

发布后端商品
*/
func TaobaoScitemAdd(clt *core.SDKClient, req *fenxiao.TaobaoScitemAddRequest, session string) (*fenxiao.TaobaoScitemAddResponse, error) {
    var resp fenxiao.TaobaoScitemAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
