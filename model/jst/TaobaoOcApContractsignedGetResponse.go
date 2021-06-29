package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户是否签署支付宝代扣协议 APIResponse
taobao.oc.ap.contractsigned.get

用户是否签署支付宝代扣协议
*/
type TaobaoOcApContractsignedGetAPIResponse struct {
    model.CommonResponse
    TaobaoOcApContractsignedGetResponse
}

type TaobaoOcApContractsignedGetResponse struct {
    XMLName xml.Name `xml:"oc_ap_contractsigned_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否开通
    
    ContractSign   bool `json:"contract_sign,omitempty" xml:"contract_sign,omitempty"`

    
    // 调用出错描述信息
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
}
