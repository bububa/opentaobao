package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表 APIResponse
taobao.ump.activities.get

查询活动列表
*/
type TaobaoUmpActivitiesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpActivitiesGetResponse `json:"ump_activities_get_response,omitempty"` 
    TaobaoUmpActivitiesGetResponse
}

/* model for simplify = false
type TaobaoUmpActivitiesGetResponse struct {

    // 营销活动内容，可以通过ump sdk来进行处理
    
    Contents  struct {
        String  []string `json:"string,omitempty"`
    } `json:"contents,omitempty"`
    

    // 记录总数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

type TaobaoUmpActivitiesGetResponse struct {

    // 营销活动内容，可以通过ump sdk来进行处理
    Contents   []string `json:"contents,omitempty"`

    // 记录总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
