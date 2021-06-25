package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
营销活动列表查询 APIResponse
taobao.ump.activities.list.get

按照营销活动id的列表ids，查询对应的营销活动列表。
*/
type TaobaoUmpActivitiesListGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpActivitiesListGetResponse `json:"taobao_ump_activities_list_get_response,omitempty"`
}

type TaobaoUmpActivitiesListGetResponse struct {

    // 营销活动列表！
    Activities   []String `json:"activities,omitempty"`

}
