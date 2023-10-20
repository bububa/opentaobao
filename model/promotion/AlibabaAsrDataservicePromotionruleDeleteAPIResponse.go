package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAsrDataservicePromotionruleDeleteAPIResponse 优惠规则删除 API返回值
// alibaba.asr.dataservice.promotionrule.delete
//
// 删除优惠规则，例如星巴克删除优惠规则
type AlibabaAsrDataservicePromotionruleDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAsrDataservicePromotionruleDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAsrDataservicePromotionruleDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAsrDataservicePromotionruleDeleteAPIResponseModel).Reset()
}

// AlibabaAsrDataservicePromotionruleDeleteAPIResponseModel is 优惠规则删除 成功返回结果
type AlibabaAsrDataservicePromotionruleDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_asr_dataservice_promotionrule_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *DataServiceResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAsrDataservicePromotionruleDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAsrDataservicePromotionruleDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAsrDataservicePromotionruleDeleteAPIResponse)
	},
}

// GetAlibabaAsrDataservicePromotionruleDeleteAPIResponse 从 sync.Pool 获取 AlibabaAsrDataservicePromotionruleDeleteAPIResponse
func GetAlibabaAsrDataservicePromotionruleDeleteAPIResponse() *AlibabaAsrDataservicePromotionruleDeleteAPIResponse {
	return poolAlibabaAsrDataservicePromotionruleDeleteAPIResponse.Get().(*AlibabaAsrDataservicePromotionruleDeleteAPIResponse)
}

// ReleaseAlibabaAsrDataservicePromotionruleDeleteAPIResponse 将 AlibabaAsrDataservicePromotionruleDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAsrDataservicePromotionruleDeleteAPIResponse(v *AlibabaAsrDataservicePromotionruleDeleteAPIResponse) {
	v.Reset()
	poolAlibabaAsrDataservicePromotionruleDeleteAPIResponse.Put(v)
}
