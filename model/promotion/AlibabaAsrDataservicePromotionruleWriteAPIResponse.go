package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAsrDataservicePromotionruleWriteAPIResponse 业务优惠规则写入 API返回值
// alibaba.asr.dataservice.promotionrule.write
//
// 星巴克优惠规则写入
type AlibabaAsrDataservicePromotionruleWriteAPIResponse struct {
	model.CommonResponse
	AlibabaAsrDataservicePromotionruleWriteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAsrDataservicePromotionruleWriteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAsrDataservicePromotionruleWriteAPIResponseModel).Reset()
}

// AlibabaAsrDataservicePromotionruleWriteAPIResponseModel is 业务优惠规则写入 成功返回结果
type AlibabaAsrDataservicePromotionruleWriteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_asr_dataservice_promotionrule_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *DataServiceResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAsrDataservicePromotionruleWriteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAsrDataservicePromotionruleWriteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAsrDataservicePromotionruleWriteAPIResponse)
	},
}

// GetAlibabaAsrDataservicePromotionruleWriteAPIResponse 从 sync.Pool 获取 AlibabaAsrDataservicePromotionruleWriteAPIResponse
func GetAlibabaAsrDataservicePromotionruleWriteAPIResponse() *AlibabaAsrDataservicePromotionruleWriteAPIResponse {
	return poolAlibabaAsrDataservicePromotionruleWriteAPIResponse.Get().(*AlibabaAsrDataservicePromotionruleWriteAPIResponse)
}

// ReleaseAlibabaAsrDataservicePromotionruleWriteAPIResponse 将 AlibabaAsrDataservicePromotionruleWriteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAsrDataservicePromotionruleWriteAPIResponse(v *AlibabaAsrDataservicePromotionruleWriteAPIResponse) {
	v.Reset()
	poolAlibabaAsrDataservicePromotionruleWriteAPIResponse.Put(v)
}
