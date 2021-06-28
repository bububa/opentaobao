package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家创建一个分组 APIResponse
taobao.crm.group.add

卖家创建一个新的分组，接口返回一个创建成功的分组的id
*/
type TaobaoCrmGroupAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_group_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 添加分组是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"