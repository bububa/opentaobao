package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardDesensitizationQueryAPIResponse 工单查询接口 API返回值
// tmall.servicecenter.workcard.desensitization.query
//
// 工单查询接口
type TmallServicecenterWorkcardDesensitizationQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardDesensitizationQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardDesensitizationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardDesensitizationQueryAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardDesensitizationQueryAPIResponseModel is 工单查询接口 成功返回结果
type TmallServicecenterWorkcardDesensitizationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_desensitization_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *TmallServicecenterWorkcardDesensitizationQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardDesensitizationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardDesensitizationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardDesensitizationQueryAPIResponse)
	},
}

// GetTmallServicecenterWorkcardDesensitizationQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardDesensitizationQueryAPIResponse
func GetTmallServicecenterWorkcardDesensitizationQueryAPIResponse() *TmallServicecenterWorkcardDesensitizationQueryAPIResponse {
	return poolTmallServicecenterWorkcardDesensitizationQueryAPIResponse.Get().(*TmallServicecenterWorkcardDesensitizationQueryAPIResponse)
}

// ReleaseTmallServicecenterWorkcardDesensitizationQueryAPIResponse 将 TmallServicecenterWorkcardDesensitizationQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardDesensitizationQueryAPIResponse(v *TmallServicecenterWorkcardDesensitizationQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardDesensitizationQueryAPIResponse.Put(v)
}
