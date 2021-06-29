package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元修改 APIResponse
taobao.feedflow.item.adgroup.modify

信息流单元修改
*/
type TaobaoFeedflowItemAdgroupModifyAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupModifyResponse
}

type TaobaoFeedflowItemAdgroupModifyResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemAdgroupModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
