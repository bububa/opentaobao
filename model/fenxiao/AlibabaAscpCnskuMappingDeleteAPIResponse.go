package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuMappingDeleteAPIResponse 货品关系解绑 API返回值
// alibaba.ascp.cnsku.mapping.delete
//
// 货品关系解绑
type AlibabaAscpCnskuMappingDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAscpCnskuMappingDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpCnskuMappingDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpCnskuMappingDeleteAPIResponseModel).Reset()
}

// AlibabaAscpCnskuMappingDeleteAPIResponseModel is 货品关系解绑 成功返回结果
type AlibabaAscpCnskuMappingDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_cnsku_mapping_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功的关联关系集合
	SuccessResultMap string `json:"success_result_map,omitempty" xml:"success_result_map,omitempty"`
	// 错误编码
	EroCode string `json:"ero_code,omitempty" xml:"ero_code,omitempty"`
	// 失败的关联关系集合
	FailResultMap string `json:"fail_result_map,omitempty" xml:"fail_result_map,omitempty"`
	// 错误信息
	EroMsg string `json:"ero_msg,omitempty" xml:"ero_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 是否系统异常
	IsSystemFailed bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpCnskuMappingDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SuccessResultMap = ""
	m.EroCode = ""
	m.FailResultMap = ""
	m.EroMsg = ""
	m.IsSuccess = false
	m.IsSystemFailed = false
}

var poolAlibabaAscpCnskuMappingDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpCnskuMappingDeleteAPIResponse)
	},
}

// GetAlibabaAscpCnskuMappingDeleteAPIResponse 从 sync.Pool 获取 AlibabaAscpCnskuMappingDeleteAPIResponse
func GetAlibabaAscpCnskuMappingDeleteAPIResponse() *AlibabaAscpCnskuMappingDeleteAPIResponse {
	return poolAlibabaAscpCnskuMappingDeleteAPIResponse.Get().(*AlibabaAscpCnskuMappingDeleteAPIResponse)
}

// ReleaseAlibabaAscpCnskuMappingDeleteAPIResponse 将 AlibabaAscpCnskuMappingDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpCnskuMappingDeleteAPIResponse(v *AlibabaAscpCnskuMappingDeleteAPIResponse) {
	v.Reset()
	poolAlibabaAscpCnskuMappingDeleteAPIResponse.Put(v)
}
