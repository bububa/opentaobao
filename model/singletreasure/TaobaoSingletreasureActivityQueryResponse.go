package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表接口 APIResponse
taobao.singletreasure.activity.query

查询活动列表接口
*/
type TaobaoSingletreasureActivityQueryAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityQueryResponse
}

type TaobaoSingletreasureActivityQueryResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoSingletreasureActivityQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
