package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
批量更新词包 
taobao.subway.wordpackage.update

批量更新词包
*/
func TaobaoSubwayWordpackageUpdate(clt *core.SDKClient, req *simba.TaobaoSubwayWordpackageUpdateRequest, session string) (*simba.TaobaoSubwayWordpackageUpdateAPIResponse, error) {
    var resp simba.TaobaoSubwayWordpackageUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
