package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品人群 APIResponse
taobao.feedflow.item.crowd.delete

删除单品人群
*/
type TaobaoFeedflowItemCrowdDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdDeleteResponse
}

type TaobaoFeedflowItemCrowdDeleteResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemCrowdDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
