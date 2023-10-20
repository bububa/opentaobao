package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterSpserviceorderUpdateAPIResponse 服务供应链服务单更新 API返回值
// alibaba.servicecenter.spserviceorder.update
//
// 服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
type AlibabaServicecenterSpserviceorderUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterSpserviceorderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterSpserviceorderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterSpserviceorderUpdateAPIResponseModel).Reset()
}

// AlibabaServicecenterSpserviceorderUpdateAPIResponseModel is 服务供应链服务单更新 成功返回结果
type AlibabaServicecenterSpserviceorderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_spserviceorder_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterSpserviceorderUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterSpserviceorderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterSpserviceorderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterSpserviceorderUpdateAPIResponse)
	},
}

// GetAlibabaServicecenterSpserviceorderUpdateAPIResponse 从 sync.Pool 获取 AlibabaServicecenterSpserviceorderUpdateAPIResponse
func GetAlibabaServicecenterSpserviceorderUpdateAPIResponse() *AlibabaServicecenterSpserviceorderUpdateAPIResponse {
	return poolAlibabaServicecenterSpserviceorderUpdateAPIResponse.Get().(*AlibabaServicecenterSpserviceorderUpdateAPIResponse)
}

// ReleaseAlibabaServicecenterSpserviceorderUpdateAPIResponse 将 AlibabaServicecenterSpserviceorderUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterSpserviceorderUpdateAPIResponse(v *AlibabaServicecenterSpserviceorderUpdateAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterSpserviceorderUpdateAPIResponse.Put(v)
}
