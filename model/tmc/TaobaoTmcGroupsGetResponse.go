package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取自定义用户分组列表 APIResponse
taobao.tmc.groups.get

获取自定义用户分组列表
*/
type TaobaoTmcGroupsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_groups_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // dasdasd
    
    Groups   []TmcGroup `json:"groups,omitempty" xml:"