package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiContractIssignAPIResponse 查询是否签约 API返回值
// alibaba.xiami.api.contract.issign
//
// 查询是否签约
type AlibabaXiamiApiContractIssignAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiContractIssignAPIResponseModel
}

// AlibabaXiamiApiContractIssignAPIResponseModel is 查询是否签约 成功返回结果
type AlibabaXiamiApiContractIssignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_contract_issign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否已经签约
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
