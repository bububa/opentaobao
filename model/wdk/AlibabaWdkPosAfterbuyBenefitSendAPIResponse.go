package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosAfterbuyBenefitSendAPIResponse 生态pos购后发放权益 API返回值
// alibaba.wdk.pos.afterbuy.benefit.send
//
// 生态pos购后发放权益接口开放
type AlibabaWdkPosAfterbuyBenefitSendAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPosAfterbuyBenefitSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPosAfterbuyBenefitSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPosAfterbuyBenefitSendAPIResponseModel).Reset()
}

// AlibabaWdkPosAfterbuyBenefitSendAPIResponseModel is 生态pos购后发放权益 成功返回结果
type AlibabaWdkPosAfterbuyBenefitSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_pos_afterbuy_benefit_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPosAfterbuyBenefitSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkPosAfterbuyBenefitSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPosAfterbuyBenefitSendAPIResponse)
	},
}

// GetAlibabaWdkPosAfterbuyBenefitSendAPIResponse 从 sync.Pool 获取 AlibabaWdkPosAfterbuyBenefitSendAPIResponse
func GetAlibabaWdkPosAfterbuyBenefitSendAPIResponse() *AlibabaWdkPosAfterbuyBenefitSendAPIResponse {
	return poolAlibabaWdkPosAfterbuyBenefitSendAPIResponse.Get().(*AlibabaWdkPosAfterbuyBenefitSendAPIResponse)
}

// ReleaseAlibabaWdkPosAfterbuyBenefitSendAPIResponse 将 AlibabaWdkPosAfterbuyBenefitSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPosAfterbuyBenefitSendAPIResponse(v *AlibabaWdkPosAfterbuyBenefitSendAPIResponse) {
	v.Reset()
	poolAlibabaWdkPosAfterbuyBenefitSendAPIResponse.Put(v)
}
