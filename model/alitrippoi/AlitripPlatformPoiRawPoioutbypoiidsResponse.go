package alitrippoi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据poiId输出飞猪poi数据 APIResponse
alitrip.platform.poi.raw.poioutbypoiids

根据poiId输出飞猪poi数据
*/
type AlitripPlatformPoiRawPoioutbypoiidsAPIResponse struct {
    model.CommonResponse
    AlitripPlatformPoiRawPoioutbypoiidsResponse
}

type AlitripPlatformPoiRawPoioutbypoiidsResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_poi_raw_poioutbypoiids_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlitripPlatformPoiRawPoioutbypoiidsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
