package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取捡漏词包分时报表数据 
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据
*/
func TaobaoSubwayMarshLandRptGet(clt *core.SDKClient, req *simba.TaobaoSubwayMarshLandRptGetRequest, session string) (*simba.TaobaoSubwayMarshLandRptGetResponse, error) {
    var resp simba.TaobaoSubwayMarshLandRptGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
