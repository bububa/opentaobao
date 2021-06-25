package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡模板列表(支持数据下行) APIRequest
alibaba.alsc.crm.card.pagetmp

查询卡模板列表(支持数据下行)
当传递了数据下行参数:
     * isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
     * 否则分页查询卡模板,返回结果带有分页信息
*/
type AlibabaAlscCrmCardPagetmpRequest struct {
    model.Params

    // 请求结果
    paramPullCardTemplateOpenReq   *PullCardTemplateOpenReq 

}

func NewAlibabaAlscCrmCardPagetmpRequest() *AlibabaAlscCrmCardPagetmpRequest{
    return &AlibabaAlscCrmCardPagetmpRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardPagetmpRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.pagetmp"
}

func (r AlibabaAlscCrmCardPagetmpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardPagetmpRequest) SetParamPullCardTemplateOpenReq(paramPullCardTemplateOpenReq *PullCardTemplateOpenReq) error {
    r.paramPullCardTemplateOpenReq = paramPullCardTemplateOpenReq
    r.Set("param_pull_card_template_open_req", paramPullCardTemplateOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardPagetmpRequest) GetParamPullCardTemplateOpenReq() *PullCardTemplateOpenReq {
    return r.paramPullCardTemplateOpenReq
}

