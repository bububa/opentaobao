package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaorefundrefundactionsdisplayAPIResponse 退货退款操作的展示信息(展现给买家) API返回值
// cainiao.refund.refundactions.display
//
// 退货退款操作的展示信息(展现给买家)
type CainiaorefundrefundactionsdisplayAPIResponse struct {
	model.CommonResponse
	CainiaorefundrefundactionsdisplayAPIResponseModel
}

// CainiaorefundrefundactionsdisplayAPIResponseModel is 退货退款操作的展示信息(展现给买家) 成功返回结果
type CainiaorefundrefundactionsdisplayAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_refund_refundactions_display_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *CainiaorefundrefundactionsdisplayBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
