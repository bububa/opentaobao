package kclub

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQueryknowledgeAPIResponse
知识云-通用知识查询服务 API返回值
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。 */
type AlibabaKclubKcQueryknowledgeAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcQueryknowledgeAPIResponseModel
}

// AlibabaKclubKcQueryknowledgeAPIResponseModel is 知识云-通用知识查询服务 成功返回结果
type AlibabaKclubKcQueryknowledgeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_queryknowledge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaKclubKcQueryknowledgeResult `json:"result,omitempty" xml:"result,omitempty"`
}
