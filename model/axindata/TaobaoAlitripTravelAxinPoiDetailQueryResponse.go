package axindata

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
景点poi详情查询-阿信 APIResponse
taobao.alitrip.travel.axin.poi.detail.query

景点poi详情查询-阿信
*/
type TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelAxinPoiDetailQueryResponse
}

type TaobaoAlitripTravelAxinPoiDetailQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_axin_poi_detail_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoAlitripTravelAxinPoiDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
