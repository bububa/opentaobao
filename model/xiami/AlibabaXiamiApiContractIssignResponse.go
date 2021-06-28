package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询是否签约 APIResponse
alibaba.xiami.api.contract.issign

查询是否签约
*/
type AlibabaXiamiApiContractIssignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_contract_issign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否已经签约
    
    Data   string `json:"data,omitempty" xml:"