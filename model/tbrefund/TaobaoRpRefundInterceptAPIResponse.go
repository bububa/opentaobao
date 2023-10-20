package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpRefundInterceptAPIResponse 卖家发起拦截 API返回值
// taobao.rp.refund.intercept
//
// 卖家发起拦截
type TaobaoRpRefundInterceptAPIResponse struct {
	model.CommonResponse
	TaobaoRpRefundInterceptAPIResponseModel
}

// TaobaoRpRefundInterceptAPIResponseModel is 卖家发起拦截 成功返回结果
type TaobaoRpRefundInterceptAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_refund_intercept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
