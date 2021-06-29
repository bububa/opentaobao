package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品优惠接口 APIResponse
taobao.singletreasure.activity.item.update

更新单品优惠接口
*/
type TaobaoSingletreasureActivityItemUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemUpdateResponse
}

type TaobaoSingletreasureActivityItemUpdateResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSingletreasureActivityItemUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
