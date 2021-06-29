package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询可用广告位列表 APIResponse
taobao.feedflow.item.adzone.list

批量查询可用广告位列表
*/
type TaobaoFeedflowItemAdzoneListAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdzoneListResponse
}

type TaobaoFeedflowItemAdzoneListResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adzone_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemAdzoneListResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
