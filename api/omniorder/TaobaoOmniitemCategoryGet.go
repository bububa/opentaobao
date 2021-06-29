package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
全渠道商品轻发布类目信息 
taobao.omniitem.category.get

全渠道商品轻发布类目信息
*/
func TaobaoOmniitemCategoryGet(clt *core.SDKClient, req *omniorder.TaobaoOmniitemCategoryGetRequest, session string) (*omniorder.TaobaoOmniitemCategoryGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniitemCategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
