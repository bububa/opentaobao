package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
批量添加/删除门店和分类的关联关系 
taobao.omniitem.classify.store.batch.operate

批量添加/删除门店和分类的关联关系
*/
func TaobaoOmniitemClassifyStoreBatchOperate(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyStoreBatchOperateRequest, session string) (*omniorder.TaobaoOmniitemClassifyStoreBatchOperateAPIResponse, error) {
    var resp omniorder.TaobaoOmniitemClassifyStoreBatchOperateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
