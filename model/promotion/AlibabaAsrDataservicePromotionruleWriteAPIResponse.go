package promotion

import (
	"encoding/xml"

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

// AlibabaAsrDataservicePromotionruleWriteAPIResponseModel is 业务优惠规则写入 成功返回结果
type AlibabaAsrDataservicePromotionruleWriteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_asr_dataservice_promotionrule_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *DataServiceResponse `json:"result,omitempty" xml:"result,omitempty"`
}
