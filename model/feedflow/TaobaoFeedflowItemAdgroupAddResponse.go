package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加单元 APIResponse
taobao.feedflow.item.adgroup.add

信息流增加单元
*/
type TaobaoFeedflowItemAdgroupAddAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupAddResponse
}

type TaobaoFeedflowItemAdgroupAddResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemAdgroupAddResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
