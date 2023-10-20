package perfect

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatcwmsoutboundpickreceiveAPIResponse 拣货接单 API返回值
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
type AlibabatcwmsoutboundpickreceiveAPIResponse struct {
	model.CommonResponse
	AlibabatcwmsoutboundpickreceiveAPIResponseModel
}

// AlibabatcwmsoutboundpickreceiveAPIResponseModel is 拣货接单 成功返回结果
type AlibabatcwmsoutboundpickreceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_pick_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *PickReceiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}
