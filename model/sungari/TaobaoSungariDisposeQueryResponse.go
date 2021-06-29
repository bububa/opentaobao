package sungari

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品商家处置结果查询 APIResponse
taobao.sungari.dispose.query

红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
*/
type TaobaoSungariDisposeQueryAPIResponse struct {
    model.CommonResponse
    TaobaoSungariDisposeQueryResponse
}

type TaobaoSungariDisposeQueryResponse struct {
    XMLName xml.Name `xml:"sungari_dispose_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSungariDisposeQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
