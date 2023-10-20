package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelNotifyAPIResponse 分销渠道各类通知接口 API返回值
// alitrip.xhotel.channel.notify
//
// 分销渠道支付通知
type AlitripXhotelChannelNotifyAPIResponse struct {
	model.CommonResponse
	AlitripXhotelChannelNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripXhotelChannelNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripXhotelChannelNotifyAPIResponseModel).Reset()
}

// AlitripXhotelChannelNotifyAPIResponseModel is 分销渠道各类通知接口 成功返回结果
type AlitripXhotelChannelNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_xhotel_channel_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripXhotelChannelNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripXhotelChannelNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripXhotelChannelNotifyAPIResponse)
	},
}

// GetAlitripXhotelChannelNotifyAPIResponse 从 sync.Pool 获取 AlitripXhotelChannelNotifyAPIResponse
func GetAlitripXhotelChannelNotifyAPIResponse() *AlitripXhotelChannelNotifyAPIResponse {
	return poolAlitripXhotelChannelNotifyAPIResponse.Get().(*AlitripXhotelChannelNotifyAPIResponse)
}

// ReleaseAlitripXhotelChannelNotifyAPIResponse 将 AlitripXhotelChannelNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripXhotelChannelNotifyAPIResponse(v *AlitripXhotelChannelNotifyAPIResponse) {
	v.Reset()
	poolAlitripXhotelChannelNotifyAPIResponse.Put(v)
}
