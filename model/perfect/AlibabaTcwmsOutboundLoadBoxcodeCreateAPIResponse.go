package perfect

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse 创建箱号 API返回值
// alibaba.tcwms.outbound.load.boxcode.create
//
// 创建箱号
type AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel
}

// AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel is 创建箱号 成功返回结果
type AlibabaTcwmsOutboundLoadBoxcodeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_load_boxcode_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BoxCodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
