package perfect

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPerfectPerformanceLocalitemEditAPIResponse 同城购定制发品编辑 API返回值
// alibaba.perfect.performance.localitem.edit
//
// 同城购业务定制化发品接口，同城购业务线专用
type AlibabaPerfectPerformanceLocalitemEditAPIResponse struct {
	model.CommonResponse
	AlibabaPerfectPerformanceLocalitemEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPerfectPerformanceLocalitemEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPerfectPerformanceLocalitemEditAPIResponseModel).Reset()
}

// AlibabaPerfectPerformanceLocalitemEditAPIResponseModel is 同城购定制发品编辑 成功返回结果
type AlibabaPerfectPerformanceLocalitemEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_perfect_performance_localitem_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPerfectPerformanceLocalitemEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPerfectPerformanceLocalitemEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPerfectPerformanceLocalitemEditAPIResponse)
	},
}

// GetAlibabaPerfectPerformanceLocalitemEditAPIResponse 从 sync.Pool 获取 AlibabaPerfectPerformanceLocalitemEditAPIResponse
func GetAlibabaPerfectPerformanceLocalitemEditAPIResponse() *AlibabaPerfectPerformanceLocalitemEditAPIResponse {
	return poolAlibabaPerfectPerformanceLocalitemEditAPIResponse.Get().(*AlibabaPerfectPerformanceLocalitemEditAPIResponse)
}

// ReleaseAlibabaPerfectPerformanceLocalitemEditAPIResponse 将 AlibabaPerfectPerformanceLocalitemEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPerfectPerformanceLocalitemEditAPIResponse(v *AlibabaPerfectPerformanceLocalitemEditAPIResponse) {
	v.Reset()
	poolAlibabaPerfectPerformanceLocalitemEditAPIResponse.Put(v)
}
