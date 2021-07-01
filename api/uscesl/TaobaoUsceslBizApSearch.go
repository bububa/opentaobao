package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
AP列表查询 
taobao.uscesl.biz.ap.search

查询当前门店下登记的AP列表
*/
func TaobaoUsceslBizApSearch(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApSearchAPIRequest, session string) (*uscesl.TaobaoUsceslBizApSearchAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizApSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
