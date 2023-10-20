package xhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldataservicehotelserviceindexAPIResponse 酒店服务指数 API返回值
// taobao.xhotel.data.service.hotel.serviceindex
//
// 酒店服务指数
type TaobaoxhoteldataservicehotelserviceindexAPIResponse struct {
	model.CommonResponse
	TaobaoxhoteldataservicehotelserviceindexAPIResponseModel
}

// TaobaoxhoteldataservicehotelserviceindexAPIResponseModel is 酒店服务指数 成功返回结果
type TaobaoxhoteldataservicehotelserviceindexAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_data_service_hotel_serviceindex_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoxhoteldataservicehotelserviceindexResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
