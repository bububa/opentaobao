package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopbizsellersignAPIResponse 淘宝订单履约-商家erp签约 API返回值
// taobao.top.biz.seller.sign
//
// 淘宝订单履约-商家erp签约，包含各场景的签约
type TaobaotopbizsellersignAPIResponse struct {
	model.CommonResponse
	TaobaotopbizsellersignAPIResponseModel
}

// TaobaotopbizsellersignAPIResponseModel is 淘宝订单履约-商家erp签约 成功返回结果
type TaobaotopbizsellersignAPIResponseModel struct {
	XMLName xml.Name `xml:"top_biz_seller_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败code
	CallErrCode string `json:"call_err_code,omitempty" xml:"call_err_code,omitempty"`
	// 失败msg
	CallErrMsg string `json:"call_err_msg,omitempty" xml:"call_err_msg,omitempty"`
	// 调用结果
	CallResult bool `json:"call_result,omitempty" xml:"call_result,omitempty"`
}
