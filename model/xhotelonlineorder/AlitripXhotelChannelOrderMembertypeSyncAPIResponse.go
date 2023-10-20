package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelOrderMembertypeSyncAPIResponse 酒店分销渠道会员类型同步 API返回值
// alitrip.xhotel.channel.order.membertype.sync
//
// 酒店分销渠道会员类型同步
type AlitripXhotelChannelOrderMembertypeSyncAPIResponse struct {
	model.CommonResponse
	AlitripXhotelChannelOrderMembertypeSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripXhotelChannelOrderMembertypeSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripXhotelChannelOrderMembertypeSyncAPIResponseModel).Reset()
}

// AlitripXhotelChannelOrderMembertypeSyncAPIResponseModel is 酒店分销渠道会员类型同步 成功返回结果
type AlitripXhotelChannelOrderMembertypeSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_xhotel_channel_order_membertype_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripXhotelChannelOrderMembertypeSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripXhotelChannelOrderMembertypeSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripXhotelChannelOrderMembertypeSyncAPIResponse)
	},
}

// GetAlitripXhotelChannelOrderMembertypeSyncAPIResponse 从 sync.Pool 获取 AlitripXhotelChannelOrderMembertypeSyncAPIResponse
func GetAlitripXhotelChannelOrderMembertypeSyncAPIResponse() *AlitripXhotelChannelOrderMembertypeSyncAPIResponse {
	return poolAlitripXhotelChannelOrderMembertypeSyncAPIResponse.Get().(*AlitripXhotelChannelOrderMembertypeSyncAPIResponse)
}

// ReleaseAlitripXhotelChannelOrderMembertypeSyncAPIResponse 将 AlitripXhotelChannelOrderMembertypeSyncAPIResponse 保存到 sync.Pool
func ReleaseAlitripXhotelChannelOrderMembertypeSyncAPIResponse(v *AlitripXhotelChannelOrderMembertypeSyncAPIResponse) {
	v.Reset()
	poolAlitripXhotelChannelOrderMembertypeSyncAPIResponse.Put(v)
}
