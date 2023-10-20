package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityDataUpdateAPIResponse 私域导购数据回流接口 API返回值
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
type AlibabaLsyCrmActivityDataUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityDataUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityDataUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivityDataUpdateAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivityDataUpdateAPIResponseModel is 私域导购数据回流接口 成功返回结果
type AlibabaLsyCrmActivityDataUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_data_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLsyCrmActivityDataUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityDataUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmActivityDataUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityDataUpdateAPIResponse)
	},
}

// GetAlibabaLsyCrmActivityDataUpdateAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivityDataUpdateAPIResponse
func GetAlibabaLsyCrmActivityDataUpdateAPIResponse() *AlibabaLsyCrmActivityDataUpdateAPIResponse {
	return poolAlibabaLsyCrmActivityDataUpdateAPIResponse.Get().(*AlibabaLsyCrmActivityDataUpdateAPIResponse)
}

// ReleaseAlibabaLsyCrmActivityDataUpdateAPIResponse 将 AlibabaLsyCrmActivityDataUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivityDataUpdateAPIResponse(v *AlibabaLsyCrmActivityDataUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivityDataUpdateAPIResponse.Put(v)
}
