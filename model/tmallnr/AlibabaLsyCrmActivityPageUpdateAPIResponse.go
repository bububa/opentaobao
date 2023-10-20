package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityPageUpdateAPIResponse ISV活动页面创建修改 API返回值
// alibaba.lsy.crm.activity.page.update
//
// ISV活动页面创建修改
type AlibabaLsyCrmActivityPageUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityPageUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityPageUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivityPageUpdateAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivityPageUpdateAPIResponseModel is ISV活动页面创建修改 成功返回结果
type AlibabaLsyCrmActivityPageUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_page_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLsyCrmActivityPageUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityPageUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmActivityPageUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityPageUpdateAPIResponse)
	},
}

// GetAlibabaLsyCrmActivityPageUpdateAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivityPageUpdateAPIResponse
func GetAlibabaLsyCrmActivityPageUpdateAPIResponse() *AlibabaLsyCrmActivityPageUpdateAPIResponse {
	return poolAlibabaLsyCrmActivityPageUpdateAPIResponse.Get().(*AlibabaLsyCrmActivityPageUpdateAPIResponse)
}

// ReleaseAlibabaLsyCrmActivityPageUpdateAPIResponse 将 AlibabaLsyCrmActivityPageUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivityPageUpdateAPIResponse(v *AlibabaLsyCrmActivityPageUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivityPageUpdateAPIResponse.Put(v)
}
