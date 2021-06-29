package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流新增并且绑定创意 APIResponse
taobao.feedflow.item.adgroup.creative.add.bind

信息流新增并且绑定创意
*/
type TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupCreativeAddBindResponse
}

type TaobaoFeedflowItemAdgroupCreativeAddBindResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_creative_add_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *TaobaoFeedflowItemAdgroupCreativeAddBindResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
