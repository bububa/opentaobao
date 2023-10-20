package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudrealnamecheckAPIResponse 验证姓名和证件号 API返回值
// alibaba.security.jaq.rp.cloud.realname.check
//
// 验证姓名和证件号
type AlibabasecurityjaqrpcloudrealnamecheckAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqrpcloudrealnamecheckAPIResponseModel
}

// AlibabasecurityjaqrpcloudrealnamecheckAPIResponseModel is 验证姓名和证件号 成功返回结果
type AlibabasecurityjaqrpcloudrealnamecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_realname_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RealNameResult `json:"data,omitempty" xml:"data,omitempty"`
}
