package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelexportsroomtypeAPIResponse 导出一个卖家房型下的所有标准房型 API返回值
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
type AlitriphotelhstdfshotelexportsroomtypeAPIResponse struct {
	model.CommonResponse
	AlitriphotelhstdfshotelexportsroomtypeAPIResponseModel
}

// AlitriphotelhstdfshotelexportsroomtypeAPIResponseModel is 导出一个卖家房型下的所有标准房型 成功返回结果
type AlitriphotelhstdfshotelexportsroomtypeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exportsroomtype_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
