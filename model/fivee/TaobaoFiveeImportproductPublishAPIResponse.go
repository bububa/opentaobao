package fivee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeImportproductPublishAPIResponse 进口商品发布 API返回值
// taobao.fivee.importproduct.publish
//
// 直营业务商家入住发布商品时，上传商品及商家证照信息
type TaobaoFiveeImportproductPublishAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeImportproductPublishAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFiveeImportproductPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFiveeImportproductPublishAPIResponseModel).Reset()
}

// TaobaoFiveeImportproductPublishAPIResponseModel is 进口商品发布 成功返回结果
type TaobaoFiveeImportproductPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_importproduct_publish_response"`
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
func (m *TaobaoFiveeImportproductPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.CodeT = 0
	m.SuccessT = false
}

var poolTaobaoFiveeImportproductPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeImportproductPublishAPIResponse)
	},
}

// GetTaobaoFiveeImportproductPublishAPIResponse 从 sync.Pool 获取 TaobaoFiveeImportproductPublishAPIResponse
func GetTaobaoFiveeImportproductPublishAPIResponse() *TaobaoFiveeImportproductPublishAPIResponse {
	return poolTaobaoFiveeImportproductPublishAPIResponse.Get().(*TaobaoFiveeImportproductPublishAPIResponse)
}

// ReleaseTaobaoFiveeImportproductPublishAPIResponse 将 TaobaoFiveeImportproductPublishAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFiveeImportproductPublishAPIResponse(v *TaobaoFiveeImportproductPublishAPIResponse) {
	v.Reset()
	poolTaobaoFiveeImportproductPublishAPIResponse.Put(v)
}
