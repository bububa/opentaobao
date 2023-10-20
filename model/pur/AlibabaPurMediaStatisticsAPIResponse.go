package pur

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurMediaStatisticsAPIResponse 新媒体统计信息 API返回值
// alibaba.pur.media.statistics
//
// 清博同步新媒体的统计信息给到采购平台
type AlibabaPurMediaStatisticsAPIResponse struct {
	model.CommonResponse
	AlibabaPurMediaStatisticsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPurMediaStatisticsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPurMediaStatisticsAPIResponseModel).Reset()
}

// AlibabaPurMediaStatisticsAPIResponseModel is 新媒体统计信息 成功返回结果
type AlibabaPurMediaStatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_media_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPurMediaStatisticsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPurMediaStatisticsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPurMediaStatisticsAPIResponse)
	},
}

// GetAlibabaPurMediaStatisticsAPIResponse 从 sync.Pool 获取 AlibabaPurMediaStatisticsAPIResponse
func GetAlibabaPurMediaStatisticsAPIResponse() *AlibabaPurMediaStatisticsAPIResponse {
	return poolAlibabaPurMediaStatisticsAPIResponse.Get().(*AlibabaPurMediaStatisticsAPIResponse)
}

// ReleaseAlibabaPurMediaStatisticsAPIResponse 将 AlibabaPurMediaStatisticsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPurMediaStatisticsAPIResponse(v *AlibabaPurMediaStatisticsAPIResponse) {
	v.Reset()
	poolAlibabaPurMediaStatisticsAPIResponse.Put(v)
}
