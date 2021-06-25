package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家创建一个分组 APIResponse
taobao.crm.group.add

卖家创建一个新的分组，接口返回一个创建成功的分组的id
*/
type TaobaoCrmGroupAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGroupAddResponse `json:"taobao_crm_group_add_response,omitempty"`
}

type TaobaoCrmGroupAddResponse struct {

    // 添加分组是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 新增分组的id
    GroupId   int64 `json:"group_id,omitempty"`

}
