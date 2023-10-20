package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRightBindBackAPIResponse 权限回流 API返回值
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
type AlibabaAlihouseNewhomeRightBindBackAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeRightBindBackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeRightBindBackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeRightBindBackAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeRightBindBackAPIResponseModel is 权限回流 成功返回结果
type AlibabaAlihouseNewhomeRightBindBackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_right_bind_back_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlibabaAlihouseNewhomeRightBindBackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeRightBindBackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeRightBindBackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRightBindBackAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeRightBindBackAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeRightBindBackAPIResponse
func GetAlibabaAlihouseNewhomeRightBindBackAPIResponse() *AlibabaAlihouseNewhomeRightBindBackAPIResponse {
	return poolAlibabaAlihouseNewhomeRightBindBackAPIResponse.Get().(*AlibabaAlihouseNewhomeRightBindBackAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeRightBindBackAPIResponse 将 AlibabaAlihouseNewhomeRightBindBackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRightBindBackAPIResponse(v *AlibabaAlihouseNewhomeRightBindBackAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRightBindBackAPIResponse.Put(v)
}
