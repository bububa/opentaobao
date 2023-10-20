package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelrefundcancelAPIResponse 渠道退款单撤销 API返回值
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
type AlibabaascpchannelrefundcancelAPIResponse struct {
	model.CommonResponse
	AlibabaascpchannelrefundcancelAPIResponseModel
}

// AlibabaascpchannelrefundcancelAPIResponseModel is 渠道退款单撤销 成功返回结果
type AlibabaascpchannelrefundcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_refund_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
