package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询租赁商品信息 APIResponse
alibaba.idle.rent.item.query

查询租赁商品信息
*/
type AlibabaIdleRentItemQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentItemQueryResponse
}

type AlibabaIdleRentItemQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaIdleRentItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
