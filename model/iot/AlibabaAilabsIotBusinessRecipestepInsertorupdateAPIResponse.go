package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse 插入或更新食谱步骤 API返回值
// alibaba.ailabs.iot.business.recipestep.insertorupdate
//
// 插入或更新食谱步骤
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponseModel).Reset()
}

// AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponseModel is 插入或更新食谱步骤 成功返回结果
type AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipestep_insertorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 响应code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 返回结果
	RetValue int64 `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TraceId = ""
	m.RetCode = 0
	m.RetValue = 0
}

var poolAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse)
	},
}

// GetAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse
func GetAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse() *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse {
	return poolAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse.Get().(*AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse)
}

// ReleaseAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse 将 AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse(v *AlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotBusinessRecipestepInsertorupdateAPIResponse.Put(v)
}
