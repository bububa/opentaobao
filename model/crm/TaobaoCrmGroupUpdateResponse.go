package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改一个已经存在的分组 APIResponse
taobao.crm.group.update

修改一个已经存在的分组，接口返回分组的修改是否成功
*/
type TaobaoCrmGroupUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGroupUpdateResponse `json:"taobao_crm_group_update_response,omitempty"`
}

type TaobaoCrmGroupUpdateResponse struct {

    // 分组修改是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
