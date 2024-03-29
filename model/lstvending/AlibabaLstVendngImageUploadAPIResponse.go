package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendngImageUploadAPIResponse 售货机商品图片上传 API返回值
// alibaba.lst.vendng.image.upload
//
// 零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
type AlibabaLstVendngImageUploadAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendngImageUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendngImageUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendngImageUploadAPIResponseModel).Reset()
}

// AlibabaLstVendngImageUploadAPIResponseModel is 售货机商品图片上传 成功返回结果
type AlibabaLstVendngImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vendng_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlibabaLstVendngImageUploadResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendngImageUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVendngImageUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendngImageUploadAPIResponse)
	},
}

// GetAlibabaLstVendngImageUploadAPIResponse 从 sync.Pool 获取 AlibabaLstVendngImageUploadAPIResponse
func GetAlibabaLstVendngImageUploadAPIResponse() *AlibabaLstVendngImageUploadAPIResponse {
	return poolAlibabaLstVendngImageUploadAPIResponse.Get().(*AlibabaLstVendngImageUploadAPIResponse)
}

// ReleaseAlibabaLstVendngImageUploadAPIResponse 将 AlibabaLstVendngImageUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendngImageUploadAPIResponse(v *AlibabaLstVendngImageUploadAPIResponse) {
	v.Reset()
	poolAlibabaLstVendngImageUploadAPIResponse.Put(v)
}
