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
type AlibabaAlscCrmRechargeAccountGetRequest struct {
    model.Params
    // 入参
    paramQueryRechargeAccountOpenReq   *QueryRechargeAccountOpenReq
}

// 初始化AlibabaAlscCrmRechargeAccountGetRequest对象
func NewAlibabaAlscCrmRechargeAccountGetRequest() *AlibabaAlscCrmRechargeAccountGetRequest{
    return &AlibabaAlscCrmRechargeAccountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.account.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQueryRechargeAccountOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountGetRequest) SetParamQueryRechargeAccountOpenReq(paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq) error {
    r.paramQueryRechargeAccountOpenReq = paramQueryRechargeAccountOpenReq
    r.Set("param_query_recharge_account_open_req", paramQueryRechargeAccountOpenReq)
    return nil
}

// ParamQueryRechargeAccountOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountGetRequest) GetParamQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
    return r.paramQueryRechargeAccountOpenReq
}
