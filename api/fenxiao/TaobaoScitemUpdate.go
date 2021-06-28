package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
根据商品ID或商家编码修改后端商品 
taobao.scitem.update

根据商品ID或商家编码修改后端商品
*/
func TaobaoScitemUpdate(clt *core.SDKClient, req *fenxiao.TaobaoScitemUpdateRequest, session string) (*fenxiao.TaobaoScitemUpdateAPIResponse, error) {
    var resp fenxiao.TaobaoScitemUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
