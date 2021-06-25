package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取词包列表 
taobao.subway.wordpackage.get

获取流量智选、捡漏词包等词包列表
*/
func TaobaoSubwayWordpackageGet(clt *core.SDKClient, req *simba.TaobaoSubwayWordpackageGetRequest, session string) (*simba.TaobaoSubwayWordpackageGetResponse, error) {
    var resp simba.TaobaoSubwayWordpackageGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
