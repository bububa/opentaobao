package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusRefundSetAPIResponse B2B退票申请接口 API返回值
// taobao.bus.refund.set
//
// B2B业务支持退票
type TaobaoBusRefundSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusRefundSetAPIResponseModel
}

// TaobaoBusRefundSetAPIResponseModel is B2B退票申请接口 成功返回结果
type TaobaoBusRefundSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_refund_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *B2BRefundOrderRp `json:"result,omitempty" xml:"result,omitempty"`
}
