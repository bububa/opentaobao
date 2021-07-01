package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品人群 API返回值 
taobao.feedflow.item.crowd.delete

删除单品人群
*/
type TaobaoFeedflowItemCrowdDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdDeleteResponse
}

// 删除单品人群 成功返回结果
type TaobaoFeedflowItemCrowdDeleteResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
