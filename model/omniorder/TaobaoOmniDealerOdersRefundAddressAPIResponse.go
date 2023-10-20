package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomnidealerodersrefundaddressAPIResponse 经销商查询逆向退货地址 API返回值
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
type TaobaoomnidealerodersrefundaddressAPIResponse struct {
	model.CommonResponse
	TaobaoomnidealerodersrefundaddressAPIResponseModel
}

// TaobaoomnidealerodersrefundaddressAPIResponseModel is 经销商查询逆向退货地址 成功返回结果
type TaobaoomnidealerodersrefundaddressAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_dealer_oders_refund_address_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经销商订单退货地址
	Result *TaobaoomnidealerodersrefundaddressResult `json:"result,omitempty" xml:"result,omitempty"`
}
