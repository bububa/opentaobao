package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情 APIResponse
taobao.ump.detail.get

查询活动详情
*/
type TaobaoUmpDetailGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpDetailGetResponse
}

type TaobaoUmpDetailGetResponse struct {
    XMLName xml.Name `xml:"ump_detail_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动详情信息，可以通过ump sdk中的MarketingTool来进行处理
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`

    
}
