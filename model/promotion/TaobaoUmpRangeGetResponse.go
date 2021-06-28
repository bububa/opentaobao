package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动范围 APIResponse
taobao.ump.range.get

查询某个卖家所有参加或者不参加某项活动的物品
*/
type TaobaoUmpRangeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ump_range_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 营销范围类列表！
    
    Ranges   []Range `json:"ranges,omitempty" xml:"