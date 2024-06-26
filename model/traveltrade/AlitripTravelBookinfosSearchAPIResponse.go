package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelBookinfosSearchAPIResponse 飞猪度假-订单预定信息列表搜索接口 API返回值
// alitrip.travel.bookinfos.search
//
// 查询订单预定信息列表
type AlitripTravelBookinfosSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelBookinfosSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelBookinfosSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelBookinfosSearchAPIResponseModel).Reset()
}

// AlitripTravelBookinfosSearchAPIResponseModel is 飞猪度假-订单预定信息列表搜索接口 成功返回结果
type AlitripTravelBookinfosSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_bookinfos_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单及预约id映射列表
	OrderBookInfoList []FirstResult `json:"order_book_info_list,omitempty" xml:"order_book_info_list>first_result,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// true或false
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelBookinfosSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderBookInfoList = m.OrderBookInfoList[:0]
	m.ErrorMsg = ""
	m.IsSuccess = ""
	m.TotalResults = 0
}

var poolAlitripTravelBookinfosSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelBookinfosSearchAPIResponse)
	},
}

// GetAlitripTravelBookinfosSearchAPIResponse 从 sync.Pool 获取 AlitripTravelBookinfosSearchAPIResponse
func GetAlitripTravelBookinfosSearchAPIResponse() *AlitripTravelBookinfosSearchAPIResponse {
	return poolAlitripTravelBookinfosSearchAPIResponse.Get().(*AlitripTravelBookinfosSearchAPIResponse)
}

// ReleaseAlitripTravelBookinfosSearchAPIResponse 将 AlitripTravelBookinfosSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelBookinfosSearchAPIResponse(v *AlitripTravelBookinfosSearchAPIResponse) {
	v.Reset()
	poolAlitripTravelBookinfosSearchAPIResponse.Put(v)
}
