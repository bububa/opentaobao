package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
修改门店发货配置内容 
taobao.omniorder.store.deliverconfig.update

修改门店发货配置内容
*/
func TaobaoOmniorderStoreDeliverconfigUpdate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreDeliverconfigUpdateRequest, session string) (*omniorder.TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}