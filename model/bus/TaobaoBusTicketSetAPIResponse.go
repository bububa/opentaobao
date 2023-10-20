package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusticketsetAPIResponse 出票接口 API返回值
// taobao.bus.ticket.set
//
// 提供给汽车票商家出票使用
type TaobaobusticketsetAPIResponse struct {
	model.CommonResponse
	TaobaobusticketsetAPIResponseModel
}

// TaobaobusticketsetAPIResponseModel is 出票接口 成功返回结果
type TaobaobusticketsetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_ticket_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ErrorCode1 string `json:"error_code1,omitempty" xml:"error_code1,omitempty"`
	// errorMsg
	ErrorMsg1 string `json:"error_msg1,omitempty" xml:"error_msg1,omitempty"`
	// success1
	Success1 bool `json:"success1,omitempty" xml:"success1,omitempty"`
}
