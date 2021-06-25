package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
为已开通用户添加用户分组 APIResponse
taobao.tmc.group.add

为已开通用户添加用户分组，授权消息使用
*/
type TaobaoTmcGroupAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcGroupAddResponse `json:"taobao_tmc_group_add_response,omitempty"`
}

type TaobaoTmcGroupAddResponse struct {

    // 创建时间
    Created   string `json:"created,omitempty"`

    // 分组名称
    GroupName   string `json:"group_name,omitempty"`

}
