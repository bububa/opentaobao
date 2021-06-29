package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店分销渠道会员类型同步 APIResponse
alitrip.xhotel.channel.order.membertype.sync

酒店分销渠道会员类型同步
*/
type AlitripXhotelChannelOrderMembertypeSyncAPIResponse struct {
    model.CommonResponse
    AlitripXhotelChannelOrderMembertypeSyncResponse
}

type AlitripXhotelChannelOrderMembertypeSyncResponse struct {
    XMLName xml.Name `xml:"alitrip_xhotel_channel_order_membertype_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *HbsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
