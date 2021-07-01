package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoChainstoreContractChangeAPIResponse
门店改签合同接口 API返回值
alibaba.ele.fengniao.chainstore.contract.change

通过调用接口，门店改签物流服务包 */
type AlibabaEleFengniaoChainstoreContractChangeAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoChainstoreContractChangeAPIResponseModel
}

// AlibabaEleFengniaoChainstoreContractChangeAPIResponseModel is 门店改签合同接口 成功返回结果
type AlibabaEleFengniaoChainstoreContractChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_contract_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
