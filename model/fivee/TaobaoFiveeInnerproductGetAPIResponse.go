package fivee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeInnerproductGetAPIResponse 国产商品资质查询 API返回值
// taobao.fivee.innerproduct.get
//
// 资质共享平台，国产商品查询
type TaobaoFiveeInnerproductGetAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeInnerproductGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFiveeInnerproductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFiveeInnerproductGetAPIResponseModel).Reset()
}

// TaobaoFiveeInnerproductGetAPIResponseModel is 国产商品资质查询 成功返回结果
type TaobaoFiveeInnerproductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_innerproduct_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFiveeInnerproductGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFiveeInnerproductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFiveeInnerproductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeInnerproductGetAPIResponse)
	},
}

// GetTaobaoFiveeInnerproductGetAPIResponse 从 sync.Pool 获取 TaobaoFiveeInnerproductGetAPIResponse
func GetTaobaoFiveeInnerproductGetAPIResponse() *TaobaoFiveeInnerproductGetAPIResponse {
	return poolTaobaoFiveeInnerproductGetAPIResponse.Get().(*TaobaoFiveeInnerproductGetAPIResponse)
}

// ReleaseTaobaoFiveeInnerproductGetAPIResponse 将 TaobaoFiveeInnerproductGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFiveeInnerproductGetAPIResponse(v *TaobaoFiveeInnerproductGetAPIResponse) {
	v.Reset()
	poolTaobaoFiveeInnerproductGetAPIResponse.Put(v)
}
