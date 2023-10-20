package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAisAssetsTagGenerateAPIResponse 基础设施资产标签生成 API返回值
// alibaba.ais.assets.tag.generate
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
type AlibabaAisAssetsTagGenerateAPIResponse struct {
	model.CommonResponse
	AlibabaAisAssetsTagGenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAisAssetsTagGenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAisAssetsTagGenerateAPIResponseModel).Reset()
}

// AlibabaAisAssetsTagGenerateAPIResponseModel is 基础设施资产标签生成 成功返回结果
type AlibabaAisAssetsTagGenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ais_assets_tag_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *BaseRep `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAisAssetsTagGenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAisAssetsTagGenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAisAssetsTagGenerateAPIResponse)
	},
}

// GetAlibabaAisAssetsTagGenerateAPIResponse 从 sync.Pool 获取 AlibabaAisAssetsTagGenerateAPIResponse
func GetAlibabaAisAssetsTagGenerateAPIResponse() *AlibabaAisAssetsTagGenerateAPIResponse {
	return poolAlibabaAisAssetsTagGenerateAPIResponse.Get().(*AlibabaAisAssetsTagGenerateAPIResponse)
}

// ReleaseAlibabaAisAssetsTagGenerateAPIResponse 将 AlibabaAisAssetsTagGenerateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAisAssetsTagGenerateAPIResponse(v *AlibabaAisAssetsTagGenerateAPIResponse) {
	v.Reset()
	poolAlibabaAisAssetsTagGenerateAPIResponse.Put(v)
}
