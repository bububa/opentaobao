package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
POI信息查询 APIResponse
alitrip.travel.poi.search

POI信息查询，用于商品更新使用
*/
type AlitripTravelPoiSearchAPIResponse struct {
    model.CommonResponse
    // Response *AlitripTravelPoiSearchResponse `json:"alitrip_travel_poi_search_response,omitempty"` 
    AlitripTravelPoiSearchResponse
}

/* model for simplify = false
type AlitripTravelPoiSearchResponse struct {

    // POI详情
    
    Results  struct {
        Poi  []Poi `json:"poi,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type AlitripTravelPoiSearchResponse struct {

    // POI详情
    Results   []Poi `json:"results,omitempty"`

}
