package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
组合货品关系查询接口 
taobao.qimen.combineitem.query

组合货品关系查询
*/
func TaobaoQimenCombineitemQuery(clt *core.SDKClient, req *qimen.TaobaoQimenCombineitemQueryAPIRequest, session string) (*qimen.TaobaoQimenCombineitemQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenCombineitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
