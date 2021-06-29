package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务详情模版查询 APIResponse
alitrip.travel.trade.template.query

通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
*/
type AlitripTravelTradeTemplateQueryAPIResponse struct {
    model.CommonResponse
    AlitripTravelTradeTemplateQueryResponse
}

type AlitripTravelTradeTemplateQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_template_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单服务标注模版获取结果
    
    Result   *AlitripTravelTradeTemplateQueryResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
