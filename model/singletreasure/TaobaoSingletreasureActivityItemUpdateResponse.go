package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品优惠接口 API返回值 
taobao.singletreasure.activity.item.update

更新单品优惠接口
*/
type TaobaoSingletreasureActivityItemUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemUpdateResponse
}

// 更新单品优惠接口 成功返回结果
type TaobaoSingletreasureActivityItemUpdateResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoSingletreasureActivityItemUpdateResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
