package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBindcardAPIResponse 绑定物理卡 API返回值
// alibaba.alsc.crm.card.bindcard
//
// 将会员卡和实例物理卡绑定在一起
type AlibabaAlscCrmCardBindcardAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardBindcardAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardBindcardAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardBindcardAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardBindcardAPIResponseModel is 绑定物理卡 成功返回结果
type AlibabaAlscCrmCardBindcardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_bindcard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardBindcardAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardBindcardAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardBindcardAPIResponse)
	},
}

// GetAlibabaAlscCrmCardBindcardAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardBindcardAPIResponse
func GetAlibabaAlscCrmCardBindcardAPIResponse() *AlibabaAlscCrmCardBindcardAPIResponse {
	return poolAlibabaAlscCrmCardBindcardAPIResponse.Get().(*AlibabaAlscCrmCardBindcardAPIResponse)
}

// ReleaseAlibabaAlscCrmCardBindcardAPIResponse 将 AlibabaAlscCrmCardBindcardAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardBindcardAPIResponse(v *AlibabaAlscCrmCardBindcardAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardBindcardAPIResponse.Put(v)
}
