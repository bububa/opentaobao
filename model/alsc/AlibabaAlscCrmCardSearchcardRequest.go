package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索卡实例列表(支持号段查询) APIRequest
alibaba.alsc.crm.card.searchcard

搜索卡实例列表(支持号段查询)
*/
type AlibabaAlscCrmCardSearchcardRequest struct {
    model.Params

    // 请求对象
    paramSearchCardOpenReq   *SearchCardOpenReq 

}

func NewAlibabaAlscCrmCardSearchcardRequest() *AlibabaAlscCrmCardSearchcardRequest{
    return &AlibabaAlscCrmCardSearchcardRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardSearchcardRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.searchcard"
}

func (r AlibabaAlscCrmCardSearchcardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardSearchcardRequest) SetParamSearchCardOpenReq(paramSearchCardOpenReq *SearchCardOpenReq) error {
    r.paramSearchCardOpenReq = paramSearchCardOpenReq
    r.Set("param_search_card_open_req", paramSearchCardOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardSearchcardRequest) GetParamSearchCardOpenReq() *SearchCardOpenReq {
    return r.paramSearchCardOpenReq
}

