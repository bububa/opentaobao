package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道销售单创建接口 API返回值 
alibaba.ascp.channel.sales.order.create

阿里巴巴供应链渠道销售订单创建接口
*/
type AlibabaAscpChannelSalesOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpChannelSalesOrderCreateAPIResponseModel
}

// 供应链渠道销售单创建接口 成功返回结果
type AlibabaAscpChannelSalesOrderCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ascp_channel_sales_order_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值包装,result为返回具体消息内容
    CreateOrderResponse   *ResultWrapper `json:"create_order_response,omitempty" xml:"create_order_response,omitempty"`
}
