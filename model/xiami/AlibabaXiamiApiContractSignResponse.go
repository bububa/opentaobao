package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提供签约链接 APIResponse
alibaba.xiami.api.contract.sign

提供签约链接。in：商家id；out：签约url
*/
type AlibabaXiamiApiContractSignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_contract_sign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 签约链接
    
    Data   string `json:"data,omitempty" xml:"