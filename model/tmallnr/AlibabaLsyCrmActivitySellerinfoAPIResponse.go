package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivitySellerinfoAPIResponse 商家信息推送 API返回值
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
type AlibabaLsyCrmActivitySellerinfoAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivitySellerinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivitySellerinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivitySellerinfoAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivitySellerinfoAPIResponseModel is 商家信息推送 成功返回结果
type AlibabaLsyCrmActivitySellerinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_sellerinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivitySellerinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmActivitySellerinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivitySellerinfoAPIResponse)
	},
}

// GetAlibabaLsyCrmActivitySellerinfoAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivitySellerinfoAPIResponse
func GetAlibabaLsyCrmActivitySellerinfoAPIResponse() *AlibabaLsyCrmActivitySellerinfoAPIResponse {
	return poolAlibabaLsyCrmActivitySellerinfoAPIResponse.Get().(*AlibabaLsyCrmActivitySellerinfoAPIResponse)
}

// ReleaseAlibabaLsyCrmActivitySellerinfoAPIResponse 将 AlibabaLsyCrmActivitySellerinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivitySellerinfoAPIResponse(v *AlibabaLsyCrmActivitySellerinfoAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivitySellerinfoAPIResponse.Put(v)
}
