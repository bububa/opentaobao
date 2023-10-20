package hotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelInfoListGetAPIResponse 酒店详细信息查询 API返回值
// taobao.xhotel.info.list.get
//
// 获取酒店详情信息
type TaobaoXhotelInfoListGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelInfoListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelInfoListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelInfoListGetAPIResponseModel).Reset()
}

// TaobaoXhotelInfoListGetAPIResponseModel is 酒店详细信息查询 成功返回结果
type TaobaoXhotelInfoListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_info_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标准酒店信息
	Hotels []SHotelInfoObject `json:"hotels,omitempty" xml:"hotels>s_hotel_info_object,omitempty"`
	// 酒店总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelInfoListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Hotels = m.Hotels[:0]
	m.TotalResults = 0
}

var poolTaobaoXhotelInfoListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelInfoListGetAPIResponse)
	},
}

// GetTaobaoXhotelInfoListGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelInfoListGetAPIResponse
func GetTaobaoXhotelInfoListGetAPIResponse() *TaobaoXhotelInfoListGetAPIResponse {
	return poolTaobaoXhotelInfoListGetAPIResponse.Get().(*TaobaoXhotelInfoListGetAPIResponse)
}

// ReleaseTaobaoXhotelInfoListGetAPIResponse 将 TaobaoXhotelInfoListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelInfoListGetAPIResponse(v *TaobaoXhotelInfoListGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelInfoListGetAPIResponse.Put(v)
}
