package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
批量添加/删除商品和分类的关联关系 
taobao.omniitem.classify.item.batch.operate

批量添加/删除商品和分类的关联关系
*/
func TaobaoOmniitemClassifyItemBatchOperate(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyItemBatchOperateAPIRequest, session string) (*omniorder.TaobaoOmniitemClassifyItemBatchOperateAPIResponse, error) {
    var resp omniorder.TaobaoOmniitemClassifyItemBatchOperateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
