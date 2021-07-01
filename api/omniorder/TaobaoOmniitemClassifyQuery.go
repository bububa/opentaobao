package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
查询分类信息 
taobao.omniitem.classify.query

通过查询关键字，分页查询分类信息
*/
func TaobaoOmniitemClassifyQuery(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyQueryAPIRequest, session string) (*omniorder.TaobaoOmniitemClassifyQueryAPIResponse, error) {
    var resp omniorder.TaobaoOmniitemClassifyQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
