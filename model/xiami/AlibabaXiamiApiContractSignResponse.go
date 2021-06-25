package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提供签约链接 APIResponse
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url
*/
type AlibabaXiamiApiContractSignAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiContractSignResponse `json:"alibaba_xiami_api_contract_sign_response,omitempty"`
}

type AlibabaXiamiApiContractSignResponse struct {

    // 签约链接
    Data   string `json:"data,omitempty"`

}
