package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikePerformancePutAPIResponse 提交客服绩效接口 API返回值
// taobao.weike.performance.put
//
// 提交客服绩效接口
type TaobaoWeikePerformancePutAPIResponse struct {
	model.CommonResponse
	TaobaoWeikePerformancePutAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWeikePerformancePutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWeikePerformancePutAPIResponseModel).Reset()
}

// TaobaoWeikePerformancePutAPIResponseModel is 提交客服绩效接口 成功返回结果
type TaobaoWeikePerformancePutAPIResponseModel struct {
	XMLName xml.Name `xml:"weike_performance_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWeikePerformancePutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoWeikePerformancePutAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWeikePerformancePutAPIResponse)
	},
}

// GetTaobaoWeikePerformancePutAPIResponse 从 sync.Pool 获取 TaobaoWeikePerformancePutAPIResponse
func GetTaobaoWeikePerformancePutAPIResponse() *TaobaoWeikePerformancePutAPIResponse {
	return poolTaobaoWeikePerformancePutAPIResponse.Get().(*TaobaoWeikePerformancePutAPIResponse)
}

// ReleaseTaobaoWeikePerformancePutAPIResponse 将 TaobaoWeikePerformancePutAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWeikePerformancePutAPIResponse(v *TaobaoWeikePerformancePutAPIResponse) {
	v.Reset()
	poolTaobaoWeikePerformancePutAPIResponse.Put(v)
}
