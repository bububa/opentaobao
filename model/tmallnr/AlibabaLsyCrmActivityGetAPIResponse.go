package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityGetAPIResponse 私域导购查询活动详情 API返回值
// alibaba.lsy.crm.activity.get
//
// 私域导购查询活动详情
type AlibabaLsyCrmActivityGetAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivityGetAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivityGetAPIResponseModel is 私域导购查询活动详情 成功返回结果
type AlibabaLsyCrmActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaLsyCrmActivityGetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmActivityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityGetAPIResponse)
	},
}

// GetAlibabaLsyCrmActivityGetAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivityGetAPIResponse
func GetAlibabaLsyCrmActivityGetAPIResponse() *AlibabaLsyCrmActivityGetAPIResponse {
	return poolAlibabaLsyCrmActivityGetAPIResponse.Get().(*AlibabaLsyCrmActivityGetAPIResponse)
}

// ReleaseAlibabaLsyCrmActivityGetAPIResponse 将 AlibabaLsyCrmActivityGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivityGetAPIResponse(v *AlibabaLsyCrmActivityGetAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivityGetAPIResponse.Put(v)
}
