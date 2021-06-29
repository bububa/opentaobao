package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单元id删除单元 APIResponse
taobao.feedflow.item.adgroup.delete

根据单元id删除单元
*/
type TaobaoFeedflowItemAdgroupDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupDeleteResponse
}

type TaobaoFeedflowItemAdgroupDeleteResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoFeedflowItemAdgroupDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
