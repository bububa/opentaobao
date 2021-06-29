package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流删除创意 APIResponse
taobao.feedflow.item.creative.delete

信息流删除创意
*/
type TaobaoFeedflowItemCreativeDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCreativeDeleteResponse
}

type TaobaoFeedflowItemCreativeDeleteResponse struct {
    XMLName xml.Name `xml:"feedflow_item_creative_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对下
    
    Result   *TaobaoFeedflowItemCreativeDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
