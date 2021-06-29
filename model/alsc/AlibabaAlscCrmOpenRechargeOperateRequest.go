package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值操作接口 API请求
alibaba.alsc.crm.open.recharge.operate

储值操作接口
*/
type AlibabaAlscCrmOpenRechargeOperateRequest struct {
    model.Params
    // 储值操作参数
    paramRechargeOperateOpenReq   *RechargeOperateOpenReq
}

// 初始化AlibabaAlscCrmOpenRechargeOperateRequest对象
func NewAlibabaAlscCrmOpenRechargeOperateRequest() *AlibabaAlscCrmOpenRechargeOperateRequest{
    return &AlibabaAlscCrmOpenRechargeOperateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenRechargeOperateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.recharge.operate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenRechargeOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRechargeOperateOpenReq Setter
// 储值操作参数
func (r *AlibabaAlscCrmOpenRechargeOperateRequest) SetParamRechargeOperateOpenReq(paramRechargeOperateOpenReq *RechargeOperateOpenReq) error {
    r.paramRechargeOperateOpenReq = paramRechargeOperateOpenReq
    r.Set("param_recharge_operate_open_req", paramRechargeOperateOpenReq)
    return nil
}

// ParamRechargeOperateOpenReq Getter
func (r AlibabaAlscCrmOpenRechargeOperateRequest) GetParamRechargeOperateOpenReq() *RechargeOperateOpenReq {
    return r.paramRechargeOperateOpenReq
}
