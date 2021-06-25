package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询是否签约 APIResponse
alibaba.xiami.api.contract.issign

查询是否签约
*/
type AlibabaXiamiApiContractIssignAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiContractIssignResponse `json:"alibaba_xiami_api_contract_issign_response,omitempty"`
}

type AlibabaXiamiApiContractIssignResponse struct {

    // 是否已经签约
    Data   string `json:"data,omitempty"`

}
