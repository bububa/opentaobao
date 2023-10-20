package fivee

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeImportproductGetAPIResponse 进口商品查询 API返回值
// taobao.fivee.importproduct.get
//
// 资质共享平台查询进口商品信息
type TaobaoFiveeImportproductGetAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeImportproductGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFiveeImportproductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFiveeImportproductGetAPIResponseModel).Reset()
}

// TaobaoFiveeImportproductGetAPIResponseModel is 进口商品查询 成功返回结果
type TaobaoFiveeImportproductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_importproduct_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFiveeImportproductGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFiveeImportproductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFiveeImportproductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeImportproductGetAPIResponse)
	},
}

// GetTaobaoFiveeImportproductGetAPIResponse 从 sync.Pool 获取 TaobaoFiveeImportproductGetAPIResponse
func GetTaobaoFiveeImportproductGetAPIResponse() *TaobaoFiveeImportproductGetAPIResponse {
	return poolTaobaoFiveeImportproductGetAPIResponse.Get().(*TaobaoFiveeImportproductGetAPIResponse)
}

// ReleaseTaobaoFiveeImportproductGetAPIResponse 将 TaobaoFiveeImportproductGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFiveeImportproductGetAPIResponse(v *TaobaoFiveeImportproductGetAPIResponse) {
	v.Reset()
	poolTaobaoFiveeImportproductGetAPIResponse.Put(v)
}
