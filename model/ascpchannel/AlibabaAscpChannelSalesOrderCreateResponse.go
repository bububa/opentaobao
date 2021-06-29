package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链渠道销售单创建接口 APIResponse
alibaba.ascp.channel.sales.order.create

阿里巴巴供应链渠道销售订单创建接口
*/
type AlibabaAscpChannelSalesOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAscpChannelSalesOrderCreateResponse
}

type AlibabaAscpChannelSalesOrderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_channel_sales_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值包装,result为返回具体消息内容
    
    CreateOrderResponse   *ResultWrapper `json:"create_order_response,omitempty" xml:"create_order_response,omitempty"`

    
}
