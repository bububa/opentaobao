package axindata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
景点poi搜索-阿信 APIResponse
taobao.alitrip.travel.axin.poi.search

给阿信提供景点poi搜索
*/
type TaobaoAlitripTravelAxinPoiSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelAxinPoiSearchResponse
}

type TaobaoAlitripTravelAxinPoiSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_axin_poi_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoAlitripTravelAxinPoiSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
