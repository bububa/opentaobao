package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元内绑定资源位 APIResponse
taobao.feedflow.item.adgroup.adzone.bind

信息流单元内绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupAdzoneBindResponse
}

type TaobaoFeedflowItemAdgroupAdzoneBindResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemAdgroupAdzoneBindResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
