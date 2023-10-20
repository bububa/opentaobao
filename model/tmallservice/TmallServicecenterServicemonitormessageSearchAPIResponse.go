package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicemonitormessagesearchAPIResponse 根据时间段查询服务商的服务预警消息列表(15分钟内) API返回值
// tmall.servicecenter.servicemonitormessage.search
//
// 根据时间段查询服务商的服务预警消息列表(15分钟内)
type TmallservicecenterservicemonitormessagesearchAPIResponse struct {
	model.CommonResponse
	TmallservicecenterservicemonitormessagesearchAPIResponseModel
}

// TmallservicecenterservicemonitormessagesearchAPIResponseModel is 根据时间段查询服务商的服务预警消息列表(15分钟内) 成功返回结果
type TmallservicecenterservicemonitormessagesearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicemonitormessage_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
