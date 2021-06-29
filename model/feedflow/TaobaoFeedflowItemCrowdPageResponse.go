package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询单品单元下人群列表 APIResponse
taobao.feedflow.item.crowd.page

分页查询单品单元下人群列表
*/
type TaobaoFeedflowItemCrowdPageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdPageResponse
}

type TaobaoFeedflowItemCrowdPageResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemCrowdPageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
