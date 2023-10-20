package fivee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeInnerproductPublishAPIResponse 国产商品发布 API返回值
// taobao.fivee.innerproduct.publish
//
// 资质共享平台国产商品发布
type TaobaoFiveeInnerproductPublishAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeInnerproductPublishAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFiveeInnerproductPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFiveeInnerproductPublishAPIResponseModel).Reset()
}

// TaobaoFiveeInnerproductPublishAPIResponseModel is 国产商品发布 成功返回结果
type TaobaoFiveeInnerproductPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_innerproduct_publish_response"`
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
func (m *TaobaoFiveeInnerproductPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Message = ""
	m.CodeT = 0
	m.SuccessT = false
}

var poolTaobaoFiveeInnerproductPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeInnerproductPublishAPIResponse)
	},
}

// GetTaobaoFiveeInnerproductPublishAPIResponse 从 sync.Pool 获取 TaobaoFiveeInnerproductPublishAPIResponse
func GetTaobaoFiveeInnerproductPublishAPIResponse() *TaobaoFiveeInnerproductPublishAPIResponse {
	return poolTaobaoFiveeInnerproductPublishAPIResponse.Get().(*TaobaoFiveeInnerproductPublishAPIResponse)
}

// ReleaseTaobaoFiveeInnerproductPublishAPIResponse 将 TaobaoFiveeInnerproductPublishAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFiveeInnerproductPublishAPIResponse(v *TaobaoFiveeInnerproductPublishAPIResponse) {
	v.Reset()
	poolTaobaoFiveeInnerproductPublishAPIResponse.Put(v)
}
