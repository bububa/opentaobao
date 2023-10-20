package youkudsp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuDspDeliveryResourceMultigetAPIResponse 优酷实时批量获取可投放设备资源 API返回值
// youku.dsp.delivery.resource.multiget
//
// 优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取
type YoukuDspDeliveryResourceMultigetAPIResponse struct {
	model.CommonResponse
	YoukuDspDeliveryResourceMultigetAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuDspDeliveryResourceMultigetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuDspDeliveryResourceMultigetAPIResponseModel).Reset()
}

// YoukuDspDeliveryResourceMultigetAPIResponseModel is 优酷实时批量获取可投放设备资源 成功返回结果
type YoukuDspDeliveryResourceMultigetAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_dsp_delivery_resource_multiget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 素材列表
	Models []DeliveryList `json:"models,omitempty" xml:"models>delivery_list,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 额外信息
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

// Reset 清空结构体
func (m *YoukuDspDeliveryResourceMultigetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Models = m.Models[:0]
	m.MsgCode = ""
	m.ExtraInfo = ""
	m.MsgInfo = ""
	m.SuccessFlag = false
}

var poolYoukuDspDeliveryResourceMultigetAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuDspDeliveryResourceMultigetAPIResponse)
	},
}

// GetYoukuDspDeliveryResourceMultigetAPIResponse 从 sync.Pool 获取 YoukuDspDeliveryResourceMultigetAPIResponse
func GetYoukuDspDeliveryResourceMultigetAPIResponse() *YoukuDspDeliveryResourceMultigetAPIResponse {
	return poolYoukuDspDeliveryResourceMultigetAPIResponse.Get().(*YoukuDspDeliveryResourceMultigetAPIResponse)
}

// ReleaseYoukuDspDeliveryResourceMultigetAPIResponse 将 YoukuDspDeliveryResourceMultigetAPIResponse 保存到 sync.Pool
func ReleaseYoukuDspDeliveryResourceMultigetAPIResponse(v *YoukuDspDeliveryResourceMultigetAPIResponse) {
	v.Reset()
	poolYoukuDspDeliveryResourceMultigetAPIResponse.Put(v)
}
