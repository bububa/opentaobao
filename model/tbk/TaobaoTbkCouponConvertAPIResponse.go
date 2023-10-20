package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkcouponconvertAPIResponse 淘宝客-推广者-单品券高效转链 API返回值
// taobao.tbk.coupon.convert
//
// 单品券高效转链API
type TaobaotbkcouponconvertAPIResponse struct {
	model.CommonResponse
	TaobaotbkcouponconvertAPIResponseModel
}

// TaobaotbkcouponconvertAPIResponseModel is 淘宝客-推广者-单品券高效转链 成功返回结果
type TaobaotbkcouponconvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_coupon_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaotbkcouponconvertRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}
