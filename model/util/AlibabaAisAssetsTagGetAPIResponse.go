package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAisAssetsTagGetAPIResponse 基础设施资产标签获取 API返回值
// alibaba.ais.assets.tag.get
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
type AlibabaAisAssetsTagGetAPIResponse struct {
	model.CommonResponse
	AlibabaAisAssetsTagGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAisAssetsTagGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAisAssetsTagGetAPIResponseModel).Reset()
}

// AlibabaAisAssetsTagGetAPIResponseModel is 基础设施资产标签获取 成功返回结果
type AlibabaAisAssetsTagGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ais_assets_tag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *BaseRep `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAisAssetsTagGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAisAssetsTagGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAisAssetsTagGetAPIResponse)
	},
}

// GetAlibabaAisAssetsTagGetAPIResponse 从 sync.Pool 获取 AlibabaAisAssetsTagGetAPIResponse
func GetAlibabaAisAssetsTagGetAPIResponse() *AlibabaAisAssetsTagGetAPIResponse {
	return poolAlibabaAisAssetsTagGetAPIResponse.Get().(*AlibabaAisAssetsTagGetAPIResponse)
}

// ReleaseAlibabaAisAssetsTagGetAPIResponse 将 AlibabaAisAssetsTagGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAisAssetsTagGetAPIResponse(v *AlibabaAisAssetsTagGetAPIResponse) {
	v.Reset()
	poolAlibabaAisAssetsTagGetAPIResponse.Put(v)
}
