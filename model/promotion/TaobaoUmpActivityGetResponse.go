package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询营销活动 APIResponse
taobao.ump.activity.get

查询营销活动
*/
type TaobaoUmpActivityGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpActivityGetResponse
}

type TaobaoUmpActivityGetResponse struct {
    XMLName xml.Name `xml:"ump_activity_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 营销活动的内容，可以通过ump sdk中的marketingTool来完成对该内容的处理
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`

    
}
