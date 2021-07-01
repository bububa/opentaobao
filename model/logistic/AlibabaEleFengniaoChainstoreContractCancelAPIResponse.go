package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoChainstoreContractCancelAPIResponse
门店解约接口 API返回值
alibaba.ele.fengniao.chainstore.contract.cancel

调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单 */
type AlibabaEleFengniaoChainstoreContractCancelAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel
}

// AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel is 门店解约接口 成功返回结果
type AlibabaEleFengniaoChainstoreContractCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_contract_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
