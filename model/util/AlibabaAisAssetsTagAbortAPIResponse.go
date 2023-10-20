package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAisAssetsTagAbortAPIResponse 基础设施资产标签废弃 API返回值
// alibaba.ais.assets.tag.abort
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
type AlibabaAisAssetsTagAbortAPIResponse struct {
	model.CommonResponse
	AlibabaAisAssetsTagAbortAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAisAssetsTagAbortAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAisAssetsTagAbortAPIResponseModel).Reset()
}

// AlibabaAisAssetsTagAbortAPIResponseModel is 基础设施资产标签废弃 成功返回结果
type AlibabaAisAssetsTagAbortAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ais_assets_tag_abort_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *BaseRep `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAisAssetsTagAbortAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAisAssetsTagAbortAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAisAssetsTagAbortAPIResponse)
	},
}

// GetAlibabaAisAssetsTagAbortAPIResponse 从 sync.Pool 获取 AlibabaAisAssetsTagAbortAPIResponse
func GetAlibabaAisAssetsTagAbortAPIResponse() *AlibabaAisAssetsTagAbortAPIResponse {
	return poolAlibabaAisAssetsTagAbortAPIResponse.Get().(*AlibabaAisAssetsTagAbortAPIResponse)
}

// ReleaseAlibabaAisAssetsTagAbortAPIResponse 将 AlibabaAisAssetsTagAbortAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAisAssetsTagAbortAPIResponse(v *AlibabaAisAssetsTagAbortAPIResponse) {
	v.Reset()
	poolAlibabaAisAssetsTagAbortAPIResponse.Put(v)
}
