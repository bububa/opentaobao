package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户是否签署支付宝代扣协议 APIResponse
taobao.oc.ap.contractsigned.get

用户是否签署支付宝代扣协议
*/
type TaobaoOcApContractsignedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOcApContractsignedGetResponse `json:"taobao_oc_ap_contractsigned_get_response,omitempty"`
}

type TaobaoOcApContractsignedGetResponse struct {

    // 是否开通
    ContractSign   bool `json:"contract_sign,omitempty"`

    // 调用出错描述信息
    ErrorDescription   string `json:"error_description,omitempty"`

}
