package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspCountsubsidyAPIResponse
计算补助金额 API返回值
alibaba.alisports.efsp.countsubsidy

计算补助金额 */
type AlibabaAlisportsEfspCountsubsidyAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsEfspCountsubsidyAPIResponseModel
}

// AlibabaAlisportsEfspCountsubsidyAPIResponseModel is 计算补助金额 成功返回结果
type AlibabaAlisportsEfspCountsubsidyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_efsp_countsubsidy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`
}
