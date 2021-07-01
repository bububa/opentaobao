package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商合约生产 API返回值 
alibaba.wtt.order.contract.subscribe

分销商合约生产
*/
type AlibabaWttOrderContractSubscribeAPIResponse struct {
    model.CommonResponse
    AlibabaWttOrderContractSubscribeAPIResponseModel
}

// 分销商合约生产 成功返回结果
type AlibabaWttOrderContractSubscribeAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wtt_order_contract_subscribe_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 合约产生陈宫
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}
