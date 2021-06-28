package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卡实例 APIResponse
alibaba.alsc.crm.card.qry

查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
*/
type AlibabaAlscCrmCardQryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardQryResponse `json:"alibaba_alsc_crm_card_qry_response,omitempty"` 
    AlibabaAlscCrmCardQryResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardQryResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardQryResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
