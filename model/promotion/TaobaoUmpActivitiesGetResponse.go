package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表 APIResponse
taobao.ump.activities.get

查询活动列表
*/
type TaobaoUmpActivitiesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ump_activities_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 营销活动内容，可以通过ump sdk来进行处理
    
    Contents   []string `json:"contents,omitempty" xml:"