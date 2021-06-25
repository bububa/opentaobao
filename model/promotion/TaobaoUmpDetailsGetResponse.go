package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情列表 APIResponse
taobao.ump.details.get

分页查询优惠详情列表
*/
type TaobaoUmpDetailsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpDetailsGetResponse `json:"taobao_ump_details_get_response,omitempty"`
}

type TaobaoUmpDetailsGetResponse struct {

    // 活动详情的信息
    Contents   []String `json:"contents,omitempty"`

    // 记录总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
