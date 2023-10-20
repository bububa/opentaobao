package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfbusinessareagetAPIResponse 根据城市查询商圈 API返回值
// alitrip.hotel.hstdf.businessarea.get
//
// 根据cityId分页查询商圈信息
type AlitriphotelhstdfbusinessareagetAPIResponse struct {
	model.CommonResponse
	AlitriphotelhstdfbusinessareagetAPIResponseModel
}

// AlitriphotelhstdfbusinessareagetAPIResponseModel is 根据城市查询商圈 成功返回结果
type AlitriphotelhstdfbusinessareagetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_businessarea_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopStdResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
