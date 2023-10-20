package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsalesorderbeforecheckAPIResponse 供应链外部订单创建前校验接口 API返回值
// alibaba.ascp.channel.sales.order.before.check
//
// 供应链外部订单创建前校验接口
type AlibabaascpchannelsalesorderbeforecheckAPIResponse struct {
	model.CommonResponse
	AlibabaascpchannelsalesorderbeforecheckAPIResponseModel
}

// AlibabaascpchannelsalesorderbeforecheckAPIResponseModel is 供应链外部订单创建前校验接口 成功返回结果
type AlibabaascpchannelsalesorderbeforecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_sales_order_before_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	OrderCheckResponse *ResultWrapper `json:"order_check_response,omitempty" xml:"order_check_response,omitempty"`
}
