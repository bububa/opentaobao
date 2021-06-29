package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
删除一个分类 
taobao.omniitem.classify.delete

删除一个分类
*/
func TaobaoOmniitemClassifyDelete(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyDeleteRequest, session string) (*omniorder.TaobaoOmniitemClassifyDeleteAPIResponse, error) {
    var resp omniorder.TaobaoOmniitemClassifyDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
