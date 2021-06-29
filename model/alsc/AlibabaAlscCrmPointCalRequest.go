package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算积分可以抵扣的金额 API请求
alibaba.alsc.crm.point.cal

计算积分可以抵扣的金额
积分的抵扣为区间型
如抵扣规则为100积分抵扣50元，则输入消费120积分的话，回返回消费100积分抵扣50元
 这里为纯计算逻辑，不会校验用户是否有足够的可用积分进行抵扣
*/
type AlibabaAlscCrmPointCalRequest struct {
    model.Params
    // 入参
    paramCalculateDeductedMoneyOpenReq   *CalculateDeductedMoneyOpenReq
}

// 初始化AlibabaAlscCrmPointCalRequest对象
func NewAlibabaAlscCrmPointCalRequest() *AlibabaAlscCrmPointCalRequest{
    return &AlibabaAlscCrmPointCalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointCalRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.cal"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointCalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCalculateDeductedMoneyOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointCalRequest) SetParamCalculateDeductedMoneyOpenReq(paramCalculateDeductedMoneyOpenReq *CalculateDeductedMoneyOpenReq) error {
    r.paramCalculateDeductedMoneyOpenReq = paramCalculateDeductedMoneyOpenReq
    r.Set("param_calculate_deducted_money_open_req", paramCalculateDeductedMoneyOpenReq)
    return nil
}

// ParamCalculateDeductedMoneyOpenReq Getter
func (r AlibabaAlscCrmPointCalRequest) GetParamCalculateDeductedMoneyOpenReq() *CalculateDeductedMoneyOpenReq {
    return r.paramCalculateDeductedMoneyOpenReq
}
