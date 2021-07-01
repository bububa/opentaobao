package alitrippoi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformPoiRawPoioutbypoiidsAPIResponse
根据poiId输出飞猪poi数据 API返回值
alitrip.platform.poi.raw.poioutbypoiids

根据poiId输出飞猪poi数据 */
type AlitripPlatformPoiRawPoioutbypoiidsAPIResponse struct {
	model.CommonResponse
	AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel
}

// AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel is 根据poiId输出飞猪poi数据 成功返回结果
type AlitripPlatformPoiRawPoioutbypoiidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_poi_raw_poioutbypoiids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripPlatformPoiRawPoioutbypoiidsResult `json:"result,omitempty" xml:"result,omitempty"`
}
