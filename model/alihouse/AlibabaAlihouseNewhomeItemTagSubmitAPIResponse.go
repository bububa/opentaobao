package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeItemTagSubmitAPIResponse ETC上翻商品打标接口 API返回值
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
type AlibabaAlihouseNewhomeItemTagSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeItemTagSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeItemTagSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeItemTagSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeItemTagSubmitAPIResponseModel is ETC上翻商品打标接口 成功返回结果
type AlibabaAlihouseNewhomeItemTagSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_item_tag_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseNewhomeItemTagSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeItemTagSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeItemTagSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeItemTagSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeItemTagSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeItemTagSubmitAPIResponse
func GetAlibabaAlihouseNewhomeItemTagSubmitAPIResponse() *AlibabaAlihouseNewhomeItemTagSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeItemTagSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeItemTagSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeItemTagSubmitAPIResponse 将 AlibabaAlihouseNewhomeItemTagSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeItemTagSubmitAPIResponse(v *AlibabaAlihouseNewhomeItemTagSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeItemTagSubmitAPIResponse.Put(v)
}
