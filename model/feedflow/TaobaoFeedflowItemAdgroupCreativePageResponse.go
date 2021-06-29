package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元下查看创意 APIResponse
taobao.feedflow.item.adgroup.creative.page

信息流单元下查看创意
*/
type TaobaoFeedflowItemAdgroupCreativePageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupCreativePageResponse
}

type TaobaoFeedflowItemAdgroupCreativePageResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_creative_page_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果返回消息
    
    Result   *TaobaoFeedflowItemAdgroupCreativePageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
