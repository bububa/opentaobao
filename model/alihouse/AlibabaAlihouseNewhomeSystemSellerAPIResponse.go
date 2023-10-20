package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeSystemSellerAPIResponse 商品发布授权 API返回值
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
type AlibabaAlihouseNewhomeSystemSellerAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeSystemSellerAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeSystemSellerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeSystemSellerAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeSystemSellerAPIResponseModel is 商品发布授权 成功返回结果
type AlibabaAlihouseNewhomeSystemSellerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_system_seller_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeSystemSellerResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeSystemSellerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeSystemSellerAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeSystemSellerAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeSystemSellerAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeSystemSellerAPIResponse
func GetAlibabaAlihouseNewhomeSystemSellerAPIResponse() *AlibabaAlihouseNewhomeSystemSellerAPIResponse {
	return poolAlibabaAlihouseNewhomeSystemSellerAPIResponse.Get().(*AlibabaAlihouseNewhomeSystemSellerAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeSystemSellerAPIResponse 将 AlibabaAlihouseNewhomeSystemSellerAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeSystemSellerAPIResponse(v *AlibabaAlihouseNewhomeSystemSellerAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeSystemSellerAPIResponse.Put(v)
}
