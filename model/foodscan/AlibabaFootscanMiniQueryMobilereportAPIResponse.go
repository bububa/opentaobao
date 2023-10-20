package foodscan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniQueryMobilereportAPIResponse 根据scanId查询报告 API返回值
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
type AlibabaFootscanMiniQueryMobilereportAPIResponse struct {
	model.CommonResponse
	AlibabaFootscanMiniQueryMobilereportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFootscanMiniQueryMobilereportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFootscanMiniQueryMobilereportAPIResponseModel).Reset()
}

// AlibabaFootscanMiniQueryMobilereportAPIResponseModel is 根据scanId查询报告 成功返回结果
type AlibabaFootscanMiniQueryMobilereportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_footscan_mini_query_mobilereport_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AlibabaFootscanMiniQueryMobilereportMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFootscanMiniQueryMobilereportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFootscanMiniQueryMobilereportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFootscanMiniQueryMobilereportAPIResponse)
	},
}

// GetAlibabaFootscanMiniQueryMobilereportAPIResponse 从 sync.Pool 获取 AlibabaFootscanMiniQueryMobilereportAPIResponse
func GetAlibabaFootscanMiniQueryMobilereportAPIResponse() *AlibabaFootscanMiniQueryMobilereportAPIResponse {
	return poolAlibabaFootscanMiniQueryMobilereportAPIResponse.Get().(*AlibabaFootscanMiniQueryMobilereportAPIResponse)
}

// ReleaseAlibabaFootscanMiniQueryMobilereportAPIResponse 将 AlibabaFootscanMiniQueryMobilereportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFootscanMiniQueryMobilereportAPIResponse(v *AlibabaFootscanMiniQueryMobilereportAPIResponse) {
	v.Reset()
	poolAlibabaFootscanMiniQueryMobilereportAPIResponse.Put(v)
}
