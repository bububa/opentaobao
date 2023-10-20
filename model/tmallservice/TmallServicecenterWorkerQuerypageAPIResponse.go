package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerQuerypageAPIResponse 查询工人列表 API返回值
// tmall.servicecenter.worker.querypage
//
// 服务商查询工人列表
type TmallServicecenterWorkerQuerypageAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkerQuerypageAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerQuerypageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkerQuerypageAPIResponseModel).Reset()
}

// TmallServicecenterWorkerQuerypageAPIResponseModel is 查询工人列表 成功返回结果
type TmallServicecenterWorkerQuerypageAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_worker_querypage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerQuerypageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkerQuerypageAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkerQuerypageAPIResponse)
	},
}

// GetTmallServicecenterWorkerQuerypageAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkerQuerypageAPIResponse
func GetTmallServicecenterWorkerQuerypageAPIResponse() *TmallServicecenterWorkerQuerypageAPIResponse {
	return poolTmallServicecenterWorkerQuerypageAPIResponse.Get().(*TmallServicecenterWorkerQuerypageAPIResponse)
}

// ReleaseTmallServicecenterWorkerQuerypageAPIResponse 将 TmallServicecenterWorkerQuerypageAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkerQuerypageAPIResponse(v *TmallServicecenterWorkerQuerypageAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkerQuerypageAPIResponse.Put(v)
}
