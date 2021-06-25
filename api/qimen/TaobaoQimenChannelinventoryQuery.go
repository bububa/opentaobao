package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
渠道库存查询接口 
taobao.qimen.channelinventory.query

渠道库存查询
*/
func TaobaoQimenChannelinventoryQuery(clt *core.SDKClient, req *qimen.TaobaoQimenChannelinventoryQueryRequest, session string) (*qimen.TaobaoQimenChannelinventoryQueryResponse, error) {
    var resp qimen.TaobaoQimenChannelinventoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
