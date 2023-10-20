package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityUpdateAPIResponse ISV活动修改 API返回值
// alibaba.lsy.crm.activity.update
//
// ISV活动修改
type AlibabaLsyCrmActivityUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivityUpdateAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivityUpdateAPIResponseModel is ISV活动修改 成功返回结果
type AlibabaLsyCrmActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLsyCrmActivityUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityUpdateAPIResponse)
	},
}

// GetAlibabaLsyCrmActivityUpdateAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivityUpdateAPIResponse
func GetAlibabaLsyCrmActivityUpdateAPIResponse() *AlibabaLsyCrmActivityUpdateAPIResponse {
	return poolAlibabaLsyCrmActivityUpdateAPIResponse.Get().(*AlibabaLsyCrmActivityUpdateAPIResponse)
}

// ReleaseAlibabaLsyCrmActivityUpdateAPIResponse 将 AlibabaLsyCrmActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivityUpdateAPIResponse(v *AlibabaLsyCrmActivityUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivityUpdateAPIResponse.Put(v)
}
