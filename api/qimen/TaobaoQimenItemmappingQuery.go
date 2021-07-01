package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
前后端商品映射查询接口 
taobao.qimen.itemmapping.query

前后端商品映射查询接口
*/
func TaobaoQimenItemmappingQuery(clt *core.SDKClient, req *qimen.TaobaoQimenItemmappingQueryAPIRequest, session string) (*qimen.TaobaoQimenItemmappingQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenItemmappingQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
