package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分组移动 APIResponse
taobao.crm.group.move

将一个分组下的所有会员移动到另一个分组，会员从原分组中删除<br/>注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
*/
type TaobaoCrmGroupMoveAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCrmGroupMoveResponse `json:"crm_group_move_response,omitempty"` 
    TaobaoCrmGroupMoveResponse
}

/* model for simplify = false
type TaobaoCrmGroupMoveResponse struct {

    // 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoCrmGroupMoveResponse struct {

    // 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测
    IsSuccess   bool `json:"is_success,omitempty"`

}
