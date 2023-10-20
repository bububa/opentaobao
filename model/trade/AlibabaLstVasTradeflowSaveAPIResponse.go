package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvastradeflowsaveAPIResponse 交易信息回流 API返回值
// alibaba.lst.vas.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabalstvastradeflowsaveAPIResponse struct {
	model.CommonResponse
	AlibabalstvastradeflowsaveAPIResponseModel
}

// AlibabalstvastradeflowsaveAPIResponseModel is 交易信息回流 成功返回结果
type AlibabalstvastradeflowsaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vas_tradeflow_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabalstvastradeflowsaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
