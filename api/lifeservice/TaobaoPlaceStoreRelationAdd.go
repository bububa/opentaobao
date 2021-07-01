package lifeservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lifeservice"
)

/* 
门店关系新增 
taobao.place.store.relation.add

新增授权用户的门店关系信息
*/
func TaobaoPlaceStoreRelationAdd(clt *core.SDKClient, req *lifeservice.TaobaoPlaceStoreRelationAddAPIRequest, session string) (*lifeservice.TaobaoPlaceStoreRelationAddAPIResponse, error) {
    var resp lifeservice.TaobaoPlaceStoreRelationAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
