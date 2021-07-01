package xhotelonlineorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripXhotelChannelOrderCreateAPIResponse
渠道分销创建订单接口 API返回值
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供） */
type AlitripXhotelChannelOrderCreateAPIResponse struct {
	model.CommonResponse
	AlitripXhotelChannelOrderCreateAPIResponseModel
}

// AlitripXhotelChannelOrderCreateAPIResponseModel is 渠道分销创建订单接口 成功返回结果
type AlitripXhotelChannelOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_xhotel_channel_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}
