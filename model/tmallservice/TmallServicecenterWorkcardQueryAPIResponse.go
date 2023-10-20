package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardQueryAPIResponse 工单查询接口 API返回值
// tmall.servicecenter.workcard.query
//
// 工单查询接口
type TmallServicecenterWorkcardQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardQueryAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardQueryAPIResponseModel is 工单查询接口 成功返回结果
type TmallServicecenterWorkcardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *TmallServicecenterWorkcardQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardQueryAPIResponse)
	},
}

// GetTmallServicecenterWorkcardQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardQueryAPIResponse
func GetTmallServicecenterWorkcardQueryAPIResponse() *TmallServicecenterWorkcardQueryAPIResponse {
	return poolTmallServicecenterWorkcardQueryAPIResponse.Get().(*TmallServicecenterWorkcardQueryAPIResponse)
}

// ReleaseTmallServicecenterWorkcardQueryAPIResponse 将 TmallServicecenterWorkcardQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardQueryAPIResponse(v *TmallServicecenterWorkcardQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardQueryAPIResponse.Put(v)
}
