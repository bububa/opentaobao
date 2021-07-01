package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
查询门店自提配置内容 
taobao.omniorder.store.collectconfig.get

查询门店自提配置内容
*/
func TaobaoOmniorderStoreCollectconfigGet(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreCollectconfigGetAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreCollectconfigGetAPIResponse, error) {
    var resp omniorder.TaobaoOmniorderStoreCollectconfigGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
