package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletransferpayqueryAPIResponse 闲鱼转账结果查询 API返回值
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
type AlibabaidletransferpayqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidletransferpayqueryAPIResponseModel
}

// AlibabaidletransferpayqueryAPIResponseModel is 闲鱼转账结果查询 成功返回结果
type AlibabaidletransferpayqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_transferpay_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidletransferpayqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
