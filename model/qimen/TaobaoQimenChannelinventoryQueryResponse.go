package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 APIResponse
taobao.qimen.channelinventory.query

渠道库存查询
*/
type TaobaoQimenChannelinventoryQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenChannelinventoryQueryResponse
}

type TaobaoQimenChannelinventoryQueryResponse struct {
    XMLName xml.Name `xml:"qimen_channelinventory_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *ResponseDO `json:"response,omitempty" xml:"response,omitempty"`

    
}
