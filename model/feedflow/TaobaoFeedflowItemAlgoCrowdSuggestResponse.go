package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单品人群建议出价 APIResponse
taobao.feedflow.item.algo.crowd.suggest

给超级推荐的广告主查看建议出价
*/
type TaobaoFeedflowItemAlgoCrowdSuggestAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAlgoCrowdSuggestResponse
}

type TaobaoFeedflowItemAlgoCrowdSuggestResponse struct {
    XMLName xml.Name `xml:"feedflow_item_algo_crowd_suggest_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemAlgoCrowdSuggestResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
