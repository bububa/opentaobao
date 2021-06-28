package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一个已经存在的分组 APIResponse
taobao.crm.group.update

修改一个已经存在的分组，接口返回分组的修改是否成功
*/
type TaobaoCrmGroupUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_group_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分组修改是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"