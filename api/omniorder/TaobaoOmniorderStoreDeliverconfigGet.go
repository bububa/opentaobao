package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
查询门店发货配置内容 
taobao.omniorder.store.deliverconfig.get

查询门店发货配置内容
*/
func TaobaoOmniorderStoreDeliverconfigGet(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreDeliverconfigGetAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreDeliverconfigGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreDeliverconfigGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
