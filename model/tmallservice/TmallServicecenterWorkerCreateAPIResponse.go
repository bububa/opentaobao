package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerCreateAPIResponse 服务商工人信息创建 API返回值
// tmall.servicecenter.worker.create
//
// 服务商工人信息创建
type TmallServicecenterWorkerCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkerCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkerCreateAPIResponseModel).Reset()
}

// TmallServicecenterWorkerCreateAPIResponseModel is 服务商工人信息创建 成功返回结果
type TmallServicecenterWorkerCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_worker_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkerCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkerCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkerCreateAPIResponse)
	},
}

// GetTmallServicecenterWorkerCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkerCreateAPIResponse
func GetTmallServicecenterWorkerCreateAPIResponse() *TmallServicecenterWorkerCreateAPIResponse {
	return poolTmallServicecenterWorkerCreateAPIResponse.Get().(*TmallServicecenterWorkerCreateAPIResponse)
}

// ReleaseTmallServicecenterWorkerCreateAPIResponse 将 TmallServicecenterWorkerCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkerCreateAPIResponse(v *TmallServicecenterWorkerCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkerCreateAPIResponse.Put(v)
}
