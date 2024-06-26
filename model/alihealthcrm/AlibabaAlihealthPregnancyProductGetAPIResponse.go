package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyProductGetAPIResponse 备孕首页获取达人配置商品 API返回值
// alibaba.alihealth.pregnancy.product.get
//
// 备孕首页获取达人配置商品
type AlibabaAlihealthPregnancyProductGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPregnancyProductGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyProductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPregnancyProductGetAPIResponseModel).Reset()
}

// AlibabaAlihealthPregnancyProductGetAPIResponseModel is 备孕首页获取达人配置商品 成功返回结果
type AlibabaAlihealthPregnancyProductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyProductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPregnancyProductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPregnancyProductGetAPIResponse)
	},
}

// GetAlibabaAlihealthPregnancyProductGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPregnancyProductGetAPIResponse
func GetAlibabaAlihealthPregnancyProductGetAPIResponse() *AlibabaAlihealthPregnancyProductGetAPIResponse {
	return poolAlibabaAlihealthPregnancyProductGetAPIResponse.Get().(*AlibabaAlihealthPregnancyProductGetAPIResponse)
}

// ReleaseAlibabaAlihealthPregnancyProductGetAPIResponse 将 AlibabaAlihealthPregnancyProductGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPregnancyProductGetAPIResponse(v *AlibabaAlihealthPregnancyProductGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPregnancyProductGetAPIResponse.Put(v)
}
