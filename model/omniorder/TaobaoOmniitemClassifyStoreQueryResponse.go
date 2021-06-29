package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据门店查分类信息 APIResponse
taobao.omniitem.classify.store.query

根据门店查分类信息
*/
type TaobaoOmniitemClassifyStoreQueryAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyStoreQueryResponse
}

type TaobaoOmniitemClassifyStoreQueryResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_store_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemClassifyStoreQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
