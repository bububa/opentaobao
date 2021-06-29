package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单品单元下，新增定向人群 APIResponse
taobao.feedflow.item.crowd.add

单品单元下，新增定向人群
*/
type TaobaoFeedflowItemCrowdAddAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdAddResponse
}

type TaobaoFeedflowItemCrowdAddResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemCrowdAddResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
