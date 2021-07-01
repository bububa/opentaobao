package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
POI信息查询 
alitrip.travel.poi.search

POI信息查询，用于商品更新使用
*/
func AlitripTravelPoiSearch(clt *core.SDKClient, req *travel.AlitripTravelPoiSearchAPIRequest, session string) (*travel.AlitripTravelPoiSearchAPIResponse, error) {
    var resp travel.AlitripTravelPoiSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
