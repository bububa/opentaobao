package hotel

import (
	"encoding/xml"

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

// TaobaoXhotelInfoListGetAPIResponseModel is 酒店详细信息查询 成功返回结果
type TaobaoXhotelInfoListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_info_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 标准酒店信息
	Hotels []SHotelInfoObject `json:"hotels,omitempty" xml:"hotels>s_hotel_info_object,omitempty"`
}
