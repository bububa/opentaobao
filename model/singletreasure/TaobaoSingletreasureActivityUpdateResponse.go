package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动接口 APIResponse
taobao.singletreasure.activity.update

修改活动接口
*/
type TaobaoSingletreasureActivityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityUpdateResponse
}

type TaobaoSingletreasureActivityUpdateResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSingletreasureActivityUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
