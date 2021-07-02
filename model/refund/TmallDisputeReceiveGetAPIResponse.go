package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallDisputeReceiveGetAPIResponse 天猫逆向纠纷查询 API返回值
// tmall.dispute.receive.get
//
// 展示商家所有退款信息
type TmallDisputeReceiveGetAPIResponse struct {
	model.CommonResponse
	TmallDisputeReceiveGetAPIResponseModel
}

// TmallDisputeReceiveGetAPIResponseModel is 天猫逆向纠纷查询 成功返回结果
type TmallDisputeReceiveGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_dispute_receive_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallDisputeReceiveGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
