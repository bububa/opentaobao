package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcGetproductbyscancodeAPIResponse POS商品查询接口 API返回值
// alibaba.mj.oc.getproductbyscancode
//
// 此API用于在银泰商场中，POS端扫码获取商品信息
type AlibabaMjOcGetproductbyscancodeAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcGetproductbyscancodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcGetproductbyscancodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcGetproductbyscancodeAPIResponseModel).Reset()
}

// AlibabaMjOcGetproductbyscancodeAPIResponseModel is POS商品查询接口 成功返回结果
type AlibabaMjOcGetproductbyscancodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_getproductbyscancode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	ProductList []ScanProduct `json:"product_list,omitempty" xml:"product_list>scan_product,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcGetproductbyscancodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductList = m.ProductList[:0]
}

var poolAlibabaMjOcGetproductbyscancodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcGetproductbyscancodeAPIResponse)
	},
}

// GetAlibabaMjOcGetproductbyscancodeAPIResponse 从 sync.Pool 获取 AlibabaMjOcGetproductbyscancodeAPIResponse
func GetAlibabaMjOcGetproductbyscancodeAPIResponse() *AlibabaMjOcGetproductbyscancodeAPIResponse {
	return poolAlibabaMjOcGetproductbyscancodeAPIResponse.Get().(*AlibabaMjOcGetproductbyscancodeAPIResponse)
}

// ReleaseAlibabaMjOcGetproductbyscancodeAPIResponse 将 AlibabaMjOcGetproductbyscancodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcGetproductbyscancodeAPIResponse(v *AlibabaMjOcGetproductbyscancodeAPIResponse) {
	v.Reset()
	poolAlibabaMjOcGetproductbyscancodeAPIResponse.Put(v)
}
