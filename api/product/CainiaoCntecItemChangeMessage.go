package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品变更消息 
cainiao.cntec.item.change.message

供货商商品信息变更消息
*/
func CainiaoCntecItemChangeMessage(clt *core.SDKClient, req *product.CainiaoCntecItemChangeMessageRequest, session string) (*product.CainiaoCntecItemChangeMessageResponse, error) {
    var resp product.CainiaoCntecItemChangeMessageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
