package idleparttime

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleParttimeJobsyncAPIResponse 兼职岗位同步 API返回值
// alibaba.idle.parttime.jobsync
//
// 服务商同步岗位信息给闲鱼
type AlibabaIdleParttimeJobsyncAPIResponse struct {
	model.CommonResponse
	AlibabaIdleParttimeJobsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleParttimeJobsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleParttimeJobsyncAPIResponseModel).Reset()
}

// AlibabaIdleParttimeJobsyncAPIResponseModel is 兼职岗位同步 成功返回结果
type AlibabaIdleParttimeJobsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_parttime_jobsync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleParttimeJobsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaIdleParttimeJobsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleParttimeJobsyncAPIResponse)
	},
}

// GetAlibabaIdleParttimeJobsyncAPIResponse 从 sync.Pool 获取 AlibabaIdleParttimeJobsyncAPIResponse
func GetAlibabaIdleParttimeJobsyncAPIResponse() *AlibabaIdleParttimeJobsyncAPIResponse {
	return poolAlibabaIdleParttimeJobsyncAPIResponse.Get().(*AlibabaIdleParttimeJobsyncAPIResponse)
}

// ReleaseAlibabaIdleParttimeJobsyncAPIResponse 将 AlibabaIdleParttimeJobsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleParttimeJobsyncAPIResponse(v *AlibabaIdleParttimeJobsyncAPIResponse) {
	v.Reset()
	poolAlibabaIdleParttimeJobsyncAPIResponse.Put(v)
}
