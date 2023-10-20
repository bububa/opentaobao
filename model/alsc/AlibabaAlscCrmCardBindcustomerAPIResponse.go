package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBindcustomerAPIResponse 卡号绑定顾客 API返回值
// alibaba.alsc.crm.card.bindcustomer
//
// 为卡号绑定顾客
type AlibabaAlscCrmCardBindcustomerAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardBindcustomerAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardBindcustomerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardBindcustomerAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardBindcustomerAPIResponseModel is 卡号绑定顾客 成功返回结果
type AlibabaAlscCrmCardBindcustomerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_bindcustomer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardBindcustomerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardBindcustomerAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardBindcustomerAPIResponse)
	},
}

// GetAlibabaAlscCrmCardBindcustomerAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardBindcustomerAPIResponse
func GetAlibabaAlscCrmCardBindcustomerAPIResponse() *AlibabaAlscCrmCardBindcustomerAPIResponse {
	return poolAlibabaAlscCrmCardBindcustomerAPIResponse.Get().(*AlibabaAlscCrmCardBindcustomerAPIResponse)
}

// ReleaseAlibabaAlscCrmCardBindcustomerAPIResponse 将 AlibabaAlscCrmCardBindcustomerAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardBindcustomerAPIResponse(v *AlibabaAlscCrmCardBindcustomerAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardBindcustomerAPIResponse.Put(v)
}
