package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
组合货品删除接口 
taobao.qimen.combineitem.delete

组合货品删除
*/
func TaobaoQimenCombineitemDelete(clt *core.SDKClient, req *qimen.TaobaoQimenCombineitemDeleteRequest, session string) (*qimen.TaobaoQimenCombineitemDeleteResponse, error) {
    var resp qimen.TaobaoQimenCombineitemDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
