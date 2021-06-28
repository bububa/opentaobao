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
    TaobaoTmcGroupsGetResponse
}

type TaobaoTmcGroupsGetResponse struct {
    XMLName xml.Name `xml:"tmc_groups_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // dasdasd
    
    Groups   []TmcGroup `json:"groups,omitempty" xml:"groups>tmc_group,omitempty"`
    
    
    // 分组总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
