package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardQuerybysellerAPIResponse 工单查询接口（面向商家） API返回值
// tmall.servicecenter.workcard.querybyseller
//
// 查询工单
type TmallServicecenterWorkcardQuerybysellerAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardQuerybysellerAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardQuerybysellerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardQuerybysellerAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardQuerybysellerAPIResponseModel is 工单查询接口（面向商家） 成功返回结果
type TmallServicecenterWorkcardQuerybysellerAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_querybyseller_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallServicecenterWorkcardQuerybysellerResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardQuerybysellerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardQuerybysellerAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardQuerybysellerAPIResponse)
	},
}

// GetTmallServicecenterWorkcardQuerybysellerAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardQuerybysellerAPIResponse
func GetTmallServicecenterWorkcardQuerybysellerAPIResponse() *TmallServicecenterWorkcardQuerybysellerAPIResponse {
	return poolTmallServicecenterWorkcardQuerybysellerAPIResponse.Get().(*TmallServicecenterWorkcardQuerybysellerAPIResponse)
}

// ReleaseTmallServicecenterWorkcardQuerybysellerAPIResponse 将 TmallServicecenterWorkcardQuerybysellerAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardQuerybysellerAPIResponse(v *TmallServicecenterWorkcardQuerybysellerAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardQuerybysellerAPIResponse.Put(v)
}
