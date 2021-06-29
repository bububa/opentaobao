package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡实例 API请求
alibaba.alsc.crm.card.qry

查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
*/
type AlibabaAlscCrmCardQryRequest struct {
    model.Params
    // 请求对象
    paramQueryCardOpenReq   *QueryCardOpenReq
}

// 初始化AlibabaAlscCrmCardQryRequest对象
func NewAlibabaAlscCrmCardQryRequest() *AlibabaAlscCrmCardQryRequest{
    return &AlibabaAlscCrmCardQryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQryRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.qry"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQueryCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardQryRequest) SetParamQueryCardOpenReq(paramQueryCardOpenReq *QueryCardOpenReq) error {
    r.paramQueryCardOpenReq = paramQueryCardOpenReq
    r.Set("param_query_card_open_req", paramQueryCardOpenReq)
    return nil
}

// ParamQueryCardOpenReq Getter
func (r AlibabaAlscCrmCardQryRequest) GetParamQueryCardOpenReq() *QueryCardOpenReq {
    return r.paramQueryCardOpenReq
}
