package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值核销预先校验 API请求
alibaba.alsc.crm.recharge.dedutprecheck.get

储值核销预先校验接口
*/
type AlibabaAlscCrmRechargeDedutprecheckGetRequest struct {
    model.Params
    // 入参
    _paramDeductPreCheckOpenReq   *DeductPreCheckOpenReq
}

// 初始化AlibabaAlscCrmRechargeDedutprecheckGetRequest对象
func NewAlibabaAlscCrmRechargeDedutprecheckGetRequest() *AlibabaAlscCrmRechargeDedutprecheckGetRequest{
    return &AlibabaAlscCrmRechargeDedutprecheckGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeDedutprecheckGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.dedutprecheck.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeDedutprecheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDeductPreCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeDedutprecheckGetRequest) SetParamDeductPreCheckOpenReq(_paramDeductPreCheckOpenReq *DeductPreCheckOpenReq) error {
    r._paramDeductPreCheckOpenReq = _paramDeductPreCheckOpenReq
    r.Set("param_deduct_pre_check_open_req", _paramDeductPreCheckOpenReq)
    return nil
}

// ParamDeductPreCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeDedutprecheckGetRequest) GetParamDeductPreCheckOpenReq() *DeductPreCheckOpenReq {
    return r._paramDeductPreCheckOpenReq
}
