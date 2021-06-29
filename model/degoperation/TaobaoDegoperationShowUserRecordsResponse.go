package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户中奖记录 APIResponse
taobao.degoperation.show.user.records

用户中奖记录
*/
type TaobaoDegoperationShowUserRecordsAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationShowUserRecordsResponse
}

type TaobaoDegoperationShowUserRecordsResponse struct {
    XMLName xml.Name `xml:"degoperation_show_user_records_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
