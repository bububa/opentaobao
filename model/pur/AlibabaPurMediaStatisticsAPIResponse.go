package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurmediastatisticsAPIResponse 新媒体统计信息 API返回值
// alibaba.pur.media.statistics
//
// 清博同步新媒体的统计信息给到采购平台
type AlibabapurmediastatisticsAPIResponse struct {
	model.CommonResponse
	AlibabapurmediastatisticsAPIResponseModel
}

// AlibabapurmediastatisticsAPIResponseModel is 新媒体统计信息 成功返回结果
type AlibabapurmediastatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_media_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
