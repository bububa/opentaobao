package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
覆盖单元下同类型定向人群 API返回值 
taobao.feedflow.item.crowd.modify

覆盖单元下同类型定向人群
*/
type TaobaoFeedflowItemCrowdModifyAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdModifyResponse
}

// 覆盖单元下同类型定向人群 成功返回结果
type TaobaoFeedflowItemCrowdModifyResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemCrowdModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
