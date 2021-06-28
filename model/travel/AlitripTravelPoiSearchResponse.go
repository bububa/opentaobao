package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
POI信息查询 APIResponse
alitrip.travel.poi.search

POI信息查询，用于商品更新使用
*/
type AlitripTravelPoiSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_poi_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // POI详情
    
    Results   []Poi `json:"results,omitempty" xml:"