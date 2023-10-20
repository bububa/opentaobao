package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaasrdataservicepromotionruledeleteAPIResponse 优惠规则删除 API返回值
// alibaba.asr.dataservice.promotionrule.delete
//
// 删除优惠规则，例如星巴克删除优惠规则
type AlibabaasrdataservicepromotionruledeleteAPIResponse struct {
	model.CommonResponse
	AlibabaasrdataservicepromotionruledeleteAPIResponseModel
}

// AlibabaasrdataservicepromotionruledeleteAPIResponseModel is 优惠规则删除 成功返回结果
type AlibabaasrdataservicepromotionruledeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_asr_dataservice_promotionrule_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *DataServiceResponse `json:"result,omitempty" xml:"result,omitempty"`
}
