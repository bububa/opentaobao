package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
差旅申请单同步接口 APIResponse
alibaba.happytrip.travel.sync

以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
*/
type AlibabaHappytripTravelSyncAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTravelSyncResponse
}

type AlibabaHappytripTravelSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_travel_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 差旅申请单在欢行内部产生的差旅单ID
    
    TravelId   int64 `json:"travel_id,omitempty" xml:"travel_id,omitempty"`

    
}
