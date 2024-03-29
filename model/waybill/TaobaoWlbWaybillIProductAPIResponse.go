package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIProductAPIResponse 商家查询物流商产品类型接口 API返回值
// taobao.wlb.waybill.i.product
//
// 商家可以查询物流商的产品类型和服务能力。
type TaobaoWlbWaybillIProductAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillIProductAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWaybillIProductAPIResponseModel).Reset()
}

// TaobaoWlbWaybillIProductAPIResponseModel is 商家查询物流商产品类型接口 成功返回结果
type TaobaoWlbWaybillIProductAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品类型返回
	ProductTypes []WaybillProductType `json:"product_types,omitempty" xml:"product_types>waybill_product_type,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductTypes = m.ProductTypes[:0]
}

var poolTaobaoWlbWaybillIProductAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWaybillIProductAPIResponse)
	},
}

// GetTaobaoWlbWaybillIProductAPIResponse 从 sync.Pool 获取 TaobaoWlbWaybillIProductAPIResponse
func GetTaobaoWlbWaybillIProductAPIResponse() *TaobaoWlbWaybillIProductAPIResponse {
	return poolTaobaoWlbWaybillIProductAPIResponse.Get().(*TaobaoWlbWaybillIProductAPIResponse)
}

// ReleaseTaobaoWlbWaybillIProductAPIResponse 将 TaobaoWlbWaybillIProductAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWaybillIProductAPIResponse(v *TaobaoWlbWaybillIProductAPIResponse) {
	v.Reset()
	poolTaobaoWlbWaybillIProductAPIResponse.Put(v)
}
