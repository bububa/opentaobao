package idleparttime

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleParttimeNotifyAPIResponse 兼职通知接口 API返回值
// alibaba.idle.parttime.notify
//
// 服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
type AlibabaIdleParttimeNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaIdleParttimeNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleParttimeNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleParttimeNotifyAPIResponseModel).Reset()
}

// AlibabaIdleParttimeNotifyAPIResponseModel is 兼职通知接口 成功返回结果
type AlibabaIdleParttimeNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_parttime_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleParttimeNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaIdleParttimeNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleParttimeNotifyAPIResponse)
	},
}

// GetAlibabaIdleParttimeNotifyAPIResponse 从 sync.Pool 获取 AlibabaIdleParttimeNotifyAPIResponse
func GetAlibabaIdleParttimeNotifyAPIResponse() *AlibabaIdleParttimeNotifyAPIResponse {
	return poolAlibabaIdleParttimeNotifyAPIResponse.Get().(*AlibabaIdleParttimeNotifyAPIResponse)
}

// ReleaseAlibabaIdleParttimeNotifyAPIResponse 将 AlibabaIdleParttimeNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleParttimeNotifyAPIResponse(v *AlibabaIdleParttimeNotifyAPIResponse) {
	v.Reset()
	poolAlibabaIdleParttimeNotifyAPIResponse.Put(v)
}
