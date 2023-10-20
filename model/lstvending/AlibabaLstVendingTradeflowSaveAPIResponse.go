package lstvending

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingtradeflowsaveAPIResponse 自动售卖机交易信息回流 API返回值
// alibaba.lst.vending.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabalstvendingtradeflowsaveAPIResponse struct {
	model.CommonResponse
	AlibabalstvendingtradeflowsaveAPIResponseModel
}

// AlibabalstvendingtradeflowsaveAPIResponseModel is 自动售卖机交易信息回流 成功返回结果
type AlibabalstvendingtradeflowsaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_tradeflow_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *MultiResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
