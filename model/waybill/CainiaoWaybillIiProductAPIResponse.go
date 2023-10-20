package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiProductAPIResponse 商家查询物流商产品类型接口 API返回值
// cainiao.waybill.ii.product
//
// 商家可以查询物流商的产品类型和服务能力。
type CainiaoWaybillIiProductAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiProductAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiProductAPIResponseModel).Reset()
}

// CainiaoWaybillIiProductAPIResponseModel is 商家查询物流商产品类型接口 成功返回结果
type CainiaoWaybillIiProductAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ProductTypes []WaybillProductType `json:"product_types,omitempty" xml:"product_types>waybill_product_type,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductTypes = m.ProductTypes[:0]
}

var poolCainiaoWaybillIiProductAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiProductAPIResponse)
	},
}

// GetCainiaoWaybillIiProductAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiProductAPIResponse
func GetCainiaoWaybillIiProductAPIResponse() *CainiaoWaybillIiProductAPIResponse {
	return poolCainiaoWaybillIiProductAPIResponse.Get().(*CainiaoWaybillIiProductAPIResponse)
}

// ReleaseCainiaoWaybillIiProductAPIResponse 将 CainiaoWaybillIiProductAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiProductAPIResponse(v *CainiaoWaybillIiProductAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiProductAPIResponse.Put(v)
}
