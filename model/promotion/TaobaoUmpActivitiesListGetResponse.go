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
    // Response *TaobaoUmpActivitiesListGetResponse `json:"ump_activities_list_get_response,omitempty"` 
    TaobaoUmpActivitiesListGetResponse
}

/* model for simplify = false
type TaobaoUmpActivitiesListGetResponse struct {

    // 营销活动列表！
    
    Activities  struct {
        String  []string `json:"string,omitempty"`
    } `json:"activities,omitempty"`
    

}
*/

type TaobaoUmpActivitiesListGetResponse struct {

    // 营销活动列表！
    Activities   []string `json:"activities,omitempty"`

}
