package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店分销渠道会员类型同步 API返回值 
alitrip.xhotel.channel.order.membertype.sync

酒店分销渠道会员类型同步
*/
type AlitripXhotelChannelOrderMembertypeSyncAPIResponse struct {
    model.CommonResponse
    AlitripXhotelChannelOrderMembertypeSyncResponse
}

// 酒店分销渠道会员类型同步 成功返回结果
type AlitripXhotelChannelOrderMembertypeSyncResponse struct {
    XMLName xml.Name `xml:"alitrip_xhotel_channel_order_membertype_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}
