package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分类信息 APIResponse
taobao.omniitem.classify.query

通过查询关键字，分页查询分类信息
*/
type TaobaoOmniitemClassifyQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyQueryResponse
}

type TaobaoOmniitemClassifyQueryResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemClassifyQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
