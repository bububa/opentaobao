package ascpqcc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpqccsampleupdateAPIResponse 品控中心更新样品信息 API返回值
// alibaba.ascp.qcc.sample.update
//
// 品控中心更新样品信息
type AlibabaascpqccsampleupdateAPIResponse struct {
	model.CommonResponse
	AlibabaascpqccsampleupdateAPIResponseModel
}

// AlibabaascpqccsampleupdateAPIResponseModel is 品控中心更新样品信息 成功返回结果
type AlibabaascpqccsampleupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_qcc_sample_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SendResult `json:"result,omitempty" xml:"result,omitempty"`
}
