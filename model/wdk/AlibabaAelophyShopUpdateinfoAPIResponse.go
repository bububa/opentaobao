package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyShopUpdateinfoAPIResponse 更新渠道店基础信息 API返回值
// alibaba.aelophy.shop.updateinfo
//
// 更新渠道店基础信息
type AlibabaAelophyShopUpdateinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyShopUpdateinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyShopUpdateinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyShopUpdateinfoAPIResponseModel).Reset()
}

// AlibabaAelophyShopUpdateinfoAPIResponseModel is 更新渠道店基础信息 成功返回结果
type AlibabaAelophyShopUpdateinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_shop_updateinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// api调用结果
	ApiResult *AlibabaAelophyShopUpdateinfoApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyShopUpdateinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyShopUpdateinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyShopUpdateinfoAPIResponse)
	},
}

// GetAlibabaAelophyShopUpdateinfoAPIResponse 从 sync.Pool 获取 AlibabaAelophyShopUpdateinfoAPIResponse
func GetAlibabaAelophyShopUpdateinfoAPIResponse() *AlibabaAelophyShopUpdateinfoAPIResponse {
	return poolAlibabaAelophyShopUpdateinfoAPIResponse.Get().(*AlibabaAelophyShopUpdateinfoAPIResponse)
}

// ReleaseAlibabaAelophyShopUpdateinfoAPIResponse 将 AlibabaAelophyShopUpdateinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyShopUpdateinfoAPIResponse(v *AlibabaAelophyShopUpdateinfoAPIResponse) {
	v.Reset()
	poolAlibabaAelophyShopUpdateinfoAPIResponse.Put(v)
}
