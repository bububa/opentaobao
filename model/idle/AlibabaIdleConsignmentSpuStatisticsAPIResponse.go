package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentSpuStatisticsAPIResponse 闲鱼帮卖同步服务商交易统计信息 API返回值
// alibaba.idle.consignment.spu.statistics
//
// 闲鱼帮卖同步服务商交易统计信息
type AlibabaIdleConsignmentSpuStatisticsAPIResponse struct {
	model.CommonResponse
	AlibabaIdleConsignmentSpuStatisticsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentSpuStatisticsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleConsignmentSpuStatisticsAPIResponseModel).Reset()
}

// AlibabaIdleConsignmentSpuStatisticsAPIResponseModel is 闲鱼帮卖同步服务商交易统计信息 成功返回结果
type AlibabaIdleConsignmentSpuStatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignment_spu_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentSpuStatisticsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaIdleConsignmentSpuStatisticsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentSpuStatisticsAPIResponse)
	},
}

// GetAlibabaIdleConsignmentSpuStatisticsAPIResponse 从 sync.Pool 获取 AlibabaIdleConsignmentSpuStatisticsAPIResponse
func GetAlibabaIdleConsignmentSpuStatisticsAPIResponse() *AlibabaIdleConsignmentSpuStatisticsAPIResponse {
	return poolAlibabaIdleConsignmentSpuStatisticsAPIResponse.Get().(*AlibabaIdleConsignmentSpuStatisticsAPIResponse)
}

// ReleaseAlibabaIdleConsignmentSpuStatisticsAPIResponse 将 AlibabaIdleConsignmentSpuStatisticsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleConsignmentSpuStatisticsAPIResponse(v *AlibabaIdleConsignmentSpuStatisticsAPIResponse) {
	v.Reset()
	poolAlibabaIdleConsignmentSpuStatisticsAPIResponse.Put(v)
}
