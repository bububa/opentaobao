package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景创建或更新组团活动 APIResponse
taobao.opentrade.group.sync

组团购场景中创建团购活动
*/
type TaobaoOpentradeGroupSyncAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupSyncResponse
}

type TaobaoOpentradeGroupSyncResponse struct {
    XMLName xml.Name `xml:"opentrade_group_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 团购活动信息
    
    Result   *GroupActivityResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
