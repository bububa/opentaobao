package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCommunitySubmitAPIResponse 提交小区信息 API返回值
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
type AlibabaAlihouseNewhomeCommunitySubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeCommunitySubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCommunitySubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeCommunitySubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeCommunitySubmitAPIResponseModel is 提交小区信息 成功返回结果
type AlibabaAlihouseNewhomeCommunitySubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_community_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeCommunitySubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCommunitySubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeCommunitySubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCommunitySubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeCommunitySubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeCommunitySubmitAPIResponse
func GetAlibabaAlihouseNewhomeCommunitySubmitAPIResponse() *AlibabaAlihouseNewhomeCommunitySubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeCommunitySubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeCommunitySubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeCommunitySubmitAPIResponse 将 AlibabaAlihouseNewhomeCommunitySubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCommunitySubmitAPIResponse(v *AlibabaAlihouseNewhomeCommunitySubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCommunitySubmitAPIResponse.Put(v)
}
