package hotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDistributionInfoAPIResponse 飞猪分销通用酒店标准信息接口 API返回值
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
type TaobaoXhotelDistributionInfoAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelDistributionInfoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelDistributionInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelDistributionInfoAPIResponseModel).Reset()
}

// TaobaoXhotelDistributionInfoAPIResponseModel is 飞猪分销通用酒店标准信息接口 成功返回结果
type TaobaoXhotelDistributionInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_distribution_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标准酒店信息
	Hotels []SHotelInfoObject `json:"hotels,omitempty" xml:"hotels>s_hotel_info_object,omitempty"`
	// 酒店总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelDistributionInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Hotels = m.Hotels[:0]
	m.TotalResults = 0
}

var poolTaobaoXhotelDistributionInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDistributionInfoAPIResponse)
	},
}

// GetTaobaoXhotelDistributionInfoAPIResponse 从 sync.Pool 获取 TaobaoXhotelDistributionInfoAPIResponse
func GetTaobaoXhotelDistributionInfoAPIResponse() *TaobaoXhotelDistributionInfoAPIResponse {
	return poolTaobaoXhotelDistributionInfoAPIResponse.Get().(*TaobaoXhotelDistributionInfoAPIResponse)
}

// ReleaseTaobaoXhotelDistributionInfoAPIResponse 将 TaobaoXhotelDistributionInfoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelDistributionInfoAPIResponse(v *TaobaoXhotelDistributionInfoAPIResponse) {
	v.Reset()
	poolTaobaoXhotelDistributionInfoAPIResponse.Put(v)
}
