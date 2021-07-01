package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询储值账户信息 API请求
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口
*/
type AlibabaAlscCrmRechargeAccountGetAPIRequest struct {
    model.Params
    // 入参
    _paramQueryRechargeAccountOpenReq   *QueryRechargeAccountOpenReq
}

// 初始化AlibabaAlscCrmRechargeAccountGetAPIRequest对象
func NewAlibabaAlscCrmRechargeAccountGetRequest() *AlibabaAlscCrmRechargeAccountGetAPIRequest{
    return &AlibabaAlscCrmRechargeAccountGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.account.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQueryRechargeAccountOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountGetAPIRequest) SetParamQueryRechargeAccountOpenReq(_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq) error {
    r._paramQueryRechargeAccountOpenReq = _paramQueryRechargeAccountOpenReq
    r.Set("param_query_recharge_account_open_req", _paramQueryRechargeAccountOpenReq)
    return nil
}

// ParamQueryRechargeAccountOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetParamQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
    return r._paramQueryRechargeAccountOpenReq
}
