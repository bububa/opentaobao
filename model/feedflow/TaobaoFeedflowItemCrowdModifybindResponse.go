package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改人群出价或状态 APIResponse
taobao.feedflow.item.crowd.modifybind

修改人群出价或状态
*/
type TaobaoFeedflowItemCrowdModifybindAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdModifybindResponse
}

type TaobaoFeedflowItemCrowdModifybindResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_modifybind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemCrowdModifybindResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
