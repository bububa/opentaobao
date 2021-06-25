package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
商品查询接口 
taobao.qimen.singleitem.query

商品查询接口
*/
func TaobaoQimenSingleitemQuery(clt *core.SDKClient, req *qimen.TaobaoQimenSingleitemQueryRequest, session string) (*qimen.TaobaoQimenSingleitemQueryResponse, error) {
    var resp qimen.TaobaoQimenSingleitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
