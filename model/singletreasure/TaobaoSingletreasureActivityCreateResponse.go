package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动创建接口 API返回值 
taobao.singletreasure.activity.create

创建优惠活动
*/
type TaobaoSingletreasureActivityCreateAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityCreateResponse
}

// 活动创建接口 成功返回结果
type TaobaoSingletreasureActivityCreateResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoSingletreasureActivityCreateResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
