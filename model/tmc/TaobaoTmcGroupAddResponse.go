package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
为已开通用户添加用户分组 APIResponse
taobao.tmc.group.add

为已开通用户添加用户分组，授权消息使用
*/
type TaobaoTmcGroupAddAPIResponse struct {
    model.CommonResponse
    TaobaoTmcGroupAddResponse
}

type TaobaoTmcGroupAddResponse struct {
    XMLName xml.Name `xml:"tmc_group_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`

    
    // 分组名称
    
    GroupName   string `json:"group_name,omitempty" xml:"group_name,omitempty"`

    
}
