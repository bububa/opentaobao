package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单是否由裹裹发送消息 APIResponse
cainiao.endpoint.locker.top.order.noticesend.query

合作公司查询消息发送的接口，判断是否裹裹发送消息
*/
type CainiaoEndpointLockerTopOrderNoticesendQueryAPIResponse struct {
    model.CommonResponse
    CainiaoEndpointLockerTopOrderNoticesendQueryResponse
}

type CainiaoEndpointLockerTopOrderNoticesendQueryResponse struct {
    XMLName xml.Name `xml:"cainiao_endpoint_locker_top_order_noticesend_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
