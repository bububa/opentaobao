package larkiot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkIotOrderGetcinemasAPIResponse 获取iot渠道开放的影院 API返回值
// taobao.lark.iot.order.getcinemas
//
// iot渠道拉取有权限访问的影院
type TaobaoLarkIotOrderGetcinemasAPIResponse struct {
	model.CommonResponse
	TaobaoLarkIotOrderGetcinemasAPIResponseModel
}

// TaobaoLarkIotOrderGetcinemasAPIResponseModel is 获取iot渠道开放的影院 成功返回结果
type TaobaoLarkIotOrderGetcinemasAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_iot_order_getcinemas_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 有权限的影院列表
	Data *BizListResult `json:"data,omitempty" xml:"data,omitempty"`
}
