package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanChannelGetAPIResponse 5-IBP同步渠道接口 API返回值
// alibaba.tmallgenie.scp.plan.channel.get
//
// IBP同步渠道接口
type AlibabaTmallgenieScpPlanChannelGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanChannelGetAPIResponseModel
}

// AlibabaTmallgenieScpPlanChannelGetAPIResponseModel is 5-IBP同步渠道接口 成功返回结果
type AlibabaTmallgenieScpPlanChannelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_channel_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
