package alitrippoi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformpoirawpoioutbypoiidsAPIResponse 根据poiId输出飞猪poi数据 API返回值
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
type AlitripplatformpoirawpoioutbypoiidsAPIResponse struct {
	model.CommonResponse
	AlitripplatformpoirawpoioutbypoiidsAPIResponseModel
}

// AlitripplatformpoirawpoioutbypoiidsAPIResponseModel is 根据poiId输出飞猪poi数据 成功返回结果
type AlitripplatformpoirawpoioutbypoiidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_poi_raw_poioutbypoiids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripplatformpoirawpoioutbypoiidsResult `json:"result,omitempty" xml:"result,omitempty"`
}
