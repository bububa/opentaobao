package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品优惠接口 APIResponse
taobao.singletreasure.activity.item.delete

删除单品优惠接口
*/
type TaobaoSingletreasureActivityItemDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemDeleteResponse
}

type TaobaoSingletreasureActivityItemDeleteResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSingletreasureActivityItemDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
