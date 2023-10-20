package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuUpdateAPIResponse 供应链中台货品修改接口 API返回值
// alibaba.ascp.cnsku.update
//
// 供应链中台货品修改接口
type AlibabaAscpCnskuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpCnskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpCnskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpCnskuUpdateAPIResponseModel).Reset()
}

// AlibabaAscpCnskuUpdateAPIResponseModel is 供应链中台货品修改接口 成功返回结果
type AlibabaAscpCnskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_cnsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常信息
	ErrorMessages []string `json:"error_messages,omitempty" xml:"error_messages>string,omitempty"`
	// 货品id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 异常信息Code
	SysErrorCode string `json:"sys_error_code,omitempty" xml:"sys_error_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 是否系统异常
	IsSystemFailed bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpCnskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessages = m.ErrorMessages[:0]
	m.Data = ""
	m.SysErrorCode = ""
	m.IsSuccess = false
	m.IsSystemFailed = false
}

var poolAlibabaAscpCnskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpCnskuUpdateAPIResponse)
	},
}

// GetAlibabaAscpCnskuUpdateAPIResponse 从 sync.Pool 获取 AlibabaAscpCnskuUpdateAPIResponse
func GetAlibabaAscpCnskuUpdateAPIResponse() *AlibabaAscpCnskuUpdateAPIResponse {
	return poolAlibabaAscpCnskuUpdateAPIResponse.Get().(*AlibabaAscpCnskuUpdateAPIResponse)
}

// ReleaseAlibabaAscpCnskuUpdateAPIResponse 将 AlibabaAscpCnskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpCnskuUpdateAPIResponse(v *AlibabaAscpCnskuUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAscpCnskuUpdateAPIResponse.Put(v)
}
