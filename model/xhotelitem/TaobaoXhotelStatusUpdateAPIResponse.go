package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelStatusUpdateAPIResponse 酒店状态更新 API返回值
// taobao.xhotel.status.update
//
// 酒店状态更新
type TaobaoXhotelStatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelStatusUpdateAPIResponseModel
}

// TaobaoXhotelStatusUpdateAPIResponseModel is 酒店状态更新 成功返回结果
type TaobaoXhotelStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否出错
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
