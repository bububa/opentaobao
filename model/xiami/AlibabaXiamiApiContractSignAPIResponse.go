package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提供签约链接 API返回值 
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url
*/
type AlibabaXiamiApiContractSignAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiContractSignAPIResponseModel
}

// 提供签约链接 成功返回结果
type AlibabaXiamiApiContractSignAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_contract_sign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 签约链接
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
}
