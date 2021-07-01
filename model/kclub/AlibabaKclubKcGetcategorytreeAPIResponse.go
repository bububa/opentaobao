package kclub

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcGetcategorytreeAPIResponse
知识云-查询租户下类目树 API返回值
alibaba.kclub.kc.getcategorytree

知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。 */
type AlibabaKclubKcGetcategorytreeAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcGetcategorytreeAPIResponseModel
}

// AlibabaKclubKcGetcategorytreeAPIResponseModel is 知识云-查询租户下类目树 成功返回结果
type AlibabaKclubKcGetcategorytreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_getcategorytree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaKclubKcGetcategorytreeResult `json:"result,omitempty" xml:"result,omitempty"`
}
