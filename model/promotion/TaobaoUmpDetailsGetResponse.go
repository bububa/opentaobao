package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情列表 APIResponse
taobao.ump.details.get

分页查询优惠详情列表
*/
type TaobaoUmpDetailsGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpDetailsGetResponse
}

type TaobaoUmpDetailsGetResponse struct {
    XMLName xml.Name `xml:"ump_details_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动详情的信息
    
    Contents   []string `json:"contents,omitempty" xml:"contents>string,omitempty"`
    
    
    // 记录总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
