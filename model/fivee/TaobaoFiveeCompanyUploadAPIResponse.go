package fivee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeCompanyUploadAPIResponse 上传商信息接口 API返回值
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
type TaobaoFiveeCompanyUploadAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeCompanyUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFiveeCompanyUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFiveeCompanyUploadAPIResponseModel).Reset()
}

// TaobaoFiveeCompanyUploadAPIResponseModel is 上传商信息接口 成功返回结果
type TaobaoFiveeCompanyUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_company_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	CodeT int64 `json:"code_t,omitempty" xml:"code_t,omitempty"`
	// 是否成功
	SuccessT bool `json:"success_t,omitempty" xml:"success_t,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFiveeCompanyUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.CodeT = 0
	m.SuccessT = false
}

var poolTaobaoFiveeCompanyUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeCompanyUploadAPIResponse)
	},
}

// GetTaobaoFiveeCompanyUploadAPIResponse 从 sync.Pool 获取 TaobaoFiveeCompanyUploadAPIResponse
func GetTaobaoFiveeCompanyUploadAPIResponse() *TaobaoFiveeCompanyUploadAPIResponse {
	return poolTaobaoFiveeCompanyUploadAPIResponse.Get().(*TaobaoFiveeCompanyUploadAPIResponse)
}

// ReleaseTaobaoFiveeCompanyUploadAPIResponse 将 TaobaoFiveeCompanyUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFiveeCompanyUploadAPIResponse(v *TaobaoFiveeCompanyUploadAPIResponse) {
	v.Reset()
	poolTaobaoFiveeCompanyUploadAPIResponse.Put(v)
}
