package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家的分组 APIResponse
taobao.crm.groups.get

查询卖家的分组，返回查询到的分组列表，分页返回分组
*/
type TaobaoCrmGroupsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_groups_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询到的当前卖家的当前页的会员
    
    Groups   []Group `json:"groups,omitempty" xml:"