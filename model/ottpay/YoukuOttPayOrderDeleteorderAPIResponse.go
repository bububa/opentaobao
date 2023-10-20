package ottpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderdeleteorderAPIResponse 退订应用中心支付订单 API返回值
// youku.ott.pay.order.deleteorder
//
// 应用中心sdk连续包月退订接口
type YoukuottpayorderdeleteorderAPIResponse struct {
	model.CommonResponse
	YoukuottpayorderdeleteorderAPIResponseModel
}

// YoukuottpayorderdeleteorderAPIResponseModel is 退订应用中心支付订单 成功返回结果
type YoukuottpayorderdeleteorderAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_deleteorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}
