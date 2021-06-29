package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询接口 APIResponse
taobao.qimen.singleitem.query

商品查询接口
*/
type TaobaoQimenSingleitemQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenSingleitemQueryResponse
}

type TaobaoQimenSingleitemQueryResponse struct {
    XMLName xml.Name `xml:"qimen_singleitem_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *ResponseDO `json:"response,omitempty" xml:"response,omitempty"`

    
}
