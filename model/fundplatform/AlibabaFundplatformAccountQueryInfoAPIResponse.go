package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformaccountqueryinfoAPIResponse 查询账户信息 API返回值
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
type AlibabafundplatformaccountqueryinfoAPIResponse struct {
	model.CommonResponse
	AlibabafundplatformaccountqueryinfoAPIResponseModel
}

// AlibabafundplatformaccountqueryinfoAPIResponseModel is 查询账户信息 成功返回结果
type AlibabafundplatformaccountqueryinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_account_query_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
