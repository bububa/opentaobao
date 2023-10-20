package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosPosAlarmAPIResponse pos故障报警 API返回值
// alibaba.mos.pos.alarm
//
// 故障报警
type AlibabaMosPosAlarmAPIResponse struct {
	model.CommonResponse
	AlibabaMosPosAlarmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosPosAlarmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosPosAlarmAPIResponseModel).Reset()
}

// AlibabaMosPosAlarmAPIResponseModel is pos故障报警 成功返回结果
type AlibabaMosPosAlarmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_pos_alarm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errMsg
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errCode
	SubErrCode int64 `json:"sub_err_code,omitempty" xml:"sub_err_code,omitempty"`
	// success/false
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosPosAlarmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessage = ""
	m.SubErrCode = 0
	m.Result = false
}

var poolAlibabaMosPosAlarmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosPosAlarmAPIResponse)
	},
}

// GetAlibabaMosPosAlarmAPIResponse 从 sync.Pool 获取 AlibabaMosPosAlarmAPIResponse
func GetAlibabaMosPosAlarmAPIResponse() *AlibabaMosPosAlarmAPIResponse {
	return poolAlibabaMosPosAlarmAPIResponse.Get().(*AlibabaMosPosAlarmAPIResponse)
}

// ReleaseAlibabaMosPosAlarmAPIResponse 将 AlibabaMosPosAlarmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosPosAlarmAPIResponse(v *AlibabaMosPosAlarmAPIResponse) {
	v.Reset()
	poolAlibabaMosPosAlarmAPIResponse.Put(v)
}
