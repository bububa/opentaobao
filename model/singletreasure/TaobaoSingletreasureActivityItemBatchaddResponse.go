package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加商品接口 APIResponse
taobao.singletreasure.activity.item.batchadd

向活动中批量添加商品优惠
*/
type TaobaoSingletreasureActivityItemBatchaddAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemBatchaddResponse
}

type TaobaoSingletreasureActivityItemBatchaddResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_batchadd_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSingletreasureActivityItemBatchaddResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
