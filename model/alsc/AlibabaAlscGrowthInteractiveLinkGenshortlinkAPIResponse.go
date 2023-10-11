package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse 短链接口 API返回值
// alibaba.alsc.growth.interactive.link.genshortlink
//
// 短链接口
type AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel
}

// AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel is 短链接口 成功返回结果
type AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_link_genshortlink_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
