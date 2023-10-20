package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCollectAPIResponse 工单揽件 API返回值
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
type TmallServicecenterWorkcardCollectAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardCollectAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardCollectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardCollectAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardCollectAPIResponseModel is 工单揽件 成功返回结果
type TmallServicecenterWorkcardCollectAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_collect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardCollectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardCollectAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardCollectAPIResponse)
	},
}

// GetTmallServicecenterWorkcardCollectAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardCollectAPIResponse
func GetTmallServicecenterWorkcardCollectAPIResponse() *TmallServicecenterWorkcardCollectAPIResponse {
	return poolTmallServicecenterWorkcardCollectAPIResponse.Get().(*TmallServicecenterWorkcardCollectAPIResponse)
}

// ReleaseTmallServicecenterWorkcardCollectAPIResponse 将 TmallServicecenterWorkcardCollectAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardCollectAPIResponse(v *TmallServicecenterWorkcardCollectAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardCollectAPIResponse.Put(v)
}
