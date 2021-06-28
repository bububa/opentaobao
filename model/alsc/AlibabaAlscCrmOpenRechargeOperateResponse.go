package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值操作接口 APIResponse
alibaba.alsc.crm.open.recharge.operate

储值操作接口
*/
type AlibabaAlscCrmOpenRechargeOperateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmOpenRechargeOperateResponse `json:"alibaba_alsc_crm_open_recharge_operate_response,omitempty"` 
    AlibabaAlscCrmOpenRechargeOperateResponse
}

/* model for simplify = false
type AlibabaAlscCrmOpenRechargeOperateResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmOpenRechargeOperateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
