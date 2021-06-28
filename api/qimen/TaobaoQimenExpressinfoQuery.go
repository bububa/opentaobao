package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
配送公司信息查询接口 
taobao.qimen.expressinfo.query

配送公司信息查询
*/
func TaobaoQimenExpressinfoQuery(clt *core.SDKClient, req *qimen.TaobaoQimenExpressinfoQueryRequest, session string) (*qimen.TaobaoQimenExpressinfoQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenExpressinfoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
