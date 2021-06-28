package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给用户打上优惠标签 APIResponse
tmall.promotag.taguser.save

给用户载体打标
*/
type TmallPromotagTaguserSaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_promotag_taguser_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 打标结果是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"