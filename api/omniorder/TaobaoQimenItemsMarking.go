package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
商品通自动打标 
taobao.qimen.items.marking

调用该接口，对商品进行XXXX标的打标、去标的动作。
*/
func TaobaoQimenItemsMarking(clt *core.SDKClient, req *omniorder.TaobaoQimenItemsMarkingRequest, session string) (*omniorder.TaobaoQimenItemsMarkingAPIResponse, error) {
    var resp omniorder.TaobaoQimenItemsMarkingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
