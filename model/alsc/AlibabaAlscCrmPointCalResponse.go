package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
计算积分可以抵扣的金额 APIResponse
alibaba.alsc.crm.point.cal

计算积分可以抵扣的金额
积分的抵扣为区间型
如抵扣规则为100积分抵扣50元，则输入消费120积分的话，回返回消费100积分抵扣50元
 这里为纯计算逻辑，不会校验用户是否有足够的可用积分进行抵扣
*/
type AlibabaAlscCrmPointCalAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmPointCalResponse `json:"alibaba_alsc_crm_point_cal_response,omitempty"`
}

type AlibabaAlscCrmPointCalResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
