package logistic

import (
	"encoding/xml"

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

// TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel is 店铺授权发货注册（催发货） 成功返回结果
type TaobaoOpenSellerBizLogisticSellerBindAPIResponseModel struct {
	XMLName xml.Name `xml:"open_seller_biz_logistic_seller_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
