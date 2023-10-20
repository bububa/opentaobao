package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenSellerBizLogisticSellerBindAPIResponse 店铺授权发货注册（催发货） API返回值
// taobao.open.seller.biz.logistic.seller.bind
//
// 店铺授权发货注册（催发货）
type TaobaoOpenSellerBizLogisticSellerBindAPIResponse struct {
	model.CommonResponse
	TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenSellerBizLogisticSellerBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel).Reset()
}

// TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel is 店铺授权发货注册（催发货） 成功返回结果
type TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel struct {
	XMLName xml.Name `xml:"open_seller_biz_logistic_seller_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoOpenSellerBizLogisticSellerBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenSellerBizLogisticSellerBindAPIResponse)
	},
}

// GetTaobaoOpenSellerBizLogisticSellerBindAPIResponse 从 sync.Pool 获取 TaobaoOpenSellerBizLogisticSellerBindAPIResponse
func GetTaobaoOpenSellerBizLogisticSellerBindAPIResponse() *TaobaoOpenSellerBizLogisticSellerBindAPIResponse {
	return poolTaobaoOpenSellerBizLogisticSellerBindAPIResponse.Get().(*TaobaoOpenSellerBizLogisticSellerBindAPIResponse)
}

// ReleaseTaobaoOpenSellerBizLogisticSellerBindAPIResponse 将 TaobaoOpenSellerBizLogisticSellerBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenSellerBizLogisticSellerBindAPIResponse(v *TaobaoOpenSellerBizLogisticSellerBindAPIResponse) {
	v.Reset()
	poolTaobaoOpenSellerBizLogisticSellerBindAPIResponse.Put(v)
}
