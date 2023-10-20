package hotelalliance

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelalliancehidgetAPIResponse 获取联盟hid API返回值
// alitrip.hotel.alliance.hid.get
//
// 获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
type AlitriphotelalliancehidgetAPIResponse struct {
	model.CommonResponse
	AlitriphotelalliancehidgetAPIResponseModel
}

// AlitriphotelalliancehidgetAPIResponseModel is 获取联盟hid 成功返回结果
type AlitriphotelalliancehidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_alliance_hid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	HmsTopResultSet *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`
}
