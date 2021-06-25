package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡模板详情 APIRequest
alibaba.alsc.crm.card.query.template

查询卡模板详情
*/
type AlibabaAlscCrmCardQueryTemplateRequest struct {
    model.Params

    // 请求对象
    paramQueryCardTemplateOpenReq   *QueryCardTemplateOpenReq 

}

func NewAlibabaAlscCrmCardQueryTemplateRequest() *AlibabaAlscCrmCardQueryTemplateRequest{
    return &AlibabaAlscCrmCardQueryTemplateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardQueryTemplateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.query.template"
}

func (r AlibabaAlscCrmCardQueryTemplateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardQueryTemplateRequest) SetParamQueryCardTemplateOpenReq(paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq) error {
    r.paramQueryCardTemplateOpenReq = paramQueryCardTemplateOpenReq
    r.Set("param_query_card_template_open_req", paramQueryCardTemplateOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardQueryTemplateRequest) GetParamQueryCardTemplateOpenReq() *QueryCardTemplateOpenReq {
    return r.paramQueryCardTemplateOpenReq
}

