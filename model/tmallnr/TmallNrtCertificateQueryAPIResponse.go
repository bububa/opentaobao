package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcertificatequeryAPIResponse 批量查询电子凭证信息 API返回值
// tmall.nrt.certificate.query
//
// 批量查询电子凭证信息
type TmallnrtcertificatequeryAPIResponse struct {
	model.CommonResponse
	TmallnrtcertificatequeryAPIResponseModel
}

// TmallnrtcertificatequeryAPIResponseModel is 批量查询电子凭证信息 成功返回结果
type TmallnrtcertificatequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_certificate_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallnrtcertificatequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
