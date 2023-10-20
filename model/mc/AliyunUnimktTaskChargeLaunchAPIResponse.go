package mc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunUnimktTaskChargeLaunchAPIResponse 云码权益查询 API返回值
// aliyun.unimkt.task.charge.launch
//
// 云码线上流量投放链路，用于判断用户是否有匹配的投放计划
type AliyunUnimktTaskChargeLaunchAPIResponse struct {
	model.CommonResponse
	AliyunUnimktTaskChargeLaunchAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunUnimktTaskChargeLaunchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunUnimktTaskChargeLaunchAPIResponseModel).Reset()
}

// AliyunUnimktTaskChargeLaunchAPIResponseModel is 云码权益查询 成功返回结果
type AliyunUnimktTaskChargeLaunchAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_unimkt_task_charge_launch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误码
	TaskErrorCode int64 `json:"task_error_code,omitempty" xml:"task_error_code,omitempty"`
	// 执行结果
	TaskSuccess bool `json:"task_success,omitempty" xml:"task_success,omitempty"`
}

// Reset 清空结构体
func (m *AliyunUnimktTaskChargeLaunchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = ""
	m.TaskErrorCode = 0
	m.TaskSuccess = false
}

var poolAliyunUnimktTaskChargeLaunchAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunUnimktTaskChargeLaunchAPIResponse)
	},
}

// GetAliyunUnimktTaskChargeLaunchAPIResponse 从 sync.Pool 获取 AliyunUnimktTaskChargeLaunchAPIResponse
func GetAliyunUnimktTaskChargeLaunchAPIResponse() *AliyunUnimktTaskChargeLaunchAPIResponse {
	return poolAliyunUnimktTaskChargeLaunchAPIResponse.Get().(*AliyunUnimktTaskChargeLaunchAPIResponse)
}

// ReleaseAliyunUnimktTaskChargeLaunchAPIResponse 将 AliyunUnimktTaskChargeLaunchAPIResponse 保存到 sync.Pool
func ReleaseAliyunUnimktTaskChargeLaunchAPIResponse(v *AliyunUnimktTaskChargeLaunchAPIResponse) {
	v.Reset()
	poolAliyunUnimktTaskChargeLaunchAPIResponse.Put(v)
}
