package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索卡实例列表(支持号段查询) API请求
alibaba.alsc.crm.card.searchcard

搜索卡实例列表(支持号段查询)
*/
type AlibabaAlscCrmCardSearchcardRequest struct {
    model.Params
    // 请求对象
    _paramSearchCardOpenReq   *SearchCardOpenReq
}

// 初始化AlibabaAlscCrmCardSearchcardRequest对象
func NewAlibabaAlscCrmCardSearchcardRequest() *AlibabaAlscCrmCardSearchcardRequest{
    return &AlibabaAlscCrmCardSearchcardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardSearchcardRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.searchcard"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardSearchcardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSearchCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardSearchcardRequest) SetParamSearchCardOpenReq(_paramSearchCardOpenReq *SearchCardOpenReq) error {
    r._paramSearchCardOpenReq = _paramSearchCardOpenReq
    r.Set("param_search_card_open_req", _paramSearchCardOpenReq)
    return nil
}

// ParamSearchCardOpenReq Getter
func (r AlibabaAlscCrmCardSearchcardRequest) GetParamSearchCardOpenReq() *SearchCardOpenReq {
    return r._paramSearchCardOpenReq
}
