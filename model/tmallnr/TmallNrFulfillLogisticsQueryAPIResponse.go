package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillLogisticsQueryAPIResponse 定时送和极速达配送物流信息查询 API返回值
// tmall.nr.fulfill.logistics.query
//
// 发布定时送&amp;极速达物流信息查询服务
type TmallNrFulfillLogisticsQueryAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillLogisticsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillLogisticsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillLogisticsQueryAPIResponseModel).Reset()
}

// TmallNrFulfillLogisticsQueryAPIResponseModel is 定时送和极速达配送物流信息查询 成功返回结果
type TmallNrFulfillLogisticsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillLogisticsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillLogisticsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillLogisticsQueryAPIResponse)
	},
}

// GetTmallNrFulfillLogisticsQueryAPIResponse 从 sync.Pool 获取 TmallNrFulfillLogisticsQueryAPIResponse
func GetTmallNrFulfillLogisticsQueryAPIResponse() *TmallNrFulfillLogisticsQueryAPIResponse {
	return poolTmallNrFulfillLogisticsQueryAPIResponse.Get().(*TmallNrFulfillLogisticsQueryAPIResponse)
}

// ReleaseTmallNrFulfillLogisticsQueryAPIResponse 将 TmallNrFulfillLogisticsQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillLogisticsQueryAPIResponse(v *TmallNrFulfillLogisticsQueryAPIResponse) {
	v.Reset()
	poolTmallNrFulfillLogisticsQueryAPIResponse.Put(v)
}
