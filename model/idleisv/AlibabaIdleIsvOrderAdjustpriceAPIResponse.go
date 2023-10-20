package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvorderadjustpriceAPIResponse 闲鱼服务商订单价格修改接口 API返回值
// alibaba.idle.isv.order.adjustprice
//
// 闲鱼用户通过授权的服务商修改订单价格和邮费
type AlibabaidleisvorderadjustpriceAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvorderadjustpriceAPIResponseModel
}

// AlibabaidleisvorderadjustpriceAPIResponseModel is 闲鱼服务商订单价格修改接口 成功返回结果
type AlibabaidleisvorderadjustpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_order_adjustprice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaidleisvorderadjustpriceResult `json:"result,omitempty" xml:"result,omitempty"`
}
