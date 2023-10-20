package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipeGetpageAPIResponse 分页查询食谱 API返回值
// alibaba.ailabs.iot.business.recipe.getpage
//
// 分页查询食谱数据
type AlibabaAilabsIotBusinessRecipeGetpageAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotBusinessRecipeGetpageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsIotBusinessRecipeGetpageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsIotBusinessRecipeGetpageAPIResponseModel).Reset()
}

// AlibabaAilabsIotBusinessRecipeGetpageAPIResponseModel is 分页查询食谱 成功返回结果
type AlibabaAilabsIotBusinessRecipeGetpageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipe_getpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsIotBusinessRecipeGetpageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsIotBusinessRecipeGetpageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsIotBusinessRecipeGetpageAPIResponse)
	},
}

// GetAlibabaAilabsIotBusinessRecipeGetpageAPIResponse 从 sync.Pool 获取 AlibabaAilabsIotBusinessRecipeGetpageAPIResponse
func GetAlibabaAilabsIotBusinessRecipeGetpageAPIResponse() *AlibabaAilabsIotBusinessRecipeGetpageAPIResponse {
	return poolAlibabaAilabsIotBusinessRecipeGetpageAPIResponse.Get().(*AlibabaAilabsIotBusinessRecipeGetpageAPIResponse)
}

// ReleaseAlibabaAilabsIotBusinessRecipeGetpageAPIResponse 将 AlibabaAilabsIotBusinessRecipeGetpageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsIotBusinessRecipeGetpageAPIResponse(v *AlibabaAilabsIotBusinessRecipeGetpageAPIResponse) {
	v.Reset()
	poolAlibabaAilabsIotBusinessRecipeGetpageAPIResponse.Put(v)
}
