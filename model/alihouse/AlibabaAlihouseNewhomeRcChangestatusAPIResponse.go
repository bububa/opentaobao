package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRcChangestatusAPIResponse 图文草稿状态更新 API返回值
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
type AlibabaAlihouseNewhomeRcChangestatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeRcChangestatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeRcChangestatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeRcChangestatusAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeRcChangestatusAPIResponseModel is 图文草稿状态更新 成功返回结果
type AlibabaAlihouseNewhomeRcChangestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_rc_changestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeRcChangestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeRcChangestatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeRcChangestatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRcChangestatusAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeRcChangestatusAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeRcChangestatusAPIResponse
func GetAlibabaAlihouseNewhomeRcChangestatusAPIResponse() *AlibabaAlihouseNewhomeRcChangestatusAPIResponse {
	return poolAlibabaAlihouseNewhomeRcChangestatusAPIResponse.Get().(*AlibabaAlihouseNewhomeRcChangestatusAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeRcChangestatusAPIResponse 将 AlibabaAlihouseNewhomeRcChangestatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRcChangestatusAPIResponse(v *AlibabaAlihouseNewhomeRcChangestatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRcChangestatusAPIResponse.Put(v)
}
