package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
营销活动列表查询 APIResponse
taobao.ump.activities.list.get

按照营销活动id的列表ids，查询对应的营销活动列表。
*/
type TaobaoUmpActivitiesListGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpActivitiesListGetResponse
}

type TaobaoUmpActivitiesListGetResponse struct {
    XMLName xml.Name `xml:"ump_activities_list_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 营销活动列表！
    
    Activities   []string `json:"activities,omitempty" xml:"activities>string,omitempty"`
    
    
}
