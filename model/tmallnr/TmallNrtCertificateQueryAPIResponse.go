package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCertificateQueryAPIResponse 批量查询电子凭证信息 API返回值
// tmall.nrt.certificate.query
//
// 批量查询电子凭证信息
type TmallNrtCertificateQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtCertificateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtCertificateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtCertificateQueryAPIResponseModel).Reset()
}

// TmallNrtCertificateQueryAPIResponseModel is 批量查询电子凭证信息 成功返回结果
type TmallNrtCertificateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_certificate_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtCertificateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtCertificateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtCertificateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtCertificateQueryAPIResponse)
	},
}

// GetTmallNrtCertificateQueryAPIResponse 从 sync.Pool 获取 TmallNrtCertificateQueryAPIResponse
func GetTmallNrtCertificateQueryAPIResponse() *TmallNrtCertificateQueryAPIResponse {
	return poolTmallNrtCertificateQueryAPIResponse.Get().(*TmallNrtCertificateQueryAPIResponse)
}

// ReleaseTmallNrtCertificateQueryAPIResponse 将 TmallNrtCertificateQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtCertificateQueryAPIResponse(v *TmallNrtCertificateQueryAPIResponse) {
	v.Reset()
	poolTmallNrtCertificateQueryAPIResponse.Put(v)
}
