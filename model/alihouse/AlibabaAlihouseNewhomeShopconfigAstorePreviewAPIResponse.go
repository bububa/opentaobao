package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse 天猫好房店铺装修-地址预览 API返回值
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
type AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponseModel is 天猫好房店铺装修-地址预览 成功返回结果
type AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_astore_preview_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *AlibabaAlihouseNewhomeShopconfigAstorePreviewResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse
func GetAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse() *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse {
	return poolAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse.Get().(*AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse 将 AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse(v *AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse.Put(v)
}
