package alitrippoi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
POI开放存储能力 API返回值 
alitrip.platform.poi.raw.saverawpoi

POI开放存储提供离线/在线/纬错更新的能力
*/
type AlitripPlatformPoiRawSaverawpoiAPIResponse struct {
    model.CommonResponse
    AlitripPlatformPoiRawSaverawpoiAPIResponseModel
}

// POI开放存储能力 成功返回结果
type AlitripPlatformPoiRawSaverawpoiAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_platform_poi_raw_saverawpoi_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlitripPlatformPoiRawSaverawpoiResult `json:"result,omitempty" xml:"result,omitempty"`
}
