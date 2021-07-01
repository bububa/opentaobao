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
func TaobaoQimenChannelinventoryQuery(clt *core.SDKClient, req *qimen.TaobaoQimenChannelinventoryQueryAPIRequest, session string) (*qimen.TaobaoQimenChannelinventoryQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenChannelinventoryQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
