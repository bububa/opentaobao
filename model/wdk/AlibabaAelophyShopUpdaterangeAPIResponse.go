package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyShopUpdaterangeAPIResponse 更新渠道店销售范围 API返回值
// alibaba.aelophy.shop.updaterange
//
// 更新渠道店销售范围
type AlibabaAelophyShopUpdaterangeAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyShopUpdaterangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyShopUpdaterangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyShopUpdaterangeAPIResponseModel).Reset()
}

// AlibabaAelophyShopUpdaterangeAPIResponseModel is 更新渠道店销售范围 成功返回结果
type AlibabaAelophyShopUpdaterangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_shop_updaterange_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// api调用结果
	ApiResult *AlibabaAelophyShopUpdaterangeApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyShopUpdaterangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyShopUpdaterangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyShopUpdaterangeAPIResponse)
	},
}

// GetAlibabaAelophyShopUpdaterangeAPIResponse 从 sync.Pool 获取 AlibabaAelophyShopUpdaterangeAPIResponse
func GetAlibabaAelophyShopUpdaterangeAPIResponse() *AlibabaAelophyShopUpdaterangeAPIResponse {
	return poolAlibabaAelophyShopUpdaterangeAPIResponse.Get().(*AlibabaAelophyShopUpdaterangeAPIResponse)
}

// ReleaseAlibabaAelophyShopUpdaterangeAPIResponse 将 AlibabaAelophyShopUpdaterangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyShopUpdaterangeAPIResponse(v *AlibabaAelophyShopUpdaterangeAPIResponse) {
	v.Reset()
	poolAlibabaAelophyShopUpdaterangeAPIResponse.Put(v)
}
