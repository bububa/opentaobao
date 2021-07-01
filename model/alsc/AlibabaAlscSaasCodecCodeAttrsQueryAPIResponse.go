package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse
码业务属性查询 API返回值
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询 */
type AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel
}

// AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel is 码业务属性查询 成功返回结果
type AlibabaAlscSaasCodecCodeAttrsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_saas_codec_code_attrs_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlscSaasCodecCodeAttrsQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
