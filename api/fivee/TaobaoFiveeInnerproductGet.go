package fivee

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fivee"
)

/* 
国产商品资质查询 
taobao.fivee.innerproduct.get

资质共享平台，国产商品查询
*/
func TaobaoFiveeInnerproductGet(clt *core.SDKClient, req *fivee.TaobaoFiveeInnerproductGetAPIRequest, session string) (*fivee.TaobaoFiveeInnerproductGetAPIResponse, error) {
    var resp fivee.TaobaoFiveeInnerproductGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
