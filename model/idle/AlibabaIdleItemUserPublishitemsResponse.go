package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发布的商品列表 APIResponse
alibaba.idle.item.user.publishitems

为服务商的卖家提供发布的闲鱼商品列表
*/
type AlibabaIdleItemUserPublishitemsAPIResponse struct {
    model.CommonResponse
    AlibabaIdleItemUserPublishitemsResponse
}

type AlibabaIdleItemUserPublishitemsResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_item_user_publishitems_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果
    
    Result   *TopPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
