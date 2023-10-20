package xhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripxhotelchannelordercreateresqueryAPIResponse 分销订单查询订单创建结果 API返回值
// alitrip.xhotel.channel.order.create.res.query
//
// 针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
type AlitripxhotelchannelordercreateresqueryAPIResponse struct {
	model.CommonResponse
	AlitripxhotelchannelordercreateresqueryAPIResponseModel
}

// AlitripxhotelchannelordercreateresqueryAPIResponseModel is 分销订单查询订单创建结果 成功返回结果
type AlitripxhotelchannelordercreateresqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_xhotel_channel_order_create_res_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}
