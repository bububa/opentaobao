package xhotelonlineorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripxhotelchannelnotifyAPIResponse 分销渠道各类通知接口 API返回值
// alitrip.xhotel.channel.notify
//
// 分销渠道支付通知
type AlitripxhotelchannelnotifyAPIResponse struct {
	model.CommonResponse
	AlitripxhotelchannelnotifyAPIResponseModel
}

// AlitripxhotelchannelnotifyAPIResponseModel is 分销渠道各类通知接口 成功返回结果
type AlitripxhotelchannelnotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_xhotel_channel_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}
