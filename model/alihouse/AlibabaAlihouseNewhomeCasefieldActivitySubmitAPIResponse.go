package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse 案场活动维护 API返回值
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
type AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponseModel is 案场活动维护 成功返回结果
type AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_casefield_activity_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabaAlihouseNewhomeCasefieldActivitySubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse
func GetAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse() *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse 将 AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse(v *AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse.Put(v)
}
