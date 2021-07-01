package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingAititlegenerateAPIResponse
标题智能优化 API返回值
alibaba.seaking.aititlegenerate

标题智能优化 */
type AlibabaSeakingAititlegenerateAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingAititlegenerateAPIResponseModel
}

// AlibabaSeakingAititlegenerateAPIResponseModel is 标题智能优化 成功返回结果
type AlibabaSeakingAititlegenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_aititlegenerate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 智能生成的标题列表
	ResultList []string `json:"result_list,omitempty" xml:"result_list>string,omitempty"`
}
