package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse 插入和更新食谱 API返回值
// alibaba.ailabs.iot.business.recipe.insertorupdate
//
// 插入和更新食谱，将isv的食谱添加到云端进行存储
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponseModel).Reset()
}

// AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponseModel is 插入和更新食谱 成功返回结果
type AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipe_insertorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 响应code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 返回结果，行业食谱Id
	RetValue int64 `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TraceId = ""
	m.RetCode = 0
	m.RetValue = 0
}

var poolAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse)
	},
}

// GetAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse
func GetAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse() *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse {
	return poolAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse.Get().(*AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse)
}

// ReleaseAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse 将 AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse(v *AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse.Put(v)
}
