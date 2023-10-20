package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaStandpointHistorykeyGetAPIResponse 查询历史数据 API返回值
// alibaba.standpoint.historykey.get
//
// 查询历史数据
type AlibabaStandpointHistorykeyGetAPIResponse struct {
	model.CommonResponse
	AlibabaStandpointHistorykeyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaStandpointHistorykeyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaStandpointHistorykeyGetAPIResponseModel).Reset()
}

// AlibabaStandpointHistorykeyGetAPIResponseModel is 查询历史数据 成功返回结果
type AlibabaStandpointHistorykeyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_standpoint_historykey_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据列表
	Content []string `json:"content,omitempty" xml:"content>string,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态吗
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaStandpointHistorykeyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = m.Content[:0]
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.SuccessRes = false
}

var poolAlibabaStandpointHistorykeyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaStandpointHistorykeyGetAPIResponse)
	},
}

// GetAlibabaStandpointHistorykeyGetAPIResponse 从 sync.Pool 获取 AlibabaStandpointHistorykeyGetAPIResponse
func GetAlibabaStandpointHistorykeyGetAPIResponse() *AlibabaStandpointHistorykeyGetAPIResponse {
	return poolAlibabaStandpointHistorykeyGetAPIResponse.Get().(*AlibabaStandpointHistorykeyGetAPIResponse)
}

// ReleaseAlibabaStandpointHistorykeyGetAPIResponse 将 AlibabaStandpointHistorykeyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaStandpointHistorykeyGetAPIResponse(v *AlibabaStandpointHistorykeyGetAPIResponse) {
	v.Reset()
	poolAlibabaStandpointHistorykeyGetAPIResponse.Put(v)
}
