package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallMsfReservationAPIResponse
喵师傅服务预约API API返回值
tmall.msf.reservation

喵师傅预约api */
type TmallMsfReservationAPIResponse struct {
	model.CommonResponse
	TmallMsfReservationAPIResponseModel
}

// TmallMsfReservationAPIResponseModel is 喵师傅服务预约API 成功返回结果
type TmallMsfReservationAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_reservation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 预约成功,json
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
