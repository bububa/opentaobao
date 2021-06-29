package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单品单元下，新增定向人群 API返回值 
taobao.feedflow.item.crowd.add

单品单元下，新增定向人群
*/
type TaobaoFeedflowItemCrowdAddAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdAddResponse
}

// 单品单元下，新增定向人群 成功返回结果
type TaobaoFeedflowItemCrowdAddResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemCrowdAddResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
