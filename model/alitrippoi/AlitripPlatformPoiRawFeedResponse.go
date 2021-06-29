package alitrippoi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
存储poi原始数据 APIResponse
alitrip.platform.poi.raw.feed

对接外部数据源，外部数据推送poi数据到飞猪
*/
type AlitripPlatformPoiRawFeedAPIResponse struct {
    model.CommonResponse
    AlitripPlatformPoiRawFeedResponse
}

type AlitripPlatformPoiRawFeedResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_poi_raw_feed_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlitripPlatformPoiRawFeedResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
