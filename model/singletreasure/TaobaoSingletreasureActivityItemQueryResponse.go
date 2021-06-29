package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动下的优惠信息 APIResponse
taobao.singletreasure.activity.item.query

分页查询活动下的商品优惠信息
*/
type TaobaoSingletreasureActivityItemQueryAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemQueryResponse
}

type TaobaoSingletreasureActivityItemQueryResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *PageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
