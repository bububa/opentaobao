package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaasrdataservicepromotionrulequeryAPIResponse 星巴克优惠规则查询 API返回值
// alibaba.asr.dataservice.promotionrule.query
//
// 查询优惠规则，例如星巴克查询优惠规则
type AlibabaasrdataservicepromotionrulequeryAPIResponse struct {
	model.CommonResponse
	AlibabaasrdataservicepromotionrulequeryAPIResponseModel
}

// AlibabaasrdataservicepromotionrulequeryAPIResponseModel is 星巴克优惠规则查询 成功返回结果
type AlibabaasrdataservicepromotionrulequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_asr_dataservice_promotionrule_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *DataServiceResponse `json:"result,omitempty" xml:"result,omitempty"`
}
