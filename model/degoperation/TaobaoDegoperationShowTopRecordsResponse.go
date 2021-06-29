package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动中奖记录 APIResponse
taobao.degoperation.show.top.records

活动中奖记录
*/
type TaobaoDegoperationShowTopRecordsAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationShowTopRecordsResponse
}

type TaobaoDegoperationShowTopRecordsResponse struct {
    XMLName xml.Name `xml:"degoperation_show_top_records_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
