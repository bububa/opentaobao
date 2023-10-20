package normalvisa

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelVisaSignSendAPIResponse 签证批量申请人送签接口 API返回值
// alitrip.travel.visa.sign.send
//
// 签证批量申请人送签接口，用于商家批量送签。
type AlitripTravelVisaSignSendAPIResponse struct {
	model.CommonResponse
	AlitripTravelVisaSignSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelVisaSignSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelVisaSignSendAPIResponseModel).Reset()
}

// AlitripTravelVisaSignSendAPIResponseModel is 签证批量申请人送签接口 成功返回结果
type AlitripTravelVisaSignSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_visa_sign_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批次信息
	BatchInfos []BatchInfo `json:"batch_infos,omitempty" xml:"batch_infos>batch_info,omitempty"`
	// 失败信息
	FailInfos []SendSignFailInfo `json:"fail_infos,omitempty" xml:"fail_infos>send_sign_fail_info,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelVisaSignSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchInfos = m.BatchInfos[:0]
	m.FailInfos = m.FailInfos[:0]
}

var poolAlitripTravelVisaSignSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelVisaSignSendAPIResponse)
	},
}

// GetAlitripTravelVisaSignSendAPIResponse 从 sync.Pool 获取 AlitripTravelVisaSignSendAPIResponse
func GetAlitripTravelVisaSignSendAPIResponse() *AlitripTravelVisaSignSendAPIResponse {
	return poolAlitripTravelVisaSignSendAPIResponse.Get().(*AlitripTravelVisaSignSendAPIResponse)
}

// ReleaseAlitripTravelVisaSignSendAPIResponse 将 AlitripTravelVisaSignSendAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelVisaSignSendAPIResponse(v *AlitripTravelVisaSignSendAPIResponse) {
	v.Reset()
	poolAlitripTravelVisaSignSendAPIResponse.Put(v)
}
