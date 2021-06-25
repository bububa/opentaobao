package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
查询单元智能出价信息 
taobao.subway.cia.get

查询单元智能出价信息
*/
func TaobaoSubwayCiaGet(clt *core.SDKClient, req *simba.TaobaoSubwayCiaGetRequest, session string) (*simba.TaobaoSubwayCiaGetResponse, error) {
    var resp simba.TaobaoSubwayCiaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
