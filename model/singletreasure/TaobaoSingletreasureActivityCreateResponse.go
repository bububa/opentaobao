package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动创建接口 APIResponse
taobao.singletreasure.activity.create

创建优惠活动
*/
type TaobaoSingletreasureActivityCreateAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityCreateResponse
}

type TaobaoSingletreasureActivityCreateResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSingletreasureActivityCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
