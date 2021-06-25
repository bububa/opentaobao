package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家的分组 APIResponse
taobao.crm.groups.get

查询卖家的分组，返回查询到的分组列表，分页返回分组
*/
type TaobaoCrmGroupsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCrmGroupsGetResponse `json:"taobao_crm_groups_get_response,omitempty"`
}

type TaobaoCrmGroupsGetResponse struct {

    // 查询到的当前卖家的当前页的会员
    Groups   []Group `json:"groups,omitempty"`

    // 记录总数
    TotalResult   int64 `json:"total_result,omitempty"`

}
