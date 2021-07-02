package icburfq

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqRecommendAPIResponse rfq推荐 API返回值
// alibaba.icbu.rfq.recommend
//
// rfq推荐
type AlibabaIcbuRfqRecommendAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuRfqRecommendAPIResponseModel
}

// AlibabaIcbuRfqRecommendAPIResponseModel is rfq推荐 成功返回结果
type AlibabaIcbuRfqRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfq_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
