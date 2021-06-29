package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品优惠接口 API返回值 
taobao.singletreasure.activity.item.delete

删除单品优惠接口
*/
type TaobaoSingletreasureActivityItemDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemDeleteResponse
}

// 删除单品优惠接口 成功返回结果
type TaobaoSingletreasureActivityItemDeleteResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoSingletreasureActivityItemDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
