package alitrippoi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitrippoi"
)

/* 
飞猪poi输出 
alitrip.platform.poi.raw.poiout

输出指定城市poi指定信息
*/
func AlitripPlatformPoiRawPoiout(clt *core.SDKClient, req *alitrippoi.AlitripPlatformPoiRawPoioutRequest, session string) (*alitrippoi.AlitripPlatformPoiRawPoioutAPIResponse, error) {
    var resp alitrippoi.AlitripPlatformPoiRawPoioutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
