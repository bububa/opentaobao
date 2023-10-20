package ottpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderauthpaywithpriceAPIResponse 委托代扣可配定价服务 API返回值
// youku.ott.pay.order.authpaywithprice
//
// 应用中心sdk连续包月委托代扣服务，次月可配置营销价
type YoukuottpayorderauthpaywithpriceAPIResponse struct {
	model.CommonResponse
	YoukuottpayorderauthpaywithpriceAPIResponseModel
}

// YoukuottpayorderauthpaywithpriceAPIResponseModel is 委托代扣可配定价服务 成功返回结果
type YoukuottpayorderauthpaywithpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_authpaywithprice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}
