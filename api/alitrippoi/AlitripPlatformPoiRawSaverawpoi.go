package alitrippoi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitrippoi"
)

/* 
POI开放存储能力 
alitrip.platform.poi.raw.saverawpoi

POI开放存储提供离线/在线/纬错更新的能力
*/
func AlitripPlatformPoiRawSaverawpoi(clt *core.SDKClient, req *alitrippoi.AlitripPlatformPoiRawSaverawpoiAPIRequest, session string) (*alitrippoi.AlitripPlatformPoiRawSaverawpoiAPIResponse, error) {
    var resp alitrippoi.AlitripPlatformPoiRawSaverawpoiAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
