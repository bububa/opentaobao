package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户移除优惠标签 APIResponse
tmall.promotag.taguser.remove

给用户载体去标
*/
type TmallPromotagTaguserRemoveAPIResponse struct {
    model.CommonResponse
    TmallPromotagTaguserRemoveResponse
}

type TmallPromotagTaguserRemoveResponse struct {
    XMLName xml.Name `xml:"tmall_promotag_taguser_remove_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 打标结果是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
