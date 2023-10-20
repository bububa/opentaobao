package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelIncrementInfoGetAPIResponse 酒店状态增量查询接口 API返回值
// taobao.xhotel.increment.info.get
//
// 酒店状态增量查询接口
type TaobaoXhotelIncrementInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelIncrementInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelIncrementInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelIncrementInfoGetAPIResponseModel).Reset()
}

// TaobaoXhotelIncrementInfoGetAPIResponseModel is 酒店状态增量查询接口 成功返回结果
type TaobaoXhotelIncrementInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_increment_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店列表
	Hotels []SHotel `json:"hotels,omitempty" xml:"hotels>s_hotel,omitempty"`
	// 酒店总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelIncrementInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Hotels = m.Hotels[:0]
	m.TotalResults = 0
}

var poolTaobaoXhotelIncrementInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelIncrementInfoGetAPIResponse)
	},
}

// GetTaobaoXhotelIncrementInfoGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelIncrementInfoGetAPIResponse
func GetTaobaoXhotelIncrementInfoGetAPIResponse() *TaobaoXhotelIncrementInfoGetAPIResponse {
	return poolTaobaoXhotelIncrementInfoGetAPIResponse.Get().(*TaobaoXhotelIncrementInfoGetAPIResponse)
}

// ReleaseTaobaoXhotelIncrementInfoGetAPIResponse 将 TaobaoXhotelIncrementInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelIncrementInfoGetAPIResponse(v *TaobaoXhotelIncrementInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelIncrementInfoGetAPIResponse.Put(v)
}
