package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡实例 APIRequest
alibaba.alsc.crm.card.qry

查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
*/
type AlibabaAlscCrmCardQryRequest struct {
    model.Params

    // 请求对象
    paramQueryCardOpenReq   *QueryCardOpenReq 

}

func NewAlibabaAlscCrmCardQryRequest() *AlibabaAlscCrmCardQryRequest{
    return &AlibabaAlscCrmCardQryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardQryRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.qry"
}

func (r AlibabaAlscCrmCardQryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardQryRequest) SetParamQueryCardOpenReq(paramQueryCardOpenReq *QueryCardOpenReq) error {
    r.paramQueryCardOpenReq = paramQueryCardOpenReq
    r.Set("param_query_card_open_req", paramQueryCardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardQryRequest) GetParamQueryCardOpenReq() *QueryCardOpenReq {
    return r.paramQueryCardOpenReq
}

