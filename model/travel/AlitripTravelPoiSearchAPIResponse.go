package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
POI信息查询 API返回值 
alitrip.travel.poi.search

POI信息查询，用于商品更新使用
*/
type AlitripTravelPoiSearchAPIResponse struct {
    model.CommonResponse
    AlitripTravelPoiSearchAPIResponseModel
}

// POI信息查询 成功返回结果
type AlitripTravelPoiSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_travel_poi_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // POI详情
    Results   []Poi `json:"results,omitempty" xml:"results>poi,omitempty"`
}
