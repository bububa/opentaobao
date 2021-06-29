package alitrippoi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪poi输出 API返回值 
alitrip.platform.poi.raw.poiout

输出指定城市poi指定信息
*/
type AlitripPlatformPoiRawPoioutAPIResponse struct {
    model.CommonResponse
    AlitripPlatformPoiRawPoioutResponse
}

// 飞猪poi输出 成功返回结果
type AlitripPlatformPoiRawPoioutResponse struct {
    XMLName xml.Name `xml:"alitrip_platform_poi_raw_poiout_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlitripPlatformPoiRawPoioutResult `json:"result,omitempty" xml:"result,omitempty"`
}
