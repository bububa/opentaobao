package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomorderpreauthorizequeryfundAPIResponse 资金流水查询 API返回值
// alibaba.alicom.order.preauthorize.query.fund
//
// 预授权-资金流水查询
type AlibabaalicomorderpreauthorizequeryfundAPIResponse struct {
	model.CommonResponse
	AlibabaalicomorderpreauthorizequeryfundAPIResponseModel
}

// AlibabaalicomorderpreauthorizequeryfundAPIResponseModel is 资金流水查询 成功返回结果
type AlibabaalicomorderpreauthorizequeryfundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_query_fund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
