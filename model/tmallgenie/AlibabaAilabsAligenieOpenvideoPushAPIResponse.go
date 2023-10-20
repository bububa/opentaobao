package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpenvideoPushAPIResponse 天猫精灵内容库视频分集数据推送接口 API返回值
// alibaba.ailabs.aligenie.openvideo.push
//
// 天猫精灵内容库视频分集数据推送接口
type AlibabaAilabsAligenieOpenvideoPushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpenvideoPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpenvideoPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieOpenvideoPushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieOpenvideoPushAPIResponseModel is 天猫精灵内容库视频分集数据推送接口 成功返回结果
type AlibabaAilabsAligenieOpenvideoPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideo_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 状态码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpenvideoPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.RetCode = 0
}

var poolAlibabaAilabsAligenieOpenvideoPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieOpenvideoPushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieOpenvideoPushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieOpenvideoPushAPIResponse
func GetAlibabaAilabsAligenieOpenvideoPushAPIResponse() *AlibabaAilabsAligenieOpenvideoPushAPIResponse {
	return poolAlibabaAilabsAligenieOpenvideoPushAPIResponse.Get().(*AlibabaAilabsAligenieOpenvideoPushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieOpenvideoPushAPIResponse 将 AlibabaAilabsAligenieOpenvideoPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieOpenvideoPushAPIResponse(v *AlibabaAilabsAligenieOpenvideoPushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieOpenvideoPushAPIResponse.Put(v)
}
