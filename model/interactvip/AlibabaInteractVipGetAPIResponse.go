package interactvip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractVipGetAPIResponse 会员淘气值获取 API返回值
// alibaba.interact.vip.get
//
// 提供用户淘气值&amp;用户角色身份查询
type AlibabaInteractVipGetAPIResponse struct {
	model.CommonResponse
	AlibabaInteractVipGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractVipGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractVipGetAPIResponseModel).Reset()
}

// AlibabaInteractVipGetAPIResponseModel is 会员淘气值获取 成功返回结果
type AlibabaInteractVipGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_vip_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractVipGetAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAlibabaInteractVipGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractVipGetAPIResponse)
	},
}

// GetAlibabaInteractVipGetAPIResponse 从 sync.Pool 获取 AlibabaInteractVipGetAPIResponse
func GetAlibabaInteractVipGetAPIResponse() *AlibabaInteractVipGetAPIResponse {
	return poolAlibabaInteractVipGetAPIResponse.Get().(*AlibabaInteractVipGetAPIResponse)
}

// ReleaseAlibabaInteractVipGetAPIResponse 将 AlibabaInteractVipGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractVipGetAPIResponse(v *AlibabaInteractVipGetAPIResponse) {
	v.Reset()
	poolAlibabaInteractVipGetAPIResponse.Put(v)
}
