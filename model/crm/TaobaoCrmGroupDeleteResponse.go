package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除分组 APIResponse
taobao.crm.group.delete

将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
type TaobaoCrmGroupDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGroupDeleteResponse `json:"taobao_crm_group_delete_response,omitempty"`
}

type TaobaoCrmGroupDeleteResponse struct {

    // 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测
    IsSuccess   bool `json:"is_success,omitempty"`

}
