package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAftersaleGetAPIResponse 查询用户售后服务模板 API返回值
// taobao.aftersale.get
//
// 查询用户设置的售后服务模板，仅返回标题和id
type TaobaoAftersaleGetAPIResponse struct {
	model.CommonResponse
	TaobaoAftersaleGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAftersaleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAftersaleGetAPIResponseModel).Reset()
}

// TaobaoAftersaleGetAPIResponseModel is 查询用户售后服务模板 成功返回结果
type TaobaoAftersaleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aftersale_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 售后服务返回对象
	AfterSales []AfterSale `json:"after_sales,omitempty" xml:"after_sales>after_sale,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAftersaleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AfterSales = m.AfterSales[:0]
}

var poolTaobaoAftersaleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAftersaleGetAPIResponse)
	},
}

// GetTaobaoAftersaleGetAPIResponse 从 sync.Pool 获取 TaobaoAftersaleGetAPIResponse
func GetTaobaoAftersaleGetAPIResponse() *TaobaoAftersaleGetAPIResponse {
	return poolTaobaoAftersaleGetAPIResponse.Get().(*TaobaoAftersaleGetAPIResponse)
}

// ReleaseTaobaoAftersaleGetAPIResponse 将 TaobaoAftersaleGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAftersaleGetAPIResponse(v *TaobaoAftersaleGetAPIResponse) {
	v.Reset()
	poolTaobaoAftersaleGetAPIResponse.Put(v)
}
