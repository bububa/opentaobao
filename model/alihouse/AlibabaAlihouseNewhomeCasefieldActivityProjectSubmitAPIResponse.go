package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse 案场活动楼盘维护 API返回值
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel is 案场活动楼盘维护 成功返回结果
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_casefield_activity_project_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse
func GetAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse() *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse 将 AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse(v *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse.Put(v)
}
