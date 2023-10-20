package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse 提交小区动态信息 API返回值
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
type AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponseModel is 提交小区动态信息 成功返回结果
type AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_dynamicdata_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeDynamicdataSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse
func GetAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse() *AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse 将 AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse(v *AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse.Put(v)
}
