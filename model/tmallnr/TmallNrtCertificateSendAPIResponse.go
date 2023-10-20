package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCertificateSendAPIResponse 有价礼包发放电子凭证 API返回值
// tmall.nrt.certificate.send
//
// 支持有价礼包发放电子凭证
type TmallNrtCertificateSendAPIResponse struct {
	model.CommonResponse
	TmallNrtCertificateSendAPIResponseModel
}

// TmallNrtCertificateSendAPIResponseModel is 有价礼包发放电子凭证 成功返回结果
type TmallNrtCertificateSendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_certificate_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtCertificateSendResult `json:"result,omitempty" xml:"result,omitempty"`
}
